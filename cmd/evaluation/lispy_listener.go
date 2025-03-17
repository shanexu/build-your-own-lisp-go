// Code generated from Lispy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // Lispy
import "github.com/antlr4-go/antlr/v4"

// LispyListener is a complete listener for a parse tree produced by LispyParser.
type LispyListener interface {
	antlr.ParseTreeListener

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterLispy is called when entering the lispy production.
	EnterLispy(c *LispyContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitLispy is called when exiting the lispy production.
	ExitLispy(c *LispyContext)
}
