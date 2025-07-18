// Code generated from SQLQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // SQLQueryParser
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

type SQLQueryParser struct {
	*antlr.BaseParser
}

var SQLQueryParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqlqueryparserParserInit() {
	staticData := &SQLQueryParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'--'", "':'", "", "", "", "", "'when'", "'<when'", "'+when'", "",
		"", "", "", "", "", "'.'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "STARTCOMMENT", "STARTEXPR", "EOL", "WS", "SQLTEXT", "STARTQUERY",
		"STARTWHEN", "STARTCHOMPWHEN", "STARTJOINWHEN", "ENDWHEN", "ENDCOMMENT",
		"RAWCOMMENT", "QUERY_NAME", "ENDQUERY", "IDENT", "IDENT_SEP", "IDENT_LITERAL",
		"SPACE_EXPR", "WS_EXPR",
	}
	staticData.RuleNames = []string{
		"querySet", "query", "sqlChunk", "sqlLine", "sqlFragment", "sqlText",
		"whenDirective", "chompWhenDirective", "joinWhen", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 116, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 4,
		0, 22, 8, 0, 11, 0, 12, 0, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 5, 2, 34, 8, 2, 10, 2, 12, 2, 37, 9, 2, 1, 2, 1, 2, 4, 2, 41, 8,
		2, 11, 2, 12, 2, 42, 1, 3, 5, 3, 46, 8, 3, 10, 3, 12, 3, 49, 9, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 3, 4, 55, 8, 4, 1, 5, 4, 5, 58, 8, 5, 11, 5, 12, 5, 59,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 67, 8, 6, 10, 6, 12, 6, 70, 9, 6, 1,
		6, 1, 6, 1, 7, 4, 7, 75, 8, 7, 11, 7, 12, 7, 76, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 5, 7, 84, 8, 7, 10, 7, 12, 7, 87, 9, 7, 1, 7, 1, 7, 1, 8, 4, 8, 92,
		8, 8, 11, 8, 12, 8, 93, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 5, 9, 105, 8, 9, 10, 9, 12, 9, 108, 9, 9, 1, 9, 1, 9, 3, 9, 112,
		8, 9, 3, 9, 114, 8, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		0, 0, 120, 0, 21, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 40, 1, 0, 0, 0, 6,
		47, 1, 0, 0, 0, 8, 54, 1, 0, 0, 0, 10, 57, 1, 0, 0, 0, 12, 61, 1, 0, 0,
		0, 14, 74, 1, 0, 0, 0, 16, 91, 1, 0, 0, 0, 18, 100, 1, 0, 0, 0, 20, 22,
		3, 2, 1, 0, 21, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0,
		23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 5, 0, 0, 1, 26, 1, 1, 0,
		0, 0, 27, 28, 5, 6, 0, 0, 28, 29, 5, 13, 0, 0, 29, 30, 3, 4, 2, 0, 30,
		3, 1, 0, 0, 0, 31, 41, 3, 6, 3, 0, 32, 34, 3, 8, 4, 0, 33, 32, 1, 0, 0,
		0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38,
		1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 41, 3, 12, 6, 0, 39, 41, 3, 14, 7,
		0, 40, 31, 1, 0, 0, 0, 40, 35, 1, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 42,
		1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 5, 1, 0, 0, 0,
		44, 46, 3, 8, 4, 0, 45, 44, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1,
		0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50,
		51, 5, 3, 0, 0, 51, 7, 1, 0, 0, 0, 52, 55, 3, 10, 5, 0, 53, 55, 3, 18,
		9, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 9, 1, 0, 0, 0, 56, 58,
		5, 5, 0, 0, 57, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 11, 1, 0, 0, 0, 61, 62, 5, 7, 0, 0, 62, 63, 3,
		18, 9, 0, 63, 64, 5, 3, 0, 0, 64, 68, 3, 4, 2, 0, 65, 67, 3, 16, 8, 0,
		66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1,
		0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 72, 5, 10, 0, 0, 72,
		13, 1, 0, 0, 0, 73, 75, 3, 8, 4, 0, 74, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0,
		0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79,
		5, 8, 0, 0, 79, 80, 3, 18, 9, 0, 80, 81, 5, 3, 0, 0, 81, 85, 3, 4, 2, 0,
		82, 84, 3, 16, 8, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1,
		0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88,
		89, 5, 10, 0, 0, 89, 15, 1, 0, 0, 0, 90, 92, 3, 8, 4, 0, 91, 90, 1, 0,
		0, 0, 92, 93, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 96, 5, 9, 0, 0, 96, 97, 3, 18, 9, 0, 97, 98, 5, 3, 0, 0,
		98, 99, 3, 4, 2, 0, 99, 17, 1, 0, 0, 0, 100, 101, 5, 2, 0, 0, 101, 106,
		5, 15, 0, 0, 102, 103, 5, 16, 0, 0, 103, 105, 5, 15, 0, 0, 104, 102, 1,
		0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0,
		0, 107, 113, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 111, 5, 17, 0, 0, 110,
		112, 5, 15, 0, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114,
		1, 0, 0, 0, 113, 109, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 19, 1, 0,
		0, 0, 14, 23, 35, 40, 42, 47, 54, 59, 68, 76, 85, 93, 106, 111, 113,
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

// SQLQueryParserInit initializes any static state used to implement SQLQueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSQLQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SQLQueryParserInit() {
	staticData := &SQLQueryParserParserStaticData
	staticData.once.Do(sqlqueryparserParserInit)
}

// NewSQLQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSQLQueryParser(input antlr.TokenStream) *SQLQueryParser {
	SQLQueryParserInit()
	this := new(SQLQueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SQLQueryParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SQLQueryParser.g4"

	return this
}

// SQLQueryParser tokens.
const (
	SQLQueryParserEOF            = antlr.TokenEOF
	SQLQueryParserSTARTCOMMENT   = 1
	SQLQueryParserSTARTEXPR      = 2
	SQLQueryParserEOL            = 3
	SQLQueryParserWS             = 4
	SQLQueryParserSQLTEXT        = 5
	SQLQueryParserSTARTQUERY     = 6
	SQLQueryParserSTARTWHEN      = 7
	SQLQueryParserSTARTCHOMPWHEN = 8
	SQLQueryParserSTARTJOINWHEN  = 9
	SQLQueryParserENDWHEN        = 10
	SQLQueryParserENDCOMMENT     = 11
	SQLQueryParserRAWCOMMENT     = 12
	SQLQueryParserQUERY_NAME     = 13
	SQLQueryParserENDQUERY       = 14
	SQLQueryParserIDENT          = 15
	SQLQueryParserIDENT_SEP      = 16
	SQLQueryParserIDENT_LITERAL  = 17
	SQLQueryParserSPACE_EXPR     = 18
	SQLQueryParserWS_EXPR        = 19
)

// SQLQueryParser rules.
const (
	SQLQueryParserRULE_querySet           = 0
	SQLQueryParserRULE_query              = 1
	SQLQueryParserRULE_sqlChunk           = 2
	SQLQueryParserRULE_sqlLine            = 3
	SQLQueryParserRULE_sqlFragment        = 4
	SQLQueryParserRULE_sqlText            = 5
	SQLQueryParserRULE_whenDirective      = 6
	SQLQueryParserRULE_chompWhenDirective = 7
	SQLQueryParserRULE_joinWhen           = 8
	SQLQueryParserRULE_expr               = 9
)

// IQuerySetContext is an interface to support dynamic dispatch.
type IQuerySetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllQuery() []IQueryContext
	Query(i int) IQueryContext

	// IsQuerySetContext differentiates from other interfaces.
	IsQuerySetContext()
}

type QuerySetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuerySetContext() *QuerySetContext {
	var p = new(QuerySetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_querySet
	return p
}

func InitEmptyQuerySetContext(p *QuerySetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_querySet
}

func (*QuerySetContext) IsQuerySetContext() {}

func NewQuerySetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuerySetContext {
	var p = new(QuerySetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_querySet

	return p
}

func (s *QuerySetContext) GetParser() antlr.Parser { return s.parser }

func (s *QuerySetContext) EOF() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserEOF, 0)
}

func (s *QuerySetContext) AllQuery() []IQueryContext {
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

func (s *QuerySetContext) Query(i int) IQueryContext {
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

func (s *QuerySetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuerySetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuerySetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterQuerySet(s)
	}
}

func (s *QuerySetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitQuerySet(s)
	}
}

func (s *QuerySetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitQuerySet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) QuerySet() (localctx IQuerySetContext) {
	localctx = NewQuerySetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SQLQueryParserRULE_querySet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SQLQueryParserSTARTQUERY {
		{
			p.SetState(20)
			p.Query()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(SQLQueryParserEOF)
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

	// GetQueryName returns the queryName token.
	GetQueryName() antlr.Token

	// SetQueryName sets the queryName token.
	SetQueryName(antlr.Token)

	// GetQueryBody returns the queryBody rule contexts.
	GetQueryBody() ISqlChunkContext

	// SetQueryBody sets the queryBody rule contexts.
	SetQueryBody(ISqlChunkContext)

	// Getter signatures
	STARTQUERY() antlr.TerminalNode
	QUERY_NAME() antlr.TerminalNode
	SqlChunk() ISqlChunkContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	queryName antlr.Token
	queryBody ISqlChunkContext
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) GetQueryName() antlr.Token { return s.queryName }

func (s *QueryContext) SetQueryName(v antlr.Token) { s.queryName = v }

func (s *QueryContext) GetQueryBody() ISqlChunkContext { return s.queryBody }

func (s *QueryContext) SetQueryBody(v ISqlChunkContext) { s.queryBody = v }

func (s *QueryContext) STARTQUERY() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserSTARTQUERY, 0)
}

func (s *QueryContext) QUERY_NAME() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserQUERY_NAME, 0)
}

func (s *QueryContext) SqlChunk() ISqlChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqlChunkContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SQLQueryParserRULE_query)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(SQLQueryParserSTARTQUERY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)

		var _m = p.Match(SQLQueryParserQUERY_NAME)

		localctx.(*QueryContext).queryName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)

		var _x = p.SqlChunk()

		localctx.(*QueryContext).queryBody = _x
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

// ISqlChunkContext is an interface to support dynamic dispatch.
type ISqlChunkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSqlLine() []ISqlLineContext
	SqlLine(i int) ISqlLineContext
	AllWhenDirective() []IWhenDirectiveContext
	WhenDirective(i int) IWhenDirectiveContext
	AllChompWhenDirective() []IChompWhenDirectiveContext
	ChompWhenDirective(i int) IChompWhenDirectiveContext
	AllSqlFragment() []ISqlFragmentContext
	SqlFragment(i int) ISqlFragmentContext

	// IsSqlChunkContext differentiates from other interfaces.
	IsSqlChunkContext()
}

type SqlChunkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlChunkContext() *SqlChunkContext {
	var p = new(SqlChunkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_sqlChunk
	return p
}

func InitEmptySqlChunkContext(p *SqlChunkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_sqlChunk
}

func (*SqlChunkContext) IsSqlChunkContext() {}

func NewSqlChunkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlChunkContext {
	var p = new(SqlChunkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_sqlChunk

	return p
}

func (s *SqlChunkContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlChunkContext) AllSqlLine() []ISqlLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISqlLineContext); ok {
			len++
		}
	}

	tst := make([]ISqlLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISqlLineContext); ok {
			tst[i] = t.(ISqlLineContext)
			i++
		}
	}

	return tst
}

func (s *SqlChunkContext) SqlLine(i int) ISqlLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlLineContext); ok {
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

	return t.(ISqlLineContext)
}

func (s *SqlChunkContext) AllWhenDirective() []IWhenDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhenDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IWhenDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhenDirectiveContext); ok {
			tst[i] = t.(IWhenDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *SqlChunkContext) WhenDirective(i int) IWhenDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhenDirectiveContext); ok {
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

	return t.(IWhenDirectiveContext)
}

func (s *SqlChunkContext) AllChompWhenDirective() []IChompWhenDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChompWhenDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IChompWhenDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChompWhenDirectiveContext); ok {
			tst[i] = t.(IChompWhenDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *SqlChunkContext) ChompWhenDirective(i int) IChompWhenDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChompWhenDirectiveContext); ok {
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

	return t.(IChompWhenDirectiveContext)
}

func (s *SqlChunkContext) AllSqlFragment() []ISqlFragmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISqlFragmentContext); ok {
			len++
		}
	}

	tst := make([]ISqlFragmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISqlFragmentContext); ok {
			tst[i] = t.(ISqlFragmentContext)
			i++
		}
	}

	return tst
}

func (s *SqlChunkContext) SqlFragment(i int) ISqlFragmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlFragmentContext); ok {
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

	return t.(ISqlFragmentContext)
}

func (s *SqlChunkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlChunkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlChunkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterSqlChunk(s)
	}
}

func (s *SqlChunkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitSqlChunk(s)
	}
}

func (s *SqlChunkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitSqlChunk(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) SqlChunk() (localctx ISqlChunkContext) {
	localctx = NewSqlChunkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SQLQueryParserRULE_sqlChunk)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(31)
					p.SqlLine()
				}

			case 2:
				p.SetState(35)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == SQLQueryParserSTARTEXPR || _la == SQLQueryParserSQLTEXT {
					{
						p.SetState(32)
						p.SqlFragment()
					}

					p.SetState(37)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(38)
					p.WhenDirective()
				}

			case 3:
				{
					p.SetState(39)
					p.ChompWhenDirective()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISqlLineContext is an interface to support dynamic dispatch.
type ISqlLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOL() antlr.TerminalNode
	AllSqlFragment() []ISqlFragmentContext
	SqlFragment(i int) ISqlFragmentContext

	// IsSqlLineContext differentiates from other interfaces.
	IsSqlLineContext()
}

type SqlLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlLineContext() *SqlLineContext {
	var p = new(SqlLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_sqlLine
	return p
}

func InitEmptySqlLineContext(p *SqlLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_sqlLine
}

func (*SqlLineContext) IsSqlLineContext() {}

func NewSqlLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlLineContext {
	var p = new(SqlLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_sqlLine

	return p
}

func (s *SqlLineContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlLineContext) EOL() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserEOL, 0)
}

func (s *SqlLineContext) AllSqlFragment() []ISqlFragmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISqlFragmentContext); ok {
			len++
		}
	}

	tst := make([]ISqlFragmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISqlFragmentContext); ok {
			tst[i] = t.(ISqlFragmentContext)
			i++
		}
	}

	return tst
}

func (s *SqlLineContext) SqlFragment(i int) ISqlFragmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlFragmentContext); ok {
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

	return t.(ISqlFragmentContext)
}

func (s *SqlLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterSqlLine(s)
	}
}

func (s *SqlLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitSqlLine(s)
	}
}

func (s *SqlLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitSqlLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) SqlLine() (localctx ISqlLineContext) {
	localctx = NewSqlLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLQueryParserRULE_sqlLine)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLQueryParserSTARTEXPR || _la == SQLQueryParserSQLTEXT {
		{
			p.SetState(44)
			p.SqlFragment()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(SQLQueryParserEOL)
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

// ISqlFragmentContext is an interface to support dynamic dispatch.
type ISqlFragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SqlText() ISqlTextContext
	Expr() IExprContext

	// IsSqlFragmentContext differentiates from other interfaces.
	IsSqlFragmentContext()
}

type SqlFragmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlFragmentContext() *SqlFragmentContext {
	var p = new(SqlFragmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_sqlFragment
	return p
}

func InitEmptySqlFragmentContext(p *SqlFragmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_sqlFragment
}

func (*SqlFragmentContext) IsSqlFragmentContext() {}

func NewSqlFragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlFragmentContext {
	var p = new(SqlFragmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_sqlFragment

	return p
}

func (s *SqlFragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlFragmentContext) SqlText() ISqlTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqlTextContext)
}

func (s *SqlFragmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SqlFragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlFragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlFragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterSqlFragment(s)
	}
}

func (s *SqlFragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitSqlFragment(s)
	}
}

func (s *SqlFragmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitSqlFragment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) SqlFragment() (localctx ISqlFragmentContext) {
	localctx = NewSqlFragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLQueryParserRULE_sqlFragment)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLQueryParserSQLTEXT:
		{
			p.SetState(52)
			p.SqlText()
		}

	case SQLQueryParserSTARTEXPR:
		{
			p.SetState(53)
			p.Expr()
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

// ISqlTextContext is an interface to support dynamic dispatch.
type ISqlTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SQLTEXT returns the _SQLTEXT token.
	Get_SQLTEXT() antlr.Token

	// Set_SQLTEXT sets the _SQLTEXT token.
	Set_SQLTEXT(antlr.Token)

	// GetTextFragments returns the textFragments token list.
	GetTextFragments() []antlr.Token

	// SetTextFragments sets the textFragments token list.
	SetTextFragments([]antlr.Token)

	// Getter signatures
	AllSQLTEXT() []antlr.TerminalNode
	SQLTEXT(i int) antlr.TerminalNode

	// IsSqlTextContext differentiates from other interfaces.
	IsSqlTextContext()
}

type SqlTextContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	_SQLTEXT      antlr.Token
	textFragments []antlr.Token
}

func NewEmptySqlTextContext() *SqlTextContext {
	var p = new(SqlTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_sqlText
	return p
}

func InitEmptySqlTextContext(p *SqlTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_sqlText
}

func (*SqlTextContext) IsSqlTextContext() {}

func NewSqlTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlTextContext {
	var p = new(SqlTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_sqlText

	return p
}

func (s *SqlTextContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlTextContext) Get_SQLTEXT() antlr.Token { return s._SQLTEXT }

func (s *SqlTextContext) Set_SQLTEXT(v antlr.Token) { s._SQLTEXT = v }

func (s *SqlTextContext) GetTextFragments() []antlr.Token { return s.textFragments }

func (s *SqlTextContext) SetTextFragments(v []antlr.Token) { s.textFragments = v }

func (s *SqlTextContext) AllSQLTEXT() []antlr.TerminalNode {
	return s.GetTokens(SQLQueryParserSQLTEXT)
}

func (s *SqlTextContext) SQLTEXT(i int) antlr.TerminalNode {
	return s.GetToken(SQLQueryParserSQLTEXT, i)
}

func (s *SqlTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterSqlText(s)
	}
}

func (s *SqlTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitSqlText(s)
	}
}

func (s *SqlTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitSqlText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) SqlText() (localctx ISqlTextContext) {
	localctx = NewSqlTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLQueryParserRULE_sqlText)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(56)

				var _m = p.Match(SQLQueryParserSQLTEXT)

				localctx.(*SqlTextContext)._SQLTEXT = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			localctx.(*SqlTextContext).textFragments = append(localctx.(*SqlTextContext).textFragments, localctx.(*SqlTextContext)._SQLTEXT)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhenDirectiveContext is an interface to support dynamic dispatch.
type IWhenDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWhenCondition returns the whenCondition rule contexts.
	GetWhenCondition() IExprContext

	// GetBody returns the body rule contexts.
	GetBody() ISqlChunkContext

	// Get_joinWhen returns the _joinWhen rule contexts.
	Get_joinWhen() IJoinWhenContext

	// SetWhenCondition sets the whenCondition rule contexts.
	SetWhenCondition(IExprContext)

	// SetBody sets the body rule contexts.
	SetBody(ISqlChunkContext)

	// Set_joinWhen sets the _joinWhen rule contexts.
	Set_joinWhen(IJoinWhenContext)

	// GetJoinWhens returns the joinWhens rule context list.
	GetJoinWhens() []IJoinWhenContext

	// SetJoinWhens sets the joinWhens rule context list.
	SetJoinWhens([]IJoinWhenContext)

	// Getter signatures
	STARTWHEN() antlr.TerminalNode
	EOL() antlr.TerminalNode
	ENDWHEN() antlr.TerminalNode
	Expr() IExprContext
	SqlChunk() ISqlChunkContext
	AllJoinWhen() []IJoinWhenContext
	JoinWhen(i int) IJoinWhenContext

	// IsWhenDirectiveContext differentiates from other interfaces.
	IsWhenDirectiveContext()
}

type WhenDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	whenCondition IExprContext
	body          ISqlChunkContext
	_joinWhen     IJoinWhenContext
	joinWhens     []IJoinWhenContext
}

func NewEmptyWhenDirectiveContext() *WhenDirectiveContext {
	var p = new(WhenDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_whenDirective
	return p
}

func InitEmptyWhenDirectiveContext(p *WhenDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_whenDirective
}

func (*WhenDirectiveContext) IsWhenDirectiveContext() {}

func NewWhenDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenDirectiveContext {
	var p = new(WhenDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_whenDirective

	return p
}

func (s *WhenDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenDirectiveContext) GetWhenCondition() IExprContext { return s.whenCondition }

func (s *WhenDirectiveContext) GetBody() ISqlChunkContext { return s.body }

func (s *WhenDirectiveContext) Get_joinWhen() IJoinWhenContext { return s._joinWhen }

func (s *WhenDirectiveContext) SetWhenCondition(v IExprContext) { s.whenCondition = v }

func (s *WhenDirectiveContext) SetBody(v ISqlChunkContext) { s.body = v }

func (s *WhenDirectiveContext) Set_joinWhen(v IJoinWhenContext) { s._joinWhen = v }

func (s *WhenDirectiveContext) GetJoinWhens() []IJoinWhenContext { return s.joinWhens }

func (s *WhenDirectiveContext) SetJoinWhens(v []IJoinWhenContext) { s.joinWhens = v }

func (s *WhenDirectiveContext) STARTWHEN() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserSTARTWHEN, 0)
}

func (s *WhenDirectiveContext) EOL() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserEOL, 0)
}

func (s *WhenDirectiveContext) ENDWHEN() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserENDWHEN, 0)
}

func (s *WhenDirectiveContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhenDirectiveContext) SqlChunk() ISqlChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqlChunkContext)
}

func (s *WhenDirectiveContext) AllJoinWhen() []IJoinWhenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinWhenContext); ok {
			len++
		}
	}

	tst := make([]IJoinWhenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinWhenContext); ok {
			tst[i] = t.(IJoinWhenContext)
			i++
		}
	}

	return tst
}

func (s *WhenDirectiveContext) JoinWhen(i int) IJoinWhenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinWhenContext); ok {
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

	return t.(IJoinWhenContext)
}

func (s *WhenDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterWhenDirective(s)
	}
}

func (s *WhenDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitWhenDirective(s)
	}
}

func (s *WhenDirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitWhenDirective(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) WhenDirective() (localctx IWhenDirectiveContext) {
	localctx = NewWhenDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLQueryParserRULE_whenDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(SQLQueryParserSTARTWHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)

		var _x = p.Expr()

		localctx.(*WhenDirectiveContext).whenCondition = _x
	}
	{
		p.SetState(63)
		p.Match(SQLQueryParserEOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)

		var _x = p.SqlChunk()

		localctx.(*WhenDirectiveContext).body = _x
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLQueryParserSTARTEXPR || _la == SQLQueryParserSQLTEXT {
		{
			p.SetState(65)

			var _x = p.JoinWhen()

			localctx.(*WhenDirectiveContext)._joinWhen = _x
		}
		localctx.(*WhenDirectiveContext).joinWhens = append(localctx.(*WhenDirectiveContext).joinWhens, localctx.(*WhenDirectiveContext)._joinWhen)

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(71)
		p.Match(SQLQueryParserENDWHEN)
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

// IChompWhenDirectiveContext is an interface to support dynamic dispatch.
type IChompWhenDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix rule contexts.
	GetPrefix() ISqlFragmentContext

	// GetWhenCondition returns the whenCondition rule contexts.
	GetWhenCondition() IExprContext

	// GetBody returns the body rule contexts.
	GetBody() ISqlChunkContext

	// Get_joinWhen returns the _joinWhen rule contexts.
	Get_joinWhen() IJoinWhenContext

	// SetPrefix sets the prefix rule contexts.
	SetPrefix(ISqlFragmentContext)

	// SetWhenCondition sets the whenCondition rule contexts.
	SetWhenCondition(IExprContext)

	// SetBody sets the body rule contexts.
	SetBody(ISqlChunkContext)

	// Set_joinWhen sets the _joinWhen rule contexts.
	Set_joinWhen(IJoinWhenContext)

	// GetJoinWhens returns the joinWhens rule context list.
	GetJoinWhens() []IJoinWhenContext

	// SetJoinWhens sets the joinWhens rule context list.
	SetJoinWhens([]IJoinWhenContext)

	// Getter signatures
	STARTCHOMPWHEN() antlr.TerminalNode
	EOL() antlr.TerminalNode
	ENDWHEN() antlr.TerminalNode
	Expr() IExprContext
	SqlChunk() ISqlChunkContext
	AllSqlFragment() []ISqlFragmentContext
	SqlFragment(i int) ISqlFragmentContext
	AllJoinWhen() []IJoinWhenContext
	JoinWhen(i int) IJoinWhenContext

	// IsChompWhenDirectiveContext differentiates from other interfaces.
	IsChompWhenDirectiveContext()
}

type ChompWhenDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	prefix        ISqlFragmentContext
	whenCondition IExprContext
	body          ISqlChunkContext
	_joinWhen     IJoinWhenContext
	joinWhens     []IJoinWhenContext
}

func NewEmptyChompWhenDirectiveContext() *ChompWhenDirectiveContext {
	var p = new(ChompWhenDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_chompWhenDirective
	return p
}

func InitEmptyChompWhenDirectiveContext(p *ChompWhenDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_chompWhenDirective
}

func (*ChompWhenDirectiveContext) IsChompWhenDirectiveContext() {}

func NewChompWhenDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChompWhenDirectiveContext {
	var p = new(ChompWhenDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_chompWhenDirective

	return p
}

func (s *ChompWhenDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *ChompWhenDirectiveContext) GetPrefix() ISqlFragmentContext { return s.prefix }

func (s *ChompWhenDirectiveContext) GetWhenCondition() IExprContext { return s.whenCondition }

func (s *ChompWhenDirectiveContext) GetBody() ISqlChunkContext { return s.body }

func (s *ChompWhenDirectiveContext) Get_joinWhen() IJoinWhenContext { return s._joinWhen }

func (s *ChompWhenDirectiveContext) SetPrefix(v ISqlFragmentContext) { s.prefix = v }

func (s *ChompWhenDirectiveContext) SetWhenCondition(v IExprContext) { s.whenCondition = v }

func (s *ChompWhenDirectiveContext) SetBody(v ISqlChunkContext) { s.body = v }

func (s *ChompWhenDirectiveContext) Set_joinWhen(v IJoinWhenContext) { s._joinWhen = v }

func (s *ChompWhenDirectiveContext) GetJoinWhens() []IJoinWhenContext { return s.joinWhens }

func (s *ChompWhenDirectiveContext) SetJoinWhens(v []IJoinWhenContext) { s.joinWhens = v }

func (s *ChompWhenDirectiveContext) STARTCHOMPWHEN() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserSTARTCHOMPWHEN, 0)
}

func (s *ChompWhenDirectiveContext) EOL() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserEOL, 0)
}

func (s *ChompWhenDirectiveContext) ENDWHEN() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserENDWHEN, 0)
}

func (s *ChompWhenDirectiveContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ChompWhenDirectiveContext) SqlChunk() ISqlChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqlChunkContext)
}

func (s *ChompWhenDirectiveContext) AllSqlFragment() []ISqlFragmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISqlFragmentContext); ok {
			len++
		}
	}

	tst := make([]ISqlFragmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISqlFragmentContext); ok {
			tst[i] = t.(ISqlFragmentContext)
			i++
		}
	}

	return tst
}

func (s *ChompWhenDirectiveContext) SqlFragment(i int) ISqlFragmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlFragmentContext); ok {
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

	return t.(ISqlFragmentContext)
}

func (s *ChompWhenDirectiveContext) AllJoinWhen() []IJoinWhenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinWhenContext); ok {
			len++
		}
	}

	tst := make([]IJoinWhenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinWhenContext); ok {
			tst[i] = t.(IJoinWhenContext)
			i++
		}
	}

	return tst
}

func (s *ChompWhenDirectiveContext) JoinWhen(i int) IJoinWhenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinWhenContext); ok {
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

	return t.(IJoinWhenContext)
}

func (s *ChompWhenDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChompWhenDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChompWhenDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterChompWhenDirective(s)
	}
}

func (s *ChompWhenDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitChompWhenDirective(s)
	}
}

func (s *ChompWhenDirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitChompWhenDirective(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) ChompWhenDirective() (localctx IChompWhenDirectiveContext) {
	localctx = NewChompWhenDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SQLQueryParserRULE_chompWhenDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SQLQueryParserSTARTEXPR || _la == SQLQueryParserSQLTEXT {
		{
			p.SetState(73)

			var _x = p.SqlFragment()

			localctx.(*ChompWhenDirectiveContext).prefix = _x
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(SQLQueryParserSTARTCHOMPWHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)

		var _x = p.Expr()

		localctx.(*ChompWhenDirectiveContext).whenCondition = _x
	}
	{
		p.SetState(80)
		p.Match(SQLQueryParserEOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)

		var _x = p.SqlChunk()

		localctx.(*ChompWhenDirectiveContext).body = _x
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLQueryParserSTARTEXPR || _la == SQLQueryParserSQLTEXT {
		{
			p.SetState(82)

			var _x = p.JoinWhen()

			localctx.(*ChompWhenDirectiveContext)._joinWhen = _x
		}
		localctx.(*ChompWhenDirectiveContext).joinWhens = append(localctx.(*ChompWhenDirectiveContext).joinWhens, localctx.(*ChompWhenDirectiveContext)._joinWhen)

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(SQLQueryParserENDWHEN)
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

// IJoinWhenContext is an interface to support dynamic dispatch.
type IJoinWhenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix rule contexts.
	GetPrefix() ISqlFragmentContext

	// GetWhenCondition returns the whenCondition rule contexts.
	GetWhenCondition() IExprContext

	// GetBody returns the body rule contexts.
	GetBody() ISqlChunkContext

	// SetPrefix sets the prefix rule contexts.
	SetPrefix(ISqlFragmentContext)

	// SetWhenCondition sets the whenCondition rule contexts.
	SetWhenCondition(IExprContext)

	// SetBody sets the body rule contexts.
	SetBody(ISqlChunkContext)

	// Getter signatures
	STARTJOINWHEN() antlr.TerminalNode
	EOL() antlr.TerminalNode
	Expr() IExprContext
	SqlChunk() ISqlChunkContext
	AllSqlFragment() []ISqlFragmentContext
	SqlFragment(i int) ISqlFragmentContext

	// IsJoinWhenContext differentiates from other interfaces.
	IsJoinWhenContext()
}

type JoinWhenContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	prefix        ISqlFragmentContext
	whenCondition IExprContext
	body          ISqlChunkContext
}

func NewEmptyJoinWhenContext() *JoinWhenContext {
	var p = new(JoinWhenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_joinWhen
	return p
}

func InitEmptyJoinWhenContext(p *JoinWhenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_joinWhen
}

func (*JoinWhenContext) IsJoinWhenContext() {}

func NewJoinWhenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinWhenContext {
	var p = new(JoinWhenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_joinWhen

	return p
}

func (s *JoinWhenContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinWhenContext) GetPrefix() ISqlFragmentContext { return s.prefix }

func (s *JoinWhenContext) GetWhenCondition() IExprContext { return s.whenCondition }

func (s *JoinWhenContext) GetBody() ISqlChunkContext { return s.body }

func (s *JoinWhenContext) SetPrefix(v ISqlFragmentContext) { s.prefix = v }

func (s *JoinWhenContext) SetWhenCondition(v IExprContext) { s.whenCondition = v }

func (s *JoinWhenContext) SetBody(v ISqlChunkContext) { s.body = v }

func (s *JoinWhenContext) STARTJOINWHEN() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserSTARTJOINWHEN, 0)
}

func (s *JoinWhenContext) EOL() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserEOL, 0)
}

func (s *JoinWhenContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *JoinWhenContext) SqlChunk() ISqlChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqlChunkContext)
}

func (s *JoinWhenContext) AllSqlFragment() []ISqlFragmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISqlFragmentContext); ok {
			len++
		}
	}

	tst := make([]ISqlFragmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISqlFragmentContext); ok {
			tst[i] = t.(ISqlFragmentContext)
			i++
		}
	}

	return tst
}

func (s *JoinWhenContext) SqlFragment(i int) ISqlFragmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlFragmentContext); ok {
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

	return t.(ISqlFragmentContext)
}

func (s *JoinWhenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinWhenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinWhenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterJoinWhen(s)
	}
}

func (s *JoinWhenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitJoinWhen(s)
	}
}

func (s *JoinWhenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitJoinWhen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) JoinWhen() (localctx IJoinWhenContext) {
	localctx = NewJoinWhenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SQLQueryParserRULE_joinWhen)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SQLQueryParserSTARTEXPR || _la == SQLQueryParserSQLTEXT {
		{
			p.SetState(90)

			var _x = p.SqlFragment()

			localctx.(*JoinWhenContext).prefix = _x
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(SQLQueryParserSTARTJOINWHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)

		var _x = p.Expr()

		localctx.(*JoinWhenContext).whenCondition = _x
	}
	{
		p.SetState(97)
		p.Match(SQLQueryParserEOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)

		var _x = p.SqlChunk()

		localctx.(*JoinWhenContext).body = _x
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

	// Get_IDENT returns the _IDENT token.
	Get_IDENT() antlr.Token

	// GetLiteralFlag returns the literalFlag token.
	GetLiteralFlag() antlr.Token

	// GetLiteralMode returns the literalMode token.
	GetLiteralMode() antlr.Token

	// Set_IDENT sets the _IDENT token.
	Set_IDENT(antlr.Token)

	// SetLiteralFlag sets the literalFlag token.
	SetLiteralFlag(antlr.Token)

	// SetLiteralMode sets the literalMode token.
	SetLiteralMode(antlr.Token)

	// GetPath returns the path token list.
	GetPath() []antlr.Token

	// SetPath sets the path token list.
	SetPath([]antlr.Token)

	// Getter signatures
	STARTEXPR() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllIDENT_SEP() []antlr.TerminalNode
	IDENT_SEP(i int) antlr.TerminalNode
	IDENT_LITERAL() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	_IDENT      antlr.Token
	path        []antlr.Token
	literalFlag antlr.Token
	literalMode antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLQueryParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLQueryParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Get_IDENT() antlr.Token { return s._IDENT }

func (s *ExprContext) GetLiteralFlag() antlr.Token { return s.literalFlag }

func (s *ExprContext) GetLiteralMode() antlr.Token { return s.literalMode }

func (s *ExprContext) Set_IDENT(v antlr.Token) { s._IDENT = v }

func (s *ExprContext) SetLiteralFlag(v antlr.Token) { s.literalFlag = v }

func (s *ExprContext) SetLiteralMode(v antlr.Token) { s.literalMode = v }

func (s *ExprContext) GetPath() []antlr.Token { return s.path }

func (s *ExprContext) SetPath(v []antlr.Token) { s.path = v }

func (s *ExprContext) STARTEXPR() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserSTARTEXPR, 0)
}

func (s *ExprContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SQLQueryParserIDENT)
}

func (s *ExprContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SQLQueryParserIDENT, i)
}

func (s *ExprContext) AllIDENT_SEP() []antlr.TerminalNode {
	return s.GetTokens(SQLQueryParserIDENT_SEP)
}

func (s *ExprContext) IDENT_SEP(i int) antlr.TerminalNode {
	return s.GetToken(SQLQueryParserIDENT_SEP, i)
}

func (s *ExprContext) IDENT_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLQueryParserIDENT_LITERAL, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLQueryParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLQueryParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLQueryParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLQueryParserRULE_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(SQLQueryParserSTARTEXPR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)

		var _m = p.Match(SQLQueryParserIDENT)

		localctx.(*ExprContext)._IDENT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ExprContext).path = append(localctx.(*ExprContext).path, localctx.(*ExprContext)._IDENT)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLQueryParserIDENT_SEP {
		{
			p.SetState(102)
			p.Match(SQLQueryParserIDENT_SEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)

			var _m = p.Match(SQLQueryParserIDENT)

			localctx.(*ExprContext)._IDENT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).path = append(localctx.(*ExprContext).path, localctx.(*ExprContext)._IDENT)

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLQueryParserIDENT_LITERAL {
		{
			p.SetState(109)

			var _m = p.Match(SQLQueryParserIDENT_LITERAL)

			localctx.(*ExprContext).literalFlag = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLQueryParserIDENT {
			{
				p.SetState(110)

				var _m = p.Match(SQLQueryParserIDENT)

				localctx.(*ExprContext).literalMode = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

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
