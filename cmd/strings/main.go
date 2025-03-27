package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chzyer/readline"
)

func main() {
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

	e := EnvNew()
	e.AddBuiltins()

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
		is := antlr.NewInputStream(line)
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
			fmt.Println(ec)
		} else {
			x := ValEval(e, ValRead(tree))
			fmt.Println(x)
		}
	}
}

type ValType int

const (
	ValErr ValType = iota
	ValNum
	ValSym
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
	v, ok := env.vals[k.Sym]
	if ok {
		return ValCopy(v)
	}
	if env.par != nil {
		return env.par.Get(k)
	}
	return NewValErr("Unbound Symbol '%s'", k.Sym)
}

func (env *Env) Put(k *Val, v *Val) {
	env.vals[k.Sym] = ValCopy(v)
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
	Type ValType

	/* Basic */
	Num int
	Err string
	Sym string

	/* Function */
	Builtin BuiltinFunc
	Env     *Env
	Formals *Val
	Body    *Val

	/* Expression */
	Cell []*Val
}

func (v *Val) Count() int {
	return len(v.Cell)
}

func NewValNum(x int) *Val {
	return &Val{Type: ValNum, Num: x}
}

func NewValErr(f string, a ...any) *Val {
	return &Val{Type: ValErr, Err: fmt.Sprintf(f, a...)}
}

func NewValSym(s string) *Val { return &Val{Type: ValSym, Sym: s} }

func NewValFun(f BuiltinFunc) *Val {
	return &Val{Type: ValFun, Builtin: f}
}

func NewLambda(formals *Val, body *Val) *Val {
	return &Val{
		Type:    ValFun,
		Env:     EnvNew(),
		Formals: formals,
		Body:    body,
	}
}

func NewValSexpr() *Val { return &Val{Type: ValSexpr} }

func NewValQexpr() *Val { return &Val{Type: ValQexpr} }

func ValCopy(v *Val) *Val {
	x := &Val{Type: v.Type}
	switch v.Type {
	case ValFun:
		if v.Builtin != nil {
			x.Builtin = v.Builtin
		} else {
			x.Env = v.Env.Copy()
			x.Formals = ValCopy(v.Formals)
			x.Body = ValCopy(v.Body)
		}
	case ValNum:
		x.Num = v.Num
	case ValErr:
		x.Err = v.Err
	case ValSym:
		x.Sym = v.Sym
	case ValSexpr:
		fallthrough
	case ValQexpr:
		for _, c := range v.Cell {
			x.Cell = append(x.Cell, ValCopy(c))
		}
	}
	return x
}

func ValAdd(v *Val, x *Val) *Val {
	v.Cell = append(v.Cell, x)
	return v
}

func ValJoin(x *Val, y *Val) *Val {
	x.Cell = append(x.Cell, y.Cell...)
	y.Cell = nil
	return x
}

func ValPop(v *Val, i int) *Val {
	x := v.Cell[i]
	v.Cell = append(v.Cell[:i], v.Cell[i+1:]...)
	return x
}

func ValTake(v *Val, i int) *Val {
	x := v.Cell[i]
	v.Cell = nil
	return x
}

func ExprToString(v *Val, open, close string) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(open)
	for i := 0; i < v.Count(); i++ {
		buf.WriteString(v.Cell[i].String())
		if i != v.Count()-1 {
			buf.WriteString(" ")
		}
	}
	buf.WriteString(close)
	return buf.String()
}

func (v *Val) String() string {
	switch v.Type {
	case ValFun:
		if v.Builtin != nil {
			return "<function>"
		} else {
			return fmt.Sprintf("(\\ %s %s)", v.Formals.String(), v.Body.String())
		}
	case ValNum:
		return fmt.Sprintf("%d", v.Num)
	case ValErr:
		return fmt.Sprintf("Error: %s", v.Err)
	case ValSym:
		return v.Sym
	case ValSexpr:
		return ExprToString(v, "(", ")")
	case ValQexpr:
		return ExprToString(v, "{", "}")
	}
	return ""
}

func (v *Val) Eq(y *Val) bool {
	x := v
	if x.Type != y.Type {
		return false
	}
	switch x.Type {
	case ValNum:
		return x.Num == y.Num
	case ValErr:
		return x.Err == y.Err
	case ValSym:
		return x.Sym == y.Sym
	case ValFun:
		if x.Builtin != nil && y.Builtin == nil {
			return false
		}
		if x.Builtin == nil && y.Builtin != nil {
			return false
		}
		if x.Builtin != nil && y.Builtin != nil {
			return fmt.Sprintf("%s", x.Builtin) == fmt.Sprintf("%s", y.Builtin)
		}
		return x.Formals.Eq(y.Formals) && x.Body.Eq(y.Body)
	case ValQexpr:
		fallthrough
	case ValSexpr:
		if x.Count() != y.Count() {
			return false
		}
		for i := range x.Count() {
			if !x.Cell[i].Eq(y.Cell[i]) {
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

	for i := range a.Cell[0].Count() {
		if err := Assert(a, a.Cell[0].Cell[i].Type == ValSym, "Cannot define non-symbol. Got %s, Expected %s.", TypeName(a.Cell[0].Cell[i].Type), TypeName(ValSym)); err != nil {
			return err
		}
	}
	formals := ValPop(a, 0)
	body := ValPop(a, 0)

	return NewLambda(formals, body)
}

func BuiltinList(e *Env, a *Val) *Val {
	a.Type = ValQexpr
	return a
}

func Assert(args *Val, cond bool, format string, a ...any) *Val {
	if !cond {
		return NewValErr(format, a...)
	}
	return nil
}

func AssertType(fun string, args *Val, index int, expect ValType) *Val {
	return Assert(args, args.Cell[index].Type == expect, "Function '%s' passed incorrect type for argument %i. Got %s, Expected %s.", fun, index, TypeName(args.Cell[index].Type), TypeName(expect))
}

func AssertNum(fun string, args *Val, num int) *Val {
	return Assert(args, args.Count() == num, "Function '%s' passed incorrect number of arguments. Got %i, Expected %i.", fun, args.Count(), num)
}

func AssertNotEmpty(fun string, args *Val, index int) *Val {
	return Assert(args, args.Cell[index].Count() != 0, "Function '%s' passed {} for argument %i.", fun, index)
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
	x.Type = ValSexpr
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
	for _, c := range a.Cell {
		if c.Type != ValNum {
			return NewValErr("Cannot operate on non-number!")
		}
	}
	x := ValPop(a, 0)
	if op == "-" && a.Count() == 0 {
		x.Num = -x.Num
	}

	for a.Count() > 0 {
		y := ValPop(a, 0)
		switch op {
		case "+":
			x.Num += y.Num
		case "-":
			x.Num -= y.Num
		case "*":
			x.Num *= y.Num
		case "/":
			if y.Num == 0 {
				x = NewValErr("Division by zero.")
			} else {
				x.Num /= y.Num
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
	syms := a.Cell[0]
	for i := range syms.Count() {
		if err := Assert(a, syms.Cell[i].Type == ValSym, "Function '%s' cannot define non-symbol. Got %s, Expected %s.", fun, TypeName(syms.Cell[i].Type), TypeName(ValSym)); err != nil {
			return err
		}
	}
	if err := Assert(a, syms.Count() == a.Count()-1, "Function '%s' passed too many arguments for symbols. Got %d, Expected %d.", fun, syms.Count(), a.Count()-1); err != nil {
		return err
	}
	for i := range syms.Count() {
		if fun == "def" {
			e.Def(syms.Cell[i], a.Cell[i+1])
		}
		if fun == "=" {
			e.Put(syms.Cell[i], a.Cell[i+1])
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
		r = a.Cell[0].Num > a.Cell[1].Num
	}
	if op == "<" {
		r = a.Cell[0].Num < a.Cell[1].Num
	}
	if op == ">=" {
		r = a.Cell[0].Num >= a.Cell[1].Num
	}
	if op == "<=" {
		r = a.Cell[0].Num <= a.Cell[1].Num
	}
	num := NewValNum(0)
	if r {
		num.Num = 1
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
		r = a.Cell[0].Eq(a.Cell[1])
	}
	if op == "!=" {
		r = !a.Cell[0].Eq(a.Cell[1])
	}
	num := NewValNum(0)
	if r {
		num.Num = 1
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
	a.Cell[1].Type = ValSexpr
	a.Cell[2].Type = ValSexpr

	var x *Val
	if a.Cell[0].Num != 0 {
		x = ValEval(e, ValPop(a, 1))
	} else {
		x = ValEval(e, ValPop(a, 2))
	}
	return x
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
}

func ValCall(e *Env, f *Val, a *Val) *Val {
	if f.Builtin != nil {
		return f.Builtin(e, a)
	}
	given := a.Count()
	total := f.Formals.Count()

	for a.Count() > 0 {
		if f.Formals.Count() == 0 {
			return NewValErr("Function passed too many arguments. Got %d, Expected %d.", given, total)
		}
		sym := ValPop(f.Formals, 0)
		if sym.Sym == "&" {
			if f.Formals.Count() != 1 {
				return NewValErr("Function format invalid. Symbol '&' not followed by single symbol.")
			}
			nsym := ValPop(f.Formals, 0)
			f.Env.Put(nsym, BuiltinList(e, a))
			break
		}
		val := ValPop(a, 0)
		f.Env.Put(sym, val)
	}
	if f.Formals.Count() > 0 && f.Formals.Cell[0].Sym == "&" {
		if f.Formals.Count() != 2 {
			return NewValErr("Function format invalid. Symbol '&' not followed by single symbol.")
		}
		ValPop(f.Formals, 0)
		sym := ValPop(f.Formals, 0)
		val := NewValQexpr()
		f.Env.Put(sym, val)
	}
	if f.Formals.Count() == 0 {
		f.Env.par = e
		return BuiltinEval(f.Env, ValAdd(NewValSexpr(), ValCopy(f.Body)))
	} else {
		return ValCopy(f)
	}
}

func ValEval(e *Env, v *Val) *Val {
	if v.Type == ValSym {
		x := e.Get(v)
		return x
	}
	if v.Type == ValSexpr {
		return ValEvalSexpr(e, v)
	}
	return v
}

func ValEvalSexpr(e *Env, v *Val) *Val {
	for i := range v.Count() {
		v.Cell[i] = ValEval(e, v.Cell[i])
	}

	for i := range v.Count() {
		if v.Cell[i].Type == ValErr {
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
	if f.Type != ValFun {
		err := NewValErr("S-Expression starts with incorrect type. Got %s, Expected %s.", TypeName(f.Type), TypeName(ValFun))
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

func ValRead(t antlr.Tree) *Val {
	var x *Val
	switch t2 := t.(type) {
	case *NumberContext:
		return ValReadNum(t2)
	case *SymbolContext:
		return NewValSym(t2.GetText())
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
	}
	for i := range t.GetChildCount() {
		v := ValRead(t.GetChild(i))
		if v != nil {
			ValAdd(x, v)
		}
	}
	return x
}
