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
	fmt.Println("Lispy Version 0.0.0.0.6")
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
			x := ValEval(ValRead(tree))
			fmt.Println(x)
		}
	}
}

type ValType int

const (
	ValErr ValType = iota
	ValNum
	ValSym
	ValSexpr
	ValQexpr
)

type Val struct {
	Type  ValType
	Num   int
	Err   string
	Sym   string
	Count int
	Cell  []*Val
}

func NewValNum(x int) *Val {
	return &Val{Type: ValNum, Num: x}
}

func NewValErr(m string) *Val {
	return &Val{Type: ValErr, Err: m}
}

func NewValSym(s string) *Val { return &Val{Type: ValSym, Sym: s} }

func NewValSexpr() *Val { return &Val{Type: ValSexpr} }

func NewValQexpr() *Val { return &Val{Type: ValQexpr} }

func ValAdd(v *Val, x *Val) *Val {
	v.Count++
	v.Cell = append(v.Cell, x)
	return v
}

func ValPop(v *Val, i int) *Val {
	x := v.Cell[i]
	v.Count--
	v.Cell = append(v.Cell[:i], v.Cell[i+1:]...)
	return x
}

func ValJoin(x *Val, y *Val) *Val {
	x.Count += y.Count
	x.Cell = append(x.Cell, y.Cell...)
	y.Count = 0
	y.Cell = nil
	return x
}

func ValTake(v *Val, i int) *Val {
	x := v.Cell[i]
	v.Count = 0
	v.Cell = nil
	return x
}

func ExprToString(v *Val, open, close string) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(open)
	for i := 0; i < v.Count; i++ {
		buf.WriteString(v.Cell[i].String())
		if i != v.Count-1 {
			buf.WriteString(" ")
		}
	}
	buf.WriteString(close)
	return buf.String()
}

func (v *Val) String() string {
	switch v.Type {
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

func BuiltinList(a *Val) *Val {
	a.Type = ValQexpr
	return a
}

func Assert(cond bool, err string) *Val {
	if !cond {
		return NewValErr(err)
	}
	return nil
}

func BuiltinHead(a *Val) *Val {
	if e := Assert(a.Count == 1, "Function 'head' passed too many arguments."); e != nil {
		return e
	}
	if e := Assert(a.Cell[0].Type == ValQexpr, "Function 'head' passed incorrect type."); e != nil {
		return e
	}
	if e := Assert(a.Cell[0].Count != 0, "Function 'head' passed {}."); e != nil {
		return e
	}
	v := ValTake(a, 0)
	v.Count = 1
	v.Cell = v.Cell[1:]
	return v
}

func BuiltinTail(a *Val) *Val {
	if e := Assert(a.Count == 1, "Function 'tail' passed too many arguments."); e != nil {
		return e
	}
	if e := Assert(a.Cell[0].Type == ValQexpr, "Function 'tail' passed incorrect type."); e != nil {
		return e
	}
	if e := Assert(a.Cell[0].Count != 0, "Function 'tail' passed {}."); e != nil {
		return e
	}
	v := ValTake(a, 0)
	ValPop(v, 0)
	return v
}

func BuiltinEval(a *Val) *Val {
	if e := Assert(a.Count == 1, "Function 'eval' passed too many arguments."); e != nil {
		return e
	}
	if e := Assert(a.Cell[0].Type == ValQexpr, "Function 'eval' passed incorrect type."); e != nil {
		return e
	}
	x := ValTake(a, 0)
	x.Type = ValSexpr
	return ValEval(x)
}

func BuiltinJoin(a *Val) *Val {
	for i := range a.Count {
		if e := Assert(a.Cell[i].Type == ValQexpr, "Function 'join' passed incorrect type."); e != nil {
			return e
		}
	}
	x := ValPop(a, 0)
	for a.Count > 0 {
		x = ValJoin(x, ValPop(a, 0))
	}
	return x
}

func BuiltinOp(a *Val, op string) *Val {
	for _, c := range a.Cell {
		if c.Type != ValNum {
			return NewValErr("Cannot operate on non-number!")
		}
	}
	x := ValPop(a, 0)
	if op == "-" && a.Count == 0 {
		x.Num = -x.Num
	}

	for a.Count > 0 {
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

func Builtin(a *Val, fun string) *Val {
	switch fun {
	case "list":
		return BuiltinList(a)
	case "head":
		return BuiltinHead(a)
	case "tail":
		return BuiltinTail(a)
	case "join":
		return BuiltinJoin(a)
	case "eval":
		return BuiltinEval(a)
	case "+":
		return BuiltinOp(a, "+")
	case "-":
		return BuiltinOp(a, "-")
	case "*":
		return BuiltinOp(a, "*")
	case "/":
		return BuiltinOp(a, "/")
	default:
		return NewValErr("Unknown Function!")
	}
}

func ValEval(v *Val) *Val {
	if v.Type == ValSexpr {
		return ValEvalSexpr(v)
	}
	return v
}

func ValEvalSexpr(v *Val) *Val {
	for i := range v.Count {
		v.Cell[i] = ValEval(v.Cell[i])
	}

	for i := range v.Count {
		if v.Cell[i].Type == ValErr {
			return ValTake(v, i)
		}
	}

	if v.Count == 0 {
		return v
	}

	if v.Count == 1 {
		return ValTake(v, 0)
	}

	f := ValPop(v, 0)
	if f.Type != ValSym {
		return NewValErr("S-expression Does not start with symbol.")
	}
	result := Builtin(v, f.Sym)
	return result
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
