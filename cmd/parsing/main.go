package main

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chzyer/readline"
)

func main() {
	fmt.Println("Lispy Version 0.0.0.0.1")
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

		tree := parser.Program()
		if len(ec.Errors) > 0 {
			fmt.Println(ec)
		} else {
			PrintParseTree(tree, parser)
		}
	}
}

type ErrorCollector struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func (ec *ErrorCollector) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("line %d:%d - %s", line, column, msg)
	ec.Errors = append(ec.Errors, errorMsg)
}

func (ec *ErrorCollector) String() string {
	return strings.Join(ec.Errors, "\n")
}

//func (c *ErrorCollector) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//	// 处理歧义错误...
//}

func PrintParseTree(tree antlr.ParseTree, parser *LispyParser) {
	printTree(tree, "", 0, parser)
}

func printTree(tree antlr.ParseTree, indent string, depth int, parser *LispyParser) {
	// 获取节点基本信息
	nodeText := antlr.TreesGetNodeText(tree, parser.RuleNames, nil)
	children := tree.GetChildren()

	// 生成节点前缀
	prefix := ""
	if depth > 0 {
		prefix += strings.Repeat("│   ", depth-1)
		prefix += "├── "
	}

	// 打印当前节点
	fmt.Printf("%s%s [%T]\n", prefix, nodeText, tree)

	// 递归打印子节点
	for i, child := range children {
		lastNode := i == len(children)-1
		newIndent := indent
		if depth > 0 {
			if lastNode {
				newIndent += "    "
			} else {
				newIndent += "│   "
			}
		}
		printTree(child.(antlr.ParseTree), newIndent, depth+1, parser)
	}
}
