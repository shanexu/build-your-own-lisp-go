// Code generated from Lispy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LispyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LispyLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lispylexerLexerInit() {
	staticData := &LispyLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "NUMBER", "OPERATOR", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "NUMBER", "OPERATOR", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 32, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 3, 2, 17, 8, 2, 1, 2, 4, 2, 20,
		8, 2, 11, 2, 12, 2, 21, 1, 3, 1, 3, 1, 4, 4, 4, 27, 8, 4, 11, 4, 12, 4,
		28, 1, 4, 1, 4, 0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0, 3, 1, 0, 48,
		57, 3, 0, 42, 43, 45, 45, 47, 47, 3, 0, 9, 10, 13, 13, 32, 32, 34, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 13, 1, 0, 0, 0, 5, 16, 1, 0, 0, 0, 7,
		23, 1, 0, 0, 0, 9, 26, 1, 0, 0, 0, 11, 12, 5, 40, 0, 0, 12, 2, 1, 0, 0,
		0, 13, 14, 5, 41, 0, 0, 14, 4, 1, 0, 0, 0, 15, 17, 5, 45, 0, 0, 16, 15,
		1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 19, 1, 0, 0, 0, 18, 20, 7, 0, 0, 0,
		19, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1,
		0, 0, 0, 22, 6, 1, 0, 0, 0, 23, 24, 7, 1, 0, 0, 24, 8, 1, 0, 0, 0, 25,
		27, 7, 2, 0, 0, 26, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 26, 1, 0, 0,
		0, 28, 29, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 6, 4, 0, 0, 31, 10,
		1, 0, 0, 0, 4, 0, 16, 21, 28, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LispyLexerInit initializes any static state used to implement LispyLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLispyLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LispyLexerInit() {
	staticData := &LispyLexerLexerStaticData
	staticData.once.Do(lispylexerLexerInit)
}

// NewLispyLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLispyLexer(input antlr.CharStream) *LispyLexer {
	LispyLexerInit()
	l := new(LispyLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LispyLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lispy.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LispyLexer tokens.
const (
	LispyLexerT__0     = 1
	LispyLexerT__1     = 2
	LispyLexerNUMBER   = 3
	LispyLexerOPERATOR = 4
	LispyLexerWS       = 5
)
