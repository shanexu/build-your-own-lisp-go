package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chzyer/readline"
)

func main() {
	fmt.Println("Lispy Version 0.0.0.0.4")
	fmt.Println("Press Ctrl+c to Exit\n")

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
			fmt.Println(Eval(tree, parser))
		}
	}
}

const (
	ErrUnknown = iota
	ErrDivZero
	ErrBadOp
	ErrBadNum
)

const (
	ValNum = iota
	ValErr
)

type Val struct {
	Type int
	Num  int
	Err  int
}

func NewValNum(x int) Val {
	return Val{Type: ValNum, Num: x}
}

func NewValErr(x int) Val {
	return Val{Type: ValErr, Err: x}
}

func (v Val) String() string {
	switch v.Type {
	case ValNum:
		return fmt.Sprintf("%d", v.Num)
	case ValErr:
		switch v.Err {
		case ErrDivZero:
			return "Error: Division By Zero!"
		case ErrBadOp:
			return "Error: Invalid Operator!"
		case ErrBadNum:
			return "Error: Invalid Number!"
		case ErrUnknown:
			return "Error: Unknown!"
		}
	}
	return ""
}

func EvalOp(x Val, op string, y Val) Val {
	if x.Type == ValErr {
		return x
	}
	if y.Type == ValErr {
		return y
	}

	switch op {
	case "+":
		return NewValNum(x.Num + y.Num)
	case "-":
		return NewValNum(x.Num - y.Num)
	case "*":
		return NewValNum(x.Num * y.Num)
	case "/":
		if y.Num == 0 {
			return NewValErr(ErrDivZero)
		} else {
			return NewValNum(x.Num / y.Num)
		}
	}

	return NewValErr(ErrBadOp)
}

func Eval(t antlr.Tree, parser *LispyParser) Val {
	name := antlr.TreesGetNodeText(t, nil, parser)
	if name == "number" {
		v, err := strconv.Atoi(t.(antlr.RuleNode).GetText())
		if err != nil {
			return NewValErr(ErrBadNum)
		}
		return NewValNum(v)
	}

	if name == "expr" {
		if t.GetChildCount() == 1 {
			return Eval(t.GetChild(0), parser)
		}
	}

	if name == "expr" || name == "lispy" {
		op := t.GetChild(1).(antlr.RuleNode).GetText()
		x := Eval(t.GetChild(2), parser)
		for i := 3; i < t.GetChildCount(); i++ {
			c := t.GetChild(i)
			if _, ok := c.(antlr.TerminalNode); ok {
				continue
			}
			x = EvalOp(x, op, Eval(c, parser))
		}
		return x
	}

	return NewValErr(ErrUnknown)
}
