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
		"", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "NUMBER", "SYMBOL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "NUMBER", "SYMBOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 61, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 3, 4, 25, 8, 4, 1, 4, 4, 4, 28, 8, 4, 11, 4, 12, 4, 29,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 53, 8, 5, 1,
		6, 4, 6, 56, 8, 6, 11, 6, 12, 6, 57, 1, 6, 1, 6, 0, 0, 7, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 1, 0, 3, 1, 0, 48, 57, 3, 0, 42, 43, 45, 45,
		47, 47, 3, 0, 9, 10, 13, 13, 32, 32, 68, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 1, 15, 1, 0, 0, 0, 3, 17, 1, 0, 0, 0, 5, 19, 1,
		0, 0, 0, 7, 21, 1, 0, 0, 0, 9, 24, 1, 0, 0, 0, 11, 52, 1, 0, 0, 0, 13,
		55, 1, 0, 0, 0, 15, 16, 5, 40, 0, 0, 16, 2, 1, 0, 0, 0, 17, 18, 5, 41,
		0, 0, 18, 4, 1, 0, 0, 0, 19, 20, 5, 123, 0, 0, 20, 6, 1, 0, 0, 0, 21, 22,
		5, 125, 0, 0, 22, 8, 1, 0, 0, 0, 23, 25, 5, 45, 0, 0, 24, 23, 1, 0, 0,
		0, 24, 25, 1, 0, 0, 0, 25, 27, 1, 0, 0, 0, 26, 28, 7, 0, 0, 0, 27, 26,
		1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0,
		30, 10, 1, 0, 0, 0, 31, 32, 5, 108, 0, 0, 32, 33, 5, 105, 0, 0, 33, 34,
		5, 115, 0, 0, 34, 53, 5, 116, 0, 0, 35, 36, 5, 104, 0, 0, 36, 37, 5, 101,
		0, 0, 37, 38, 5, 97, 0, 0, 38, 53, 5, 100, 0, 0, 39, 40, 5, 116, 0, 0,
		40, 41, 5, 97, 0, 0, 41, 42, 5, 105, 0, 0, 42, 53, 5, 108, 0, 0, 43, 44,
		5, 101, 0, 0, 44, 45, 5, 118, 0, 0, 45, 46, 5, 97, 0, 0, 46, 53, 5, 108,
		0, 0, 47, 48, 5, 106, 0, 0, 48, 49, 5, 111, 0, 0, 49, 50, 5, 105, 0, 0,
		50, 53, 5, 110, 0, 0, 51, 53, 7, 1, 0, 0, 52, 31, 1, 0, 0, 0, 52, 35, 1,
		0, 0, 0, 52, 39, 1, 0, 0, 0, 52, 43, 1, 0, 0, 0, 52, 47, 1, 0, 0, 0, 52,
		51, 1, 0, 0, 0, 53, 12, 1, 0, 0, 0, 54, 56, 7, 2, 0, 0, 55, 54, 1, 0, 0,
		0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59,
		1, 0, 0, 0, 59, 60, 6, 6, 0, 0, 60, 14, 1, 0, 0, 0, 5, 0, 24, 29, 52, 57,
		1, 6, 0, 0,
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
	LispyLexerT__0   = 1
	LispyLexerT__1   = 2
	LispyLexerT__2   = 3
	LispyLexerT__3   = 4
	LispyLexerNUMBER = 5
	LispyLexerSYMBOL = 6
	LispyLexerWS     = 7
)
