// Code generated from Lispy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // Lispy
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LispyParser struct {
	*antlr.BaseParser
}

var LispyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lispyParserInit() {
	staticData := &LispyParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "NUMBER", "SYMBOL", "STRING", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"number", "symbol", "string", "comment", "sexpr", "qexpr", "expr", "lispy",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 57, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 5, 4, 27, 8, 4, 10, 4, 12, 4, 30, 9, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 5, 5, 36, 8, 5, 10, 5, 12, 5, 39, 9, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 49, 8, 6, 1, 7, 5, 7, 52, 8, 7, 10,
		7, 12, 7, 55, 9, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 0, 56,
		0, 16, 1, 0, 0, 0, 2, 18, 1, 0, 0, 0, 4, 20, 1, 0, 0, 0, 6, 22, 1, 0, 0,
		0, 8, 24, 1, 0, 0, 0, 10, 33, 1, 0, 0, 0, 12, 48, 1, 0, 0, 0, 14, 53, 1,
		0, 0, 0, 16, 17, 5, 5, 0, 0, 17, 1, 1, 0, 0, 0, 18, 19, 5, 6, 0, 0, 19,
		3, 1, 0, 0, 0, 20, 21, 5, 7, 0, 0, 21, 5, 1, 0, 0, 0, 22, 23, 5, 8, 0,
		0, 23, 7, 1, 0, 0, 0, 24, 28, 5, 1, 0, 0, 25, 27, 3, 12, 6, 0, 26, 25,
		1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0,
		29, 31, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 5, 2, 0, 0, 32, 9, 1, 0,
		0, 0, 33, 37, 5, 3, 0, 0, 34, 36, 3, 12, 6, 0, 35, 34, 1, 0, 0, 0, 36,
		39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0,
		0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 4, 0, 0, 41, 11, 1, 0, 0, 0, 42, 49,
		3, 0, 0, 0, 43, 49, 3, 2, 1, 0, 44, 49, 3, 4, 2, 0, 45, 49, 3, 6, 3, 0,
		46, 49, 3, 8, 4, 0, 47, 49, 3, 10, 5, 0, 48, 42, 1, 0, 0, 0, 48, 43, 1,
		0, 0, 0, 48, 44, 1, 0, 0, 0, 48, 45, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48,
		47, 1, 0, 0, 0, 49, 13, 1, 0, 0, 0, 50, 52, 3, 12, 6, 0, 51, 50, 1, 0,
		0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 15,
		1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 4, 28, 37, 48, 53,
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

// LispyParserInit initializes any static state used to implement LispyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLispyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LispyParserInit() {
	staticData := &LispyParserStaticData
	staticData.once.Do(lispyParserInit)
}

// NewLispyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLispyParser(input antlr.TokenStream) *LispyParser {
	LispyParserInit()
	this := new(LispyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LispyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lispy.g4"

	return this
}

// LispyParser tokens.
const (
	LispyParserEOF     = antlr.TokenEOF
	LispyParserT__0    = 1
	LispyParserT__1    = 2
	LispyParserT__2    = 3
	LispyParserT__3    = 4
	LispyParserNUMBER  = 5
	LispyParserSYMBOL  = 6
	LispyParserSTRING  = 7
	LispyParserCOMMENT = 8
	LispyParserWS      = 9
)

// LispyParser rules.
const (
	LispyParserRULE_number  = 0
	LispyParserRULE_symbol  = 1
	LispyParserRULE_string  = 2
	LispyParserRULE_comment = 3
	LispyParserRULE_sexpr   = 4
	LispyParserRULE_qexpr   = 5
	LispyParserRULE_expr    = 6
	LispyParserRULE_lispy   = 7
)

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LispyParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LispyParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *LispyParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LispyParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(LispyParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYMBOL() antlr.TerminalNode

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_symbol
	return p
}

func InitEmptySymbolContext(p *SymbolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_symbol
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LispyParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(LispyParserSYMBOL, 0)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *LispyParser) Symbol() (localctx ISymbolContext) {
	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LispyParserRULE_symbol)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(LispyParserSYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LispyParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(LispyParserSTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *LispyParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LispyParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(LispyParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LispyParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(LispyParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *LispyParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LispyParserRULE_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(LispyParserCOMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISexprContext is an interface to support dynamic dispatch.
type ISexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsSexprContext differentiates from other interfaces.
	IsSexprContext()
}

type SexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySexprContext() *SexprContext {
	var p = new(SexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_sexpr
	return p
}

func InitEmptySexprContext(p *SexprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_sexpr
}

func (*SexprContext) IsSexprContext() {}

func NewSexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexprContext {
	var p = new(SexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LispyParserRULE_sexpr

	return p
}

func (s *SexprContext) GetParser() antlr.Parser { return s.parser }

func (s *SexprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SexprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.EnterSexpr(s)
	}
}

func (s *SexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.ExitSexpr(s)
	}
}

func (p *LispyParser) Sexpr() (localctx ISexprContext) {
	localctx = NewSexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LispyParserRULE_sexpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(LispyParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&490) != 0 {
		{
			p.SetState(25)
			p.Expr()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(LispyParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQexprContext is an interface to support dynamic dispatch.
type IQexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsQexprContext differentiates from other interfaces.
	IsQexprContext()
}

type QexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQexprContext() *QexprContext {
	var p = new(QexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_qexpr
	return p
}

func InitEmptyQexprContext(p *QexprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_qexpr
}

func (*QexprContext) IsQexprContext() {}

func NewQexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QexprContext {
	var p = new(QexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LispyParserRULE_qexpr

	return p
}

func (s *QexprContext) GetParser() antlr.Parser { return s.parser }

func (s *QexprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *QexprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *QexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.EnterQexpr(s)
	}
}

func (s *QexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.ExitQexpr(s)
	}
}

func (p *LispyParser) Qexpr() (localctx IQexprContext) {
	localctx = NewQexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LispyParserRULE_qexpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(LispyParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&490) != 0 {
		{
			p.SetState(34)
			p.Expr()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(LispyParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	Symbol() ISymbolContext
	String_() IStringContext
	Comment() ICommentContext
	Sexpr() ISexprContext
	Qexpr() IQexprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LispyParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ExprContext) Symbol() ISymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *ExprContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *ExprContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *ExprContext) Sexpr() ISexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISexprContext)
}

func (s *ExprContext) Qexpr() IQexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQexprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *LispyParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LispyParserRULE_expr)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LispyParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Number()
		}

	case LispyParserSYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Symbol()
		}

	case LispyParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.String_()
		}

	case LispyParserCOMMENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.Comment()
		}

	case LispyParserT__0:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(46)
			p.Sexpr()
		}

	case LispyParserT__2:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(47)
			p.Qexpr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILispyContext is an interface to support dynamic dispatch.
type ILispyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsLispyContext differentiates from other interfaces.
	IsLispyContext()
}

type LispyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLispyContext() *LispyContext {
	var p = new(LispyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_lispy
	return p
}

func InitEmptyLispyContext(p *LispyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LispyParserRULE_lispy
}

func (*LispyContext) IsLispyContext() {}

func NewLispyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LispyContext {
	var p = new(LispyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LispyParserRULE_lispy

	return p
}

func (s *LispyContext) GetParser() antlr.Parser { return s.parser }

func (s *LispyContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LispyContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LispyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LispyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LispyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.EnterLispy(s)
	}
}

func (s *LispyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LispyListener); ok {
		listenerT.ExitLispy(s)
	}
}

func (p *LispyParser) Lispy() (localctx ILispyContext) {
	localctx = NewLispyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LispyParserRULE_lispy)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&490) != 0 {
		{
			p.SetState(50)
			p.Expr()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
