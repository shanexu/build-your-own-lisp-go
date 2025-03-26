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
	fmt.Println("Lispy Version 0.0.0.0.7")
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
		return v
	}
	return NewValErr("Unbound Symbol '%s'", k.Sym)
}

func (env *Env) Put(k *Val, v *Val) {
	env.vals[k.Sym] = ValCopy(v)
}

type Val struct {
	Type ValType
	Num  int
	Err  string
	Sym  string
	fun  BuiltinFunc
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
	return &Val{Type: ValFun, fun: f}
}

func NewValSexpr() *Val { return &Val{Type: ValSexpr} }

func NewValQexpr() *Val { return &Val{Type: ValQexpr} }

func ValCopy(v *Val) *Val {
	x := &Val{Type: v.Type}
	switch v.Type {
	case ValFun:
		x.fun = v.fun
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
		return "<function>"
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
	v.Cell = v.Cell[1:]
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

func BuiltinDef(e *Env, a *Val) *Val {
	if err := AssertType("def", a, 0, ValQexpr); err != nil {
		return err
	}
	syms := a.Cell[0]
	for i := range syms.Count() {
		if err := Assert(a, syms.Cell[i].Type == ValSym, "Function 'def' cannot define non-symbol. Got %s, Expected %s.", TypeName(syms.Cell[i].Type), TypeName(ValSym)); err != nil {
			return err
		}
	}
	if err := Assert(a, syms.Count() == a.Count()-1, "Function 'def' passed too many arguments for symbols. Got %d, Expected %d.", syms.Count(), a.Count()); err != nil {
		return err
	}
	for i := range syms.Count() {
		e.Put(syms.Cell[i], a.Cell[i+1])
	}
	return NewValSexpr()
}

func (env *Env) AddBuiltin(name string, builtinFunc BuiltinFunc) {
	k := NewValSym(name)
	v := NewValFun(builtinFunc)
	env.Put(k, v)
}

func (env *Env) AddBuiltins() {
	/* Variable Functions */
	env.AddBuiltin("def", BuiltinDef)

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
		return ValTake(v, 0)
	}

	/* Ensure first element is a function after evaluation */
	f := ValPop(v, 0)
	if f.Type != ValFun {
		err := NewValErr("S-Expression starts with incorrect type. Got %s, Expected %s.", TypeName(f.Type), TypeName(ValFun))
		return err
	}
	return f.fun(e, v)
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
