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
	fmt.Println("Lispy Version 0.0.0.0.5")
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

const (
	ValErr = iota
	ValNum
	ValSym
	ValSexpr
)

type Val struct {
	Type  int
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

func ValTake(v *Val, i int) *Val {
	return v.Cell[i]
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
	}
	return ""
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
	result := BuiltinOp(v, f.Sym)
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
