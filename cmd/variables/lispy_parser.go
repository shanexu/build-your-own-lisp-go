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
		"", "", "", "", "", "NUMBER", "SYMBOL", "WS",
	}
	staticData.RuleNames = []string{
		"number", "symbol", "sexpr", "qexpr", "expr", "lispy",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 7, 47, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 19, 8, 2, 10, 2,
		12, 2, 22, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 5, 3, 28, 8, 3, 10, 3, 12, 3,
		31, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 39, 8, 4, 1, 5, 5,
		5, 42, 8, 5, 10, 5, 12, 5, 45, 9, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10,
		0, 0, 46, 0, 12, 1, 0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 16, 1, 0, 0, 0, 6, 25,
		1, 0, 0, 0, 8, 38, 1, 0, 0, 0, 10, 43, 1, 0, 0, 0, 12, 13, 5, 5, 0, 0,
		13, 1, 1, 0, 0, 0, 14, 15, 5, 6, 0, 0, 15, 3, 1, 0, 0, 0, 16, 20, 5, 1,
		0, 0, 17, 19, 3, 8, 4, 0, 18, 17, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18,
		1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 23, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0,
		23, 24, 5, 2, 0, 0, 24, 5, 1, 0, 0, 0, 25, 29, 5, 3, 0, 0, 26, 28, 3, 8,
		4, 0, 27, 26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30,
		1, 0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 4, 0, 0,
		33, 7, 1, 0, 0, 0, 34, 39, 3, 0, 0, 0, 35, 39, 3, 2, 1, 0, 36, 39, 3, 4,
		2, 0, 37, 39, 3, 6, 3, 0, 38, 34, 1, 0, 0, 0, 38, 35, 1, 0, 0, 0, 38, 36,
		1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 9, 1, 0, 0, 0, 40, 42, 3, 8, 4, 0,
		41, 40, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1,
		0, 0, 0, 44, 11, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 4, 20, 29, 38, 43,
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
	LispyParserEOF    = antlr.TokenEOF
	LispyParserT__0   = 1
	LispyParserT__1   = 2
	LispyParserT__2   = 3
	LispyParserT__3   = 4
	LispyParserNUMBER = 5
	LispyParserSYMBOL = 6
	LispyParserWS     = 7
)

// LispyParser rules.
const (
	LispyParserRULE_number = 0
	LispyParserRULE_symbol = 1
	LispyParserRULE_sexpr  = 2
	LispyParserRULE_qexpr  = 3
	LispyParserRULE_expr   = 4
	LispyParserRULE_lispy  = 5
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
		p.SetState(12)
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
		p.SetState(14)
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
	p.EnterRule(localctx, 4, LispyParserRULE_sexpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(LispyParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&106) != 0 {
		{
			p.SetState(17)
			p.Expr()
		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
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
	p.EnterRule(localctx, 6, LispyParserRULE_qexpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(LispyParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&106) != 0 {
		{
			p.SetState(26)
			p.Expr()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
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
	p.EnterRule(localctx, 8, LispyParserRULE_expr)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LispyParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Number()
		}

	case LispyParserSYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Symbol()
		}

	case LispyParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Sexpr()
		}

	case LispyParserT__2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
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
	p.EnterRule(localctx, 10, LispyParserRULE_lispy)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&106) != 0 {
		{
			p.SetState(40)
			p.Expr()
		}

		p.SetState(45)
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
