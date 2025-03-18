// Code generated from Lispy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // Lispy
import "github.com/antlr4-go/antlr/v4"

// LispyListener is a complete listener for a parse tree produced by LispyParser.
type LispyListener interface {
	antlr.ParseTreeListener

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterSexpr is called when entering the sexpr production.
	EnterSexpr(c *SexprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterLispy is called when entering the lispy production.
	EnterLispy(c *LispyContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitSexpr is called when exiting the sexpr production.
	ExitSexpr(c *SexprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitLispy is called when exiting the lispy production.
	ExitLispy(c *LispyContext)
}
