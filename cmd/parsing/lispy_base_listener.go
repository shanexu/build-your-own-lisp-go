// Code generated from Lispy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // Lispy
import "github.com/antlr4-go/antlr/v4"

// BaseLispyListener is a complete listener for a parse tree produced by LispyParser.
type BaseLispyListener struct{}

var _ LispyListener = &BaseLispyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLispyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLispyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLispyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLispyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseLispyListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseLispyListener) ExitExpr(ctx *ExprContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLispyListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLispyListener) ExitProgram(ctx *ProgramContext) {}
