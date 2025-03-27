package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chzyer/readline"
)

func main() {
	e := EnvNew()
	e.AddBuiltins()

	if len(os.Args) == 1 {
		fmt.Println("Lispy Version 0.0.0.1.0")
		fmt.Printf("Press Ctrl+c to Exit\n\n")

		rl, err := readline.NewEx(&readline.Config{
			Prompt:          "lispy> ",
			HistoryFile:     ".lispy.his",
			InterruptPrompt: "^C",
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()

		for {
			line, err := rl.Readline()
			if err != nil {
				if errors.Is(err, readline.ErrInterrupt) || err == io.EOF {
					break
				}
				fmt.Printf("Error: %v\n", err)
				continue
			}
			if line == "" {
				continue
			}
			tree, err := ParseContents(line)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				x := ValEval(e, ValRead(tree))
				fmt.Println(x)
			}
		}
	}

	if len(os.Args) >= 2 {
		for i := 1; i < len(os.Args); i++ {
			args := ValAdd(NewValSexpr(), NewValStr(os.Args[i]))
			x := BuiltinLoad(e, args)

			if x.t == ValErr {
				fmt.Println(x)
			}
		}
	}

}

func ParseContents(input string) (antlr.Tree, error) {
	is := antlr.NewInputStream(input)
	ec := &ErrorCollector{}
	lexer := NewLispyLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ec)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := NewLispyParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(ec)

	tree := parser.Lispy()
	if len(ec.Errors) > 0 {
		return nil, errors.New(strings.Join(ec.Errors, "\n"))
	}
	return tree, nil
}

type ValType int

const (
	ValErr ValType = iota
	ValNum
	ValSym
	ValStr
	ValFun
	ValSexpr
	ValQexpr
)

func TypeName(t ValType) string {
	switch t {
	case ValFun:
		return "Function"
	case ValNum:
		return "Number"
	case ValErr:
		return "Error"
	case ValSym:
		return "Symbol"
	case ValStr:
		return "String"
	case ValSexpr:
		return "S-Expression"
	case ValQexpr:
		return "Q-Expression"
	default:
		return "Unknown"
	}
}

type BuiltinFunc func(*Env, *Val) *Val

type Env struct {
	par  *Env
	vals map[string]*Val
}

func EnvNew() *Env {
	return &Env{
		vals: make(map[string]*Val),
	}
}

func (env *Env) Get(k *Val) *Val {
	v, ok := env.vals[k.sym]
	if ok {
		return ValCopy(v)
	}
	if env.par != nil {
		return env.par.Get(k)
	}
	return NewValErr("Unbound Symbol '%s'", k.sym)
}

func (env *Env) Put(k *Val, v *Val) {
	env.vals[k.sym] = ValCopy(v)
}

func (env *Env) Copy() *Env {
	n := EnvNew()
	n.par = env.par
	for k, v := range env.vals {
		n.Put(NewValSym(k), ValCopy(v))
	}
	return n
}

func (env *Env) Def(k *Val, v *Val) {
	/* Iterate till e has no parent */
	e := env
	for e.par != nil {
		e = e.par
	}
	/* Put value in e */
	e.Put(k, v)
}

type Val struct {
	t ValType

	/* Basic */
	num int
	err string
	sym string
	str string

	/* Function */
	builtin BuiltinFunc
	env     *Env
	formals *Val
	body    *Val

	/* Expression */
	cell []*Val
}

func (v *Val) Count() int {
	return len(v.cell)
}

func NewValNum(x int) *Val {
	return &Val{t: ValNum, num: x}
}

func NewValErr(f string, a ...any) *Val {
	return &Val{t: ValErr, err: fmt.Sprintf(f, a...)}
}

func NewValSym(s string) *Val { return &Val{t: ValSym, sym: s} }

func NewValStr(s string) *Val { return &Val{t: ValStr, str: s} }

func NewValFun(f BuiltinFunc) *Val {
	return &Val{t: ValFun, builtin: f}
}

func NewLambda(formals *Val, body *Val) *Val {
	return &Val{
		t:       ValFun,
		env:     EnvNew(),
		formals: formals,
		body:    body,
	}
}

func NewValSexpr() *Val { return &Val{t: ValSexpr} }

func NewValQexpr() *Val { return &Val{t: ValQexpr} }

func ValCopy(v *Val) *Val {
	x := &Val{t: v.t}
	switch v.t {
	case ValFun:
		if v.builtin != nil {
			x.builtin = v.builtin
		} else {
			x.env = v.env.Copy()
			x.formals = ValCopy(v.formals)
			x.body = ValCopy(v.body)
		}
	case ValNum:
		x.num = v.num
	case ValErr:
		x.err = v.err
	case ValSym:
		x.sym = v.sym
	case ValStr:
		x.str = v.str
	case ValSexpr:
		fallthrough
	case ValQexpr:
		for _, c := range v.cell {
			x.cell = append(x.cell, ValCopy(c))
		}
	}
	return x
}

func ValAdd(v *Val, x *Val) *Val {
	v.cell = append(v.cell, x)
	return v
}

func ValJoin(x *Val, y *Val) *Val {
	x.cell = append(x.cell, y.cell...)
	y.cell = nil
	return x
}

func ValPop(v *Val, i int) *Val {
	x := v.cell[i]
	v.cell = append(v.cell[:i], v.cell[i+1:]...)
	return x
}

func ValTake(v *Val, i int) *Val {
	x := v.cell[i]
	v.cell = nil
	return x
}

func ExprToString(v *Val, open, close string) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(open)
	for i := 0; i < v.Count(); i++ {
		buf.WriteString(v.cell[i].String())
		if i != v.Count()-1 {
			buf.WriteString(" ")
		}
	}
	buf.WriteString(close)
	return buf.String()
}

func (v *Val) String() string {
	switch v.t {
	case ValFun:
		if v.builtin != nil {
			return "<function>"
		}
		return fmt.Sprintf("(\\ %s %s)", v.formals.String(), v.body.String())
	case ValNum:
		return fmt.Sprintf("%d", v.num)
	case ValErr:
		return fmt.Sprintf("Error: %s", v.err)
	case ValSym:
		return v.sym
	case ValStr:
		return fmt.Sprintf("%q", v.str)
	case ValSexpr:
		return ExprToString(v, "(", ")")
	case ValQexpr:
		return ExprToString(v, "{", "}")
	}
	return ""
}

func (v *Val) Eq(y *Val) bool {
	x := v
	if x.t != y.t {
		return false
	}
	switch x.t {
	case ValNum:
		return x.num == y.num
	case ValErr:
		return x.err == y.err
	case ValSym:
		return x.sym == y.sym
	case ValStr:
		return x.str == y.str
	case ValFun:
		if x.builtin != nil && y.builtin == nil {
			return false
		}
		if x.builtin == nil && y.builtin != nil {
			return false
		}
		if x.builtin != nil && y.builtin != nil {
			return fmt.Sprintf("%s", x.builtin) == fmt.Sprintf("%s", y.builtin)
		}
		return x.formals.Eq(y.formals) && x.body.Eq(y.body)
	case ValQexpr:
		fallthrough
	case ValSexpr:
		if x.Count() != y.Count() {
			return false
		}
		for i := range x.Count() {
			if !x.cell[i].Eq(y.cell[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func BuiltinLambda(e *Env, a *Val) *Val {
	if err := AssertNum("\\", a, 2); err != nil {
		return err
	}
	if err := AssertType("\\", a, 0, ValQexpr); err != nil {
		return err
	}
	if err := AssertType("\\", a, 0, ValQexpr); err != nil {
		return err
	}

	for i := range a.cell[0].Count() {
		if err := Assert(a, a.cell[0].cell[i].t == ValSym, "Cannot define non-symbol. Got %s, Expected %s.", TypeName(a.cell[0].cell[i].t), TypeName(ValSym)); err != nil {
			return err
		}
	}
	formals := ValPop(a, 0)
	body := ValPop(a, 0)

	return NewLambda(formals, body)
}

func BuiltinList(e *Env, a *Val) *Val {
	a.t = ValQexpr
	return a
}

func Assert(args *Val, cond bool, format string, a ...any) *Val {
	if !cond {
		return NewValErr(format, a...)
	}
	return nil
}

func AssertType(fun string, args *Val, index int, expect ValType) *Val {
	return Assert(args, args.cell[index].t == expect, "Function '%s' passed incorrect type for argument %i. Got %s, Expected %s.", fun, index, TypeName(args.cell[index].t), TypeName(expect))
}

func AssertNum(fun string, args *Val, num int) *Val {
	return Assert(args, args.Count() == num, "Function '%s' passed incorrect number of arguments. Got %i, Expected %i.", fun, args.Count(), num)
}

func AssertNotEmpty(fun string, args *Val, index int) *Val {
	return Assert(args, args.cell[index].Count() != 0, "Function '%s' passed {} for argument %i.", fun, index)
}

func BuiltinHead(e *Env, a *Val) *Val {
	if err := AssertNum("head", a, 1); err != nil {
		return err
	}
	if err := AssertType("head", a, 0, ValQexpr); err != nil {
		return err
	}
	if err := AssertNotEmpty("head", a, 0); err != nil {
		return err
	}
	v := ValTake(a, 0)
	for v.Count() > 1 {
		ValPop(v, 1)
	}
	return v
}

func BuiltinTail(e *Env, a *Val) *Val {
	if err := AssertNum("tail", a, 1); err != nil {
		return err
	}
	if err := AssertType("tail", a, 0, ValQexpr); err != nil {
		return err
	}
	if err := AssertNotEmpty("tail", a, 0); err != nil {
		return err
	}
	v := ValTake(a, 0)
	ValPop(v, 0)
	return v
}

func BuiltinEval(e *Env, a *Val) *Val {
	if err := AssertNum("eval", a, 1); err != nil {
		return err
	}
	if err := AssertType("eval", a, 0, ValQexpr); err != nil {
		return err
	}
	x := ValTake(a, 0)
	x.t = ValSexpr
	return ValEval(e, x)
}

func BuiltinJoin(e *Env, a *Val) *Val {
	for i := range a.Count() {
		if err := AssertType("join", a, i, ValQexpr); err != nil {
			return err
		}
	}
	x := ValPop(a, 0)
	for a.Count() > 0 {
		x = ValJoin(x, ValPop(a, 0))
	}
	return x
}

func BuiltinOp(e *Env, a *Val, op string) *Val {
	for _, c := range a.cell {
		if c.t != ValNum {
			return NewValErr("Cannot operate on non-number!")
		}
	}
	x := ValPop(a, 0)
	if op == "-" && a.Count() == 0 {
		x.num = -x.num
	}

	for a.Count() > 0 {
		y := ValPop(a, 0)
		switch op {
		case "+":
			x.num += y.num
		case "-":
			x.num -= y.num
		case "*":
			x.num *= y.num
		case "/":
			if y.num == 0 {
				x = NewValErr("Division by zero.")
			} else {
				x.num /= y.num
			}
		}
	}
	return x
}

func BuiltinAdd(e *Env, a *Val) *Val {
	return BuiltinOp(e, a, "+")
}

func BuiltinSub(e *Env, a *Val) *Val {
	return BuiltinOp(e, a, "-")
}

func BuiltinMul(e *Env, a *Val) *Val {
	return BuiltinOp(e, a, "*")
}

func BuiltinDiv(e *Env, a *Val) *Val {
	return BuiltinOp(e, a, "/")
}

func BuiltinVar(e *Env, a *Val, fun string) *Val {
	if err := AssertType(fun, a, 0, ValQexpr); err != nil {
		return err
	}
	syms := a.cell[0]
	for i := range syms.Count() {
		if err := Assert(a, syms.cell[i].t == ValSym, "Function '%s' cannot define non-symbol. Got %s, Expected %s.", fun, TypeName(syms.cell[i].t), TypeName(ValSym)); err != nil {
			return err
		}
	}
	if err := Assert(a, syms.Count() == a.Count()-1, "Function '%s' passed too many arguments for symbols. Got %d, Expected %d.", fun, syms.Count(), a.Count()-1); err != nil {
		return err
	}
	for i := range syms.Count() {
		if fun == "def" {
			e.Def(syms.cell[i], a.cell[i+1])
		}
		if fun == "=" {
			e.Put(syms.cell[i], a.cell[i+1])
		}
	}
	return NewValSexpr()
}

func BuiltinDef(e *Env, a *Val) *Val {
	return BuiltinVar(e, a, "def")
}

func BuiltinPut(e *Env, a *Val) *Val {
	return BuiltinVar(e, a, "=")
}

func BuiltinOrd(e *Env, a *Val, op string) *Val {
	if err := AssertNum(op, a, 2); err != nil {
		return err
	}
	if err := AssertType(op, a, 0, ValNum); err != nil {
		return err
	}
	if err := AssertType(op, a, 1, ValNum); err != nil {
		return err
	}
	var r bool
	if op == ">" {
		r = a.cell[0].num > a.cell[1].num
	}
	if op == "<" {
		r = a.cell[0].num < a.cell[1].num
	}
	if op == ">=" {
		r = a.cell[0].num >= a.cell[1].num
	}
	if op == "<=" {
		r = a.cell[0].num <= a.cell[1].num
	}
	num := NewValNum(0)
	if r {
		num.num = 1
	}
	return num
}

func BuiltinGt(e *Env, a *Val) *Val {
	return BuiltinOrd(e, a, ">")
}

func BuiltinLt(e *Env, a *Val) *Val {
	return BuiltinOrd(e, a, "<")
}

func BuiltinGe(e *Env, a *Val) *Val {
	return BuiltinOrd(e, a, ">=")
}

func BuiltinLe(e *Env, a *Val) *Val {
	return BuiltinOrd(e, a, "<=")
}

func BuiltinCmp(e *Env, a *Val, op string) *Val {
	if err := AssertNum(op, a, 2); err != nil {
		return err
	}
	var r bool
	if op == "==" {
		r = a.cell[0].Eq(a.cell[1])
	}
	if op == "!=" {
		r = !a.cell[0].Eq(a.cell[1])
	}
	num := NewValNum(0)
	if r {
		num.num = 1
	}
	return num
}

func BuiltinEq(e *Env, a *Val) *Val {
	return BuiltinCmp(e, a, "==")
}

func BuiltinNe(e *Env, a *Val) *Val {
	return BuiltinCmp(e, a, "!=")
}

func BuiltinIf(e *Env, a *Val) *Val {
	if err := AssertNum("if", a, 3); err != nil {
		return err
	}
	if err := AssertType("if", a, 0, ValNum); err != nil {
		return err
	}
	if err := AssertType("if", a, 1, ValQexpr); err != nil {
		return err
	}
	if err := AssertType("if", a, 2, ValQexpr); err != nil {
		return err
	}
	a.cell[1].t = ValSexpr
	a.cell[2].t = ValSexpr

	var x *Val
	if a.cell[0].num != 0 {
		x = ValEval(e, ValPop(a, 1))
	} else {
		x = ValEval(e, ValPop(a, 2))
	}
	return x
}

func BuiltinLoad(e *Env, a *Val) *Val {
	if err := AssertNum("load", a, 1); err != nil {
		return err
	}
	if err := AssertType("load", a, 0, ValStr); err != nil {
		return err
	}
	var (
		content []byte
		err     error
	)
	if a.cell[0].str == "-" {
		content, err = io.ReadAll(os.Stdin)
	} else {
		content, err = os.ReadFile(a.cell[0].str)
	}
	if err != nil {
		return NewValErr("Could not load Library %s", err.Error())
	}
	tree, err := ParseContents(string(content))
	if err != nil {
		return NewValErr("Could not load Library %s", err.Error())
	}
	expr := ValRead(tree)
	for expr.Count() > 0 {
		x := ValEval(e, ValPop(expr, 0))
		if x.t == ValErr {
			fmt.Println(x)
		}
	}
	return NewValSexpr()
}

func BuiltinPrint(e *Env, a *Val) *Val {
	cnt := a.Count()
	for i := range cnt {
		fmt.Print(a.cell[i])
		if i != cnt-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	return NewValSexpr()
}

func BuiltinError(e *Env, a *Val) *Val {
	if err := AssertNum("error", a, 1); err != nil {
		return err
	}
	if err := AssertType("error", a, 0, ValStr); err != nil {
		return err
	}

	return NewValErr(a.cell[0].str)
}

func (env *Env) AddBuiltin(name string, builtinFunc BuiltinFunc) {
	k := NewValSym(name)
	v := NewValFun(builtinFunc)
	env.Put(k, v)
}

func (env *Env) AddBuiltins() {
	/* Variable Functions */
	env.AddBuiltin("\\", BuiltinLambda)
	env.AddBuiltin("def", BuiltinDef)
	env.AddBuiltin("=", BuiltinPut)

	/* List Functions */
	env.AddBuiltin("list", BuiltinList)
	env.AddBuiltin("head", BuiltinHead)
	env.AddBuiltin("tail", BuiltinTail)
	env.AddBuiltin("eval", BuiltinEval)
	env.AddBuiltin("join", BuiltinJoin)

	/* Mathematical Functions */
	env.AddBuiltin("+", BuiltinAdd)
	env.AddBuiltin("-", BuiltinSub)
	env.AddBuiltin("*", BuiltinMul)
	env.AddBuiltin("/", BuiltinDiv)

	/* Comparison Functions */
	env.AddBuiltin("if", BuiltinIf)
	env.AddBuiltin("==", BuiltinEq)
	env.AddBuiltin("!=", BuiltinNe)
	env.AddBuiltin(">", BuiltinGt)
	env.AddBuiltin("<", BuiltinLt)
	env.AddBuiltin(">=", BuiltinGe)
	env.AddBuiltin("<=", BuiltinLe)

	/* String Functions */
	env.AddBuiltin("load", BuiltinLoad)
	env.AddBuiltin("error", BuiltinError)
	env.AddBuiltin("print", BuiltinPrint)
}

func ValCall(e *Env, f *Val, a *Val) *Val {
	if f.builtin != nil {
		return f.builtin(e, a)
	}
	given := a.Count()
	total := f.formals.Count()

	for a.Count() > 0 {
		if f.formals.Count() == 0 {
			return NewValErr("Function passed too many arguments. Got %d, Expected %d.", given, total)
		}
		sym := ValPop(f.formals, 0)
		if sym.sym == "&" {
			if f.formals.Count() != 1 {
				return NewValErr("Function format invalid. Symbol '&' not followed by single symbol.")
			}
			nsym := ValPop(f.formals, 0)
			f.env.Put(nsym, BuiltinList(e, a))
			break
		}
		val := ValPop(a, 0)
		f.env.Put(sym, val)
	}
	if f.formals.Count() > 0 && f.formals.cell[0].sym == "&" {
		if f.formals.Count() != 2 {
			return NewValErr("Function format invalid. Symbol '&' not followed by single symbol.")
		}
		ValPop(f.formals, 0)
		sym := ValPop(f.formals, 0)
		val := NewValQexpr()
		f.env.Put(sym, val)
	}
	if f.formals.Count() == 0 {
		f.env.par = e
		return BuiltinEval(f.env, ValAdd(NewValSexpr(), ValCopy(f.body)))
	}
	return ValCopy(f)
}

func ValEval(e *Env, v *Val) *Val {
	if v.t == ValSym {
		x := e.Get(v)
		return x
	}
	if v.t == ValSexpr {
		return ValEvalSexpr(e, v)
	}
	return v
}

func ValEvalSexpr(e *Env, v *Val) *Val {
	for i := range v.Count() {
		v.cell[i] = ValEval(e, v.cell[i])
	}

	for i := range v.Count() {
		if v.cell[i].t == ValErr {
			return ValTake(v, i)
		}
	}

	if v.Count() == 0 {
		return v
	}

	if v.Count() == 1 {
		return ValEval(e, ValTake(v, 0))
	}

	/* Ensure first element is a function after evaluation */
	f := ValPop(v, 0)
	if f.t != ValFun {
		err := NewValErr("S-Expression starts with incorrect type. Got %s, Expected %s.", TypeName(f.t), TypeName(ValFun))
		return err
	}
	return ValCall(e, f, v)
}

func ValReadNum(t *NumberContext) *Val {
	x, err := strconv.Atoi(t.GetText())
	if err != nil {
		return NewValErr(err.Error())
	}
	return NewValNum(x)
}

func ValReadStr(t *StringContext) *Val {
	s, err := strconv.Unquote(t.GetText())
	if err != nil {
		return NewValErr(err.Error())
	}
	return NewValStr(s)
}

func ValRead(t antlr.Tree) *Val {
	var x *Val
	switch t2 := t.(type) {
	case *NumberContext:
		return ValReadNum(t2)
	case *SymbolContext:
		return NewValSym(t2.GetText())
	case *StringContext:
		return ValReadStr(t2)
	case *LispyContext:
		x = NewValSexpr()
	case *SexprContext:
		x = NewValSexpr()
	case *QexprContext:
		x = NewValQexpr()
	case *ExprContext:
		return ValRead(t2.GetChild(0))
	case antlr.TerminalNode:
		return nil
	case *CommentContext:
		return nil
	}
	for i := range t.GetChildCount() {
		v := ValRead(t.GetChild(i))
		if v != nil {
			ValAdd(x, v)
		}
	}
	return x
}
