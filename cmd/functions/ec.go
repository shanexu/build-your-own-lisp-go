package main

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

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
