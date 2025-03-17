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
	fmt.Println("Lispy Version 0.0.0.0.3")
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

func EvalOp(x int, op string, y int) int {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	}
	return 0
}

func Eval(t antlr.Tree, parser *LispyParser) int {
	name := antlr.TreesGetNodeText(t, nil, parser)
	if name == "number" {
		v, _ := strconv.Atoi(t.(antlr.RuleNode).GetText())
		return v
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

	return 0
}
