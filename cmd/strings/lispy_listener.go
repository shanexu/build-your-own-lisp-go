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

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterSexpr is called when entering the sexpr production.
	EnterSexpr(c *SexprContext)

	// EnterQexpr is called when entering the qexpr production.
	EnterQexpr(c *QexprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterLispy is called when entering the lispy production.
	EnterLispy(c *LispyContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitSexpr is called when exiting the sexpr production.
	ExitSexpr(c *SexprContext)

	// ExitQexpr is called when exiting the qexpr production.
	ExitQexpr(c *QexprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitLispy is called when exiting the lispy production.
	ExitLispy(c *LispyContext)
}
