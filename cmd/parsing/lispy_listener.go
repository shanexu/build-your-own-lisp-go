// Code generated from Lispy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // Lispy
import "github.com/antlr4-go/antlr/v4"

// LispyListener is a complete listener for a parse tree produced by LispyParser.
type LispyListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)
}
