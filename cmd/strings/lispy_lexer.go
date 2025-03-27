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
		"", "", "", "", "", "NUMBER", "SYMBOL", "STRING", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "NUMBER", "SYMBOL", "STRING", "COMMENT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 65, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 3, 4, 29, 8, 4, 1, 4, 4, 4, 32,
		8, 4, 11, 4, 12, 4, 33, 1, 5, 4, 5, 37, 8, 5, 11, 5, 12, 5, 38, 1, 6, 1,
		6, 1, 6, 1, 6, 5, 6, 45, 8, 6, 10, 6, 12, 6, 48, 9, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 5, 7, 54, 8, 7, 10, 7, 12, 7, 57, 9, 7, 1, 8, 4, 8, 60, 8, 8, 11,
		8, 12, 8, 61, 1, 8, 1, 8, 1, 46, 0, 9, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 1, 0, 4, 1, 0, 48, 57, 10, 0, 33, 33, 38, 38, 42,
		43, 45, 45, 47, 57, 60, 62, 65, 90, 92, 92, 95, 95, 97, 122, 2, 0, 10,
		10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 71, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1, 19,
		1, 0, 0, 0, 3, 21, 1, 0, 0, 0, 5, 23, 1, 0, 0, 0, 7, 25, 1, 0, 0, 0, 9,
		28, 1, 0, 0, 0, 11, 36, 1, 0, 0, 0, 13, 40, 1, 0, 0, 0, 15, 51, 1, 0, 0,
		0, 17, 59, 1, 0, 0, 0, 19, 20, 5, 40, 0, 0, 20, 2, 1, 0, 0, 0, 21, 22,
		5, 41, 0, 0, 22, 4, 1, 0, 0, 0, 23, 24, 5, 123, 0, 0, 24, 6, 1, 0, 0, 0,
		25, 26, 5, 125, 0, 0, 26, 8, 1, 0, 0, 0, 27, 29, 5, 45, 0, 0, 28, 27, 1,
		0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 31, 1, 0, 0, 0, 30, 32, 7, 0, 0, 0, 31,
		30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0,
		0, 34, 10, 1, 0, 0, 0, 35, 37, 7, 1, 0, 0, 36, 35, 1, 0, 0, 0, 37, 38,
		1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 12, 1, 0, 0, 0,
		40, 46, 5, 34, 0, 0, 41, 42, 5, 92, 0, 0, 42, 45, 5, 34, 0, 0, 43, 45,
		9, 0, 0, 0, 44, 41, 1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 46, 1,
		0, 0, 0, 49, 50, 5, 34, 0, 0, 50, 14, 1, 0, 0, 0, 51, 55, 5, 59, 0, 0,
		52, 54, 8, 2, 0, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1,
		0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 16, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58,
		60, 7, 3, 0, 0, 59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0,
		0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 6, 8, 0, 0, 64, 18,
		1, 0, 0, 0, 8, 0, 28, 33, 38, 44, 46, 55, 61, 1, 6, 0, 0,
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
	LispyLexerT__0    = 1
	LispyLexerT__1    = 2
	LispyLexerT__2    = 3
	LispyLexerT__3    = 4
	LispyLexerNUMBER  = 5
	LispyLexerSYMBOL  = 6
	LispyLexerSTRING  = 7
	LispyLexerCOMMENT = 8
	LispyLexerWS      = 9
)
