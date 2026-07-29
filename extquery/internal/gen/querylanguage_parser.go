// Code generated from QueryLanguageParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // QueryLanguageParser
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

type QueryLanguageParser struct {
	*antlr.BaseParser
}

var QueryLanguageParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func querylanguageparserParserInit() {
	staticData := &QueryLanguageParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "'('", "')'", "','", "'='", "'!='",
		"'=*'", "'!=*'", "'~'", "'!~'", "'~*'", "'!~*'", "'>'", "'>='", "'<'",
		"'<='",
	}
	staticData.SymbolicNames = []string{
		"", "AND", "OR", "NOT", "IS", "IN", "PRESENT", "COUNT", "LPAREN", "RPAREN",
		"COMMA", "OP_EQUAL", "OP_NOT_EQUAL", "OP_EQUAL_IGNORE_CASE", "OP_NOT_EQUAL_IGNORE_CASE",
		"OP_TILDE", "OP_NOT_TILDE", "OP_TILDE_IGNORE_CASE", "OP_NOT_TILDE_IGNORE_CASE",
		"OP_GREATER_THAN", "OP_GREATER_THAN_EQUAL", "OP_LESS_THAN", "OP_LESS_THAN_EQUAL",
		"INTEGER", "QUOTED", "VARIABLE", "PLACEHOLDER", "TERM", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"topLevelQuery", "query", "clause", "isPresentClause", "inClause", "compareClause",
		"countClause", "valuelist", "value", "fieldName", "int_value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 96, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 3, 0, 24, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		44, 8, 1, 10, 1, 12, 1, 47, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 53, 8,
		2, 1, 3, 1, 3, 1, 3, 3, 3, 58, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 64,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 85, 8, 7, 10, 7,
		12, 7, 88, 9, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 0, 1, 2,
		11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 4, 1, 0, 11, 18, 2, 0, 11,
		12, 19, 22, 1, 0, 23, 27, 2, 0, 24, 24, 27, 27, 95, 0, 23, 1, 0, 0, 0,
		2, 35, 1, 0, 0, 0, 4, 52, 1, 0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 61, 1, 0, 0,
		0, 10, 70, 1, 0, 0, 0, 12, 74, 1, 0, 0, 0, 14, 81, 1, 0, 0, 0, 16, 89,
		1, 0, 0, 0, 18, 91, 1, 0, 0, 0, 20, 93, 1, 0, 0, 0, 22, 24, 3, 2, 1, 0,
		23, 22, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 5,
		0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 28, 6, 1, -1, 0, 28, 36, 3, 4, 2, 0, 29,
		30, 5, 8, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 9, 0, 0, 32, 36, 1, 0, 0,
		0, 33, 34, 5, 3, 0, 0, 34, 36, 3, 2, 1, 3, 35, 27, 1, 0, 0, 0, 35, 29,
		1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 45, 1, 0, 0, 0, 37, 38, 10, 2, 0, 0,
		38, 39, 5, 1, 0, 0, 39, 44, 3, 2, 1, 3, 40, 41, 10, 1, 0, 0, 41, 42, 5,
		2, 0, 0, 42, 44, 3, 2, 1, 2, 43, 37, 1, 0, 0, 0, 43, 40, 1, 0, 0, 0, 44,
		47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 3, 1, 0, 0,
		0, 47, 45, 1, 0, 0, 0, 48, 53, 3, 6, 3, 0, 49, 53, 3, 8, 4, 0, 50, 53,
		3, 10, 5, 0, 51, 53, 3, 12, 6, 0, 52, 48, 1, 0, 0, 0, 52, 49, 1, 0, 0,
		0, 52, 50, 1, 0, 0, 0, 52, 51, 1, 0, 0, 0, 53, 5, 1, 0, 0, 0, 54, 55, 3,
		18, 9, 0, 55, 57, 5, 4, 0, 0, 56, 58, 5, 3, 0, 0, 57, 56, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 5, 6, 0, 0, 60, 7, 1, 0, 0,
		0, 61, 63, 3, 18, 9, 0, 62, 64, 5, 3, 0, 0, 63, 62, 1, 0, 0, 0, 63, 64,
		1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 5, 5, 0, 0, 66, 67, 5, 8, 0, 0,
		67, 68, 3, 14, 7, 0, 68, 69, 5, 9, 0, 0, 69, 9, 1, 0, 0, 0, 70, 71, 3,
		18, 9, 0, 71, 72, 7, 0, 0, 0, 72, 73, 3, 16, 8, 0, 73, 11, 1, 0, 0, 0,
		74, 75, 5, 7, 0, 0, 75, 76, 5, 8, 0, 0, 76, 77, 3, 18, 9, 0, 77, 78, 5,
		9, 0, 0, 78, 79, 7, 1, 0, 0, 79, 80, 3, 20, 10, 0, 80, 13, 1, 0, 0, 0,
		81, 86, 3, 16, 8, 0, 82, 83, 5, 10, 0, 0, 83, 85, 3, 16, 8, 0, 84, 82,
		1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0,
		87, 15, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 7, 2, 0, 0, 90, 17, 1,
		0, 0, 0, 91, 92, 7, 3, 0, 0, 92, 19, 1, 0, 0, 0, 93, 94, 5, 23, 0, 0, 94,
		21, 1, 0, 0, 0, 8, 23, 35, 43, 45, 52, 57, 63, 86,
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

// QueryLanguageParserInit initializes any static state used to implement QueryLanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQueryLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLanguageParserInit() {
	staticData := &QueryLanguageParserParserStaticData
	staticData.once.Do(querylanguageparserParserInit)
}

// NewQueryLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQueryLanguageParser(input antlr.TokenStream) *QueryLanguageParser {
	QueryLanguageParserInit()
	this := new(QueryLanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &QueryLanguageParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "QueryLanguageParser.g4"

	return this
}

// QueryLanguageParser tokens.
const (
	QueryLanguageParserEOF                      = antlr.TokenEOF
	QueryLanguageParserAND                      = 1
	QueryLanguageParserOR                       = 2
	QueryLanguageParserNOT                      = 3
	QueryLanguageParserIS                       = 4
	QueryLanguageParserIN                       = 5
	QueryLanguageParserPRESENT                  = 6
	QueryLanguageParserCOUNT                    = 7
	QueryLanguageParserLPAREN                   = 8
	QueryLanguageParserRPAREN                   = 9
	QueryLanguageParserCOMMA                    = 10
	QueryLanguageParserOP_EQUAL                 = 11
	QueryLanguageParserOP_NOT_EQUAL             = 12
	QueryLanguageParserOP_EQUAL_IGNORE_CASE     = 13
	QueryLanguageParserOP_NOT_EQUAL_IGNORE_CASE = 14
	QueryLanguageParserOP_TILDE                 = 15
	QueryLanguageParserOP_NOT_TILDE             = 16
	QueryLanguageParserOP_TILDE_IGNORE_CASE     = 17
	QueryLanguageParserOP_NOT_TILDE_IGNORE_CASE = 18
	QueryLanguageParserOP_GREATER_THAN          = 19
	QueryLanguageParserOP_GREATER_THAN_EQUAL    = 20
	QueryLanguageParserOP_LESS_THAN             = 21
	QueryLanguageParserOP_LESS_THAN_EQUAL       = 22
	QueryLanguageParserINTEGER                  = 23
	QueryLanguageParserQUOTED                   = 24
	QueryLanguageParserVARIABLE                 = 25
	QueryLanguageParserPLACEHOLDER              = 26
	QueryLanguageParserTERM                     = 27
	QueryLanguageParserWHITESPACE               = 28
)

// QueryLanguageParser rules.
const (
	QueryLanguageParserRULE_topLevelQuery   = 0
	QueryLanguageParserRULE_query           = 1
	QueryLanguageParserRULE_clause          = 2
	QueryLanguageParserRULE_isPresentClause = 3
	QueryLanguageParserRULE_inClause        = 4
	QueryLanguageParserRULE_compareClause   = 5
	QueryLanguageParserRULE_countClause     = 6
	QueryLanguageParserRULE_valuelist       = 7
	QueryLanguageParserRULE_value           = 8
	QueryLanguageParserRULE_fieldName       = 9
	QueryLanguageParserRULE_int_value       = 10
)

// ITopLevelQueryContext is an interface to support dynamic dispatch.
type ITopLevelQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Query() IQueryContext

	// IsTopLevelQueryContext differentiates from other interfaces.
	IsTopLevelQueryContext()
}

type TopLevelQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelQueryContext() *TopLevelQueryContext {
	var p = new(TopLevelQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_topLevelQuery
	return p
}

func InitEmptyTopLevelQueryContext(p *TopLevelQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_topLevelQuery
}

func (*TopLevelQueryContext) IsTopLevelQueryContext() {}

func NewTopLevelQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelQueryContext {
	var p = new(TopLevelQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_topLevelQuery

	return p
}

func (s *TopLevelQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelQueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserEOF, 0)
}

func (s *TopLevelQueryContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *TopLevelQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelQueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitTopLevelQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) TopLevelQuery() (localctx ITopLevelQueryContext) {
	localctx = NewTopLevelQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryLanguageParserRULE_topLevelQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&150995336) != 0 {
		{
			p.SetState(22)
			p.query(0)
		}

	}
	{
		p.SetState(25)
		p.Match(QueryLanguageParserEOF)
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

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Clause() IClauseContext
	LPAREN() antlr.TerminalNode
	AllQuery() []IQueryContext
	Query(i int) IQueryContext
	RPAREN() antlr.TerminalNode
	NOT() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Clause() IClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *QueryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserLPAREN, 0)
}

func (s *QueryContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *QueryContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
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

	return t.(IQueryContext)
}

func (s *QueryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserRPAREN, 0)
}

func (s *QueryContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserNOT, 0)
}

func (s *QueryContext) AND() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserAND, 0)
}

func (s *QueryContext) OR() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOR, 0)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *QueryLanguageParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, QueryLanguageParserRULE_query, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case QueryLanguageParserCOUNT, QueryLanguageParserQUOTED, QueryLanguageParserTERM:
		{
			p.SetState(28)
			p.Clause()
		}

	case QueryLanguageParserLPAREN:
		{
			p.SetState(29)
			p.Match(QueryLanguageParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.query(0)
		}
		{
			p.SetState(31)
			p.Match(QueryLanguageParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case QueryLanguageParserNOT:
		{
			p.SetState(33)
			p.Match(QueryLanguageParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.query(3)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewQueryContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, QueryLanguageParserRULE_query)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(38)
					p.Match(QueryLanguageParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(39)
					p.query(3)
				}

			case 2:
				localctx = NewQueryContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, QueryLanguageParserRULE_query)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(41)
					p.Match(QueryLanguageParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(42)
					p.query(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IsPresentClause() IIsPresentClauseContext
	InClause() IInClauseContext
	CompareClause() ICompareClauseContext
	CountClause() ICountClauseContext

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_clause
	return p
}

func InitEmptyClauseContext(p *ClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_clause
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) IsPresentClause() IIsPresentClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsPresentClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsPresentClauseContext)
}

func (s *ClauseContext) InClause() IInClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInClauseContext)
}

func (s *ClauseContext) CompareClause() ICompareClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareClauseContext)
}

func (s *ClauseContext) CountClause() ICountClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountClauseContext)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) Clause() (localctx IClauseContext) {
	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryLanguageParserRULE_clause)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.IsPresentClause()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.InClause()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.CompareClause()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.CountClause()
		}

	case antlr.ATNInvalidAltNumber:
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

// IIsPresentClauseContext is an interface to support dynamic dispatch.
type IIsPresentClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	IS() antlr.TerminalNode
	PRESENT() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsIsPresentClauseContext differentiates from other interfaces.
	IsIsPresentClauseContext()
}

type IsPresentClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsPresentClauseContext() *IsPresentClauseContext {
	var p = new(IsPresentClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_isPresentClause
	return p
}

func InitEmptyIsPresentClauseContext(p *IsPresentClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_isPresentClause
}

func (*IsPresentClauseContext) IsIsPresentClauseContext() {}

func NewIsPresentClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsPresentClauseContext {
	var p = new(IsPresentClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_isPresentClause

	return p
}

func (s *IsPresentClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *IsPresentClauseContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *IsPresentClauseContext) IS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserIS, 0)
}

func (s *IsPresentClauseContext) PRESENT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserPRESENT, 0)
}

func (s *IsPresentClauseContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserNOT, 0)
}

func (s *IsPresentClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsPresentClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsPresentClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitIsPresentClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) IsPresentClause() (localctx IIsPresentClauseContext) {
	localctx = NewIsPresentClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryLanguageParserRULE_isPresentClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.FieldName()
	}
	{
		p.SetState(55)
		p.Match(QueryLanguageParserIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserNOT {
		{
			p.SetState(56)
			p.Match(QueryLanguageParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(59)
		p.Match(QueryLanguageParserPRESENT)
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

// IInClauseContext is an interface to support dynamic dispatch.
type IInClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	IN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Valuelist() IValuelistContext
	RPAREN() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsInClauseContext differentiates from other interfaces.
	IsInClauseContext()
}

type InClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInClauseContext() *InClauseContext {
	var p = new(InClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_inClause
	return p
}

func InitEmptyInClauseContext(p *InClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_inClause
}

func (*InClauseContext) IsInClauseContext() {}

func NewInClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InClauseContext {
	var p = new(InClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_inClause

	return p
}

func (s *InClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *InClauseContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *InClauseContext) IN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserIN, 0)
}

func (s *InClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserLPAREN, 0)
}

func (s *InClauseContext) Valuelist() IValuelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuelistContext)
}

func (s *InClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserRPAREN, 0)
}

func (s *InClauseContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserNOT, 0)
}

func (s *InClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitInClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) InClause() (localctx IInClauseContext) {
	localctx = NewInClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryLanguageParserRULE_inClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.FieldName()
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserNOT {
		{
			p.SetState(62)
			p.Match(QueryLanguageParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(65)
		p.Match(QueryLanguageParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(QueryLanguageParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Valuelist()
	}
	{
		p.SetState(68)
		p.Match(QueryLanguageParserRPAREN)
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

// ICompareClauseContext is an interface to support dynamic dispatch.
type ICompareClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	FieldName() IFieldNameContext
	Value() IValueContext
	OP_TILDE() antlr.TerminalNode
	OP_NOT_TILDE() antlr.TerminalNode
	OP_EQUAL() antlr.TerminalNode
	OP_NOT_EQUAL() antlr.TerminalNode
	OP_TILDE_IGNORE_CASE() antlr.TerminalNode
	OP_NOT_TILDE_IGNORE_CASE() antlr.TerminalNode
	OP_EQUAL_IGNORE_CASE() antlr.TerminalNode
	OP_NOT_EQUAL_IGNORE_CASE() antlr.TerminalNode

	// IsCompareClauseContext differentiates from other interfaces.
	IsCompareClauseContext()
}

type CompareClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyCompareClauseContext() *CompareClauseContext {
	var p = new(CompareClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_compareClause
	return p
}

func InitEmptyCompareClauseContext(p *CompareClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_compareClause
}

func (*CompareClauseContext) IsCompareClauseContext() {}

func NewCompareClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareClauseContext {
	var p = new(CompareClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_compareClause

	return p
}

func (s *CompareClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareClauseContext) GetOp() antlr.Token { return s.op }

func (s *CompareClauseContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareClauseContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *CompareClauseContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CompareClauseContext) OP_TILDE() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_TILDE, 0)
}

func (s *CompareClauseContext) OP_NOT_TILDE() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_NOT_TILDE, 0)
}

func (s *CompareClauseContext) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_EQUAL, 0)
}

func (s *CompareClauseContext) OP_NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_NOT_EQUAL, 0)
}

func (s *CompareClauseContext) OP_TILDE_IGNORE_CASE() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_TILDE_IGNORE_CASE, 0)
}

func (s *CompareClauseContext) OP_NOT_TILDE_IGNORE_CASE() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_NOT_TILDE_IGNORE_CASE, 0)
}

func (s *CompareClauseContext) OP_EQUAL_IGNORE_CASE() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_EQUAL_IGNORE_CASE, 0)
}

func (s *CompareClauseContext) OP_NOT_EQUAL_IGNORE_CASE() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_NOT_EQUAL_IGNORE_CASE, 0)
}

func (s *CompareClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitCompareClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) CompareClause() (localctx ICompareClauseContext) {
	localctx = NewCompareClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QueryLanguageParserRULE_compareClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.FieldName()
	}
	{
		p.SetState(71)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CompareClauseContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&522240) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CompareClauseContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(72)
		p.Value()
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

// ICountClauseContext is an interface to support dynamic dispatch.
type ICountClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	COUNT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FieldName() IFieldNameContext
	RPAREN() antlr.TerminalNode
	Int_value() IInt_valueContext
	OP_EQUAL() antlr.TerminalNode
	OP_NOT_EQUAL() antlr.TerminalNode
	OP_GREATER_THAN() antlr.TerminalNode
	OP_GREATER_THAN_EQUAL() antlr.TerminalNode
	OP_LESS_THAN() antlr.TerminalNode
	OP_LESS_THAN_EQUAL() antlr.TerminalNode

	// IsCountClauseContext differentiates from other interfaces.
	IsCountClauseContext()
}

type CountClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyCountClauseContext() *CountClauseContext {
	var p = new(CountClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_countClause
	return p
}

func InitEmptyCountClauseContext(p *CountClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_countClause
}

func (*CountClauseContext) IsCountClauseContext() {}

func NewCountClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountClauseContext {
	var p = new(CountClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_countClause

	return p
}

func (s *CountClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CountClauseContext) GetOp() antlr.Token { return s.op }

func (s *CountClauseContext) SetOp(v antlr.Token) { s.op = v }

func (s *CountClauseContext) COUNT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserCOUNT, 0)
}

func (s *CountClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserLPAREN, 0)
}

func (s *CountClauseContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *CountClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserRPAREN, 0)
}

func (s *CountClauseContext) Int_value() IInt_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInt_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInt_valueContext)
}

func (s *CountClauseContext) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_EQUAL, 0)
}

func (s *CountClauseContext) OP_NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_NOT_EQUAL, 0)
}

func (s *CountClauseContext) OP_GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_GREATER_THAN, 0)
}

func (s *CountClauseContext) OP_GREATER_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_GREATER_THAN_EQUAL, 0)
}

func (s *CountClauseContext) OP_LESS_THAN() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_LESS_THAN, 0)
}

func (s *CountClauseContext) OP_LESS_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserOP_LESS_THAN_EQUAL, 0)
}

func (s *CountClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitCountClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) CountClause() (localctx ICountClauseContext) {
	localctx = NewCountClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QueryLanguageParserRULE_countClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(QueryLanguageParserCOUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(QueryLanguageParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.FieldName()
	}
	{
		p.SetState(77)
		p.Match(QueryLanguageParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CountClauseContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7870464) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CountClauseContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(79)
		p.Int_value()
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

// IValuelistContext is an interface to support dynamic dispatch.
type IValuelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsValuelistContext differentiates from other interfaces.
	IsValuelistContext()
}

type ValuelistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuelistContext() *ValuelistContext {
	var p = new(ValuelistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_valuelist
	return p
}

func InitEmptyValuelistContext(p *ValuelistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_valuelist
}

func (*ValuelistContext) IsValuelistContext() {}

func NewValuelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuelistContext {
	var p = new(ValuelistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_valuelist

	return p
}

func (s *ValuelistContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuelistContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ValuelistContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ValuelistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QueryLanguageParserCOMMA)
}

func (s *ValuelistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserCOMMA, i)
}

func (s *ValuelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuelistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitValuelist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) Valuelist() (localctx IValuelistContext) {
	localctx = NewValuelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QueryLanguageParserRULE_valuelist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Value()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserCOMMA {
		{
			p.SetState(82)
			p.Match(QueryLanguageParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Value()
		}

		p.SetState(88)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED() antlr.TerminalNode
	TERM() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	VARIABLE() antlr.TerminalNode
	PLACEHOLDER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserQUOTED, 0)
}

func (s *ValueContext) TERM() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTERM, 0)
}

func (s *ValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINTEGER, 0)
}

func (s *ValueContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserVARIABLE, 0)
}

func (s *ValueContext) PLACEHOLDER() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserPLACEHOLDER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QueryLanguageParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&260046848) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED() antlr.TerminalNode
	TERM() antlr.TerminalNode

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserQUOTED, 0)
}

func (s *FieldNameContext) TERM() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTERM, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QueryLanguageParserRULE_fieldName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryLanguageParserQUOTED || _la == QueryLanguageParserTERM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IInt_valueContext is an interface to support dynamic dispatch.
type IInt_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode

	// IsInt_valueContext differentiates from other interfaces.
	IsInt_valueContext()
}

type Int_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt_valueContext() *Int_valueContext {
	var p = new(Int_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_int_value
	return p
}

func InitEmptyInt_valueContext(p *Int_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_int_value
}

func (*Int_valueContext) IsInt_valueContext() {}

func NewInt_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int_valueContext {
	var p = new(Int_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_int_value

	return p
}

func (s *Int_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Int_valueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINTEGER, 0)
}

func (s *Int_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageParserVisitor:
		return t.VisitInt_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryLanguageParser) Int_value() (localctx IInt_valueContext) {
	localctx = NewInt_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QueryLanguageParserRULE_int_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(QueryLanguageParserINTEGER)
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

func (p *QueryLanguageParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QueryLanguageParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
