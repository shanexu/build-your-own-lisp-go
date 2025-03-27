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

// EnterNumber is called when production number is entered.
func (s *BaseLispyListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseLispyListener) ExitNumber(ctx *NumberContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseLispyListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseLispyListener) ExitSymbol(ctx *SymbolContext) {}

// EnterString is called when production string is entered.
func (s *BaseLispyListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseLispyListener) ExitString(ctx *StringContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseLispyListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseLispyListener) ExitComment(ctx *CommentContext) {}

// EnterSexpr is called when production sexpr is entered.
func (s *BaseLispyListener) EnterSexpr(ctx *SexprContext) {}

// ExitSexpr is called when production sexpr is exited.
func (s *BaseLispyListener) ExitSexpr(ctx *SexprContext) {}

// EnterQexpr is called when production qexpr is entered.
func (s *BaseLispyListener) EnterQexpr(ctx *QexprContext) {}

// ExitQexpr is called when production qexpr is exited.
func (s *BaseLispyListener) ExitQexpr(ctx *QexprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseLispyListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseLispyListener) ExitExpr(ctx *ExprContext) {}

// EnterLispy is called when production lispy is entered.
func (s *BaseLispyListener) EnterLispy(ctx *LispyContext) {}

// ExitLispy is called when production lispy is exited.
func (s *BaseLispyListener) ExitLispy(ctx *LispyContext) {}
