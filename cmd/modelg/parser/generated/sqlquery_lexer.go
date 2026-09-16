// Code generated from SQLQueryLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated

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

type SQLQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SQLQueryLexerLexerStaticData struct {
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

func sqlquerylexerLexerInit() {
	staticData := &SQLQueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "Comment", "Query", "Expr",
	}
	staticData.LiteralNames = []string{
		"", "'--'", "", "'--when'", "'--<when'", "'--+when'", "", "':'", "",
		"", "", "", "", "", "", "", "'.'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "STARTCOMMENT", "STARTQUERY", "STARTWHEN", "STARTCHOMPWHEN", "STARTJOINWHEN",
		"ENDWHEN", "STARTEXPR", "EOL", "WS", "SQLTEXT", "RAWCOMMENT", "ENDCOMMENT",
		"QUERY_NAME", "ENDQUERY", "IDENT", "IDENT_SEP", "IDENT_LITERAL", "WS_EXPR",
	}
	staticData.RuleNames = []string{
		"LINEWS", "NL", "STRING_LITERAL", "QUOTED_IDENTIFIER", "STARTCOMMENT",
		"STARTQUERY", "STARTWHEN", "STARTCHOMPWHEN", "STARTJOINWHEN", "ENDWHEN",
		"STARTEXPR", "EOL", "WS", "SQLTEXT", "RAWCOMMENT", "ENDCOMMENT", "QUERY_NAME",
		"ENDQUERY", "IDENT", "IDENT_SEP", "IDENT_LITERAL", "WS_EXPR", "EOL_EXPR",
		"SQLTEXT_EXPR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 236, 6, -1, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2,
		7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8,
		7, 8, 2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13,
		2, 14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2,
		19, 7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0,
		4, 0, 54, 8, 0, 11, 0, 12, 0, 55, 1, 1, 3, 1, 59, 8, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 67, 8, 2, 10, 2, 12, 2, 70, 9, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 78, 8, 3, 10, 3, 12, 3, 81, 9, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 95,
		8, 5, 10, 5, 12, 5, 98, 9, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 107, 8, 5, 10, 5, 12, 5, 110, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 148, 8, 9, 10, 9, 12, 9,
		151, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		4, 12, 162, 8, 12, 11, 12, 12, 12, 163, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 4, 13, 175, 8, 13, 11, 13, 12, 13, 176,
		3, 13, 179, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 5, 16, 192, 8, 16, 10, 16, 12, 16, 195, 9, 16,
		1, 17, 5, 17, 198, 8, 17, 10, 17, 12, 17, 201, 9, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 5, 18, 210, 8, 18, 10, 18, 12, 18, 213,
		9, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 3, 23, 232, 8, 23,
		1, 23, 1, 23, 1, 23, 0, 0, 24, 4, 0, 6, 0, 8, 0, 10, 0, 12, 1, 14, 2, 16,
		3, 18, 4, 20, 5, 22, 6, 24, 7, 26, 8, 28, 9, 30, 10, 32, 11, 34, 12, 36,
		13, 38, 14, 40, 15, 42, 16, 44, 17, 46, 18, 48, 0, 50, 0, 4, 0, 1, 2, 3,
		9, 2, 0, 9, 9, 32, 32, 1, 0, 39, 39, 2, 0, 0, 0, 34, 34, 3, 0, 9, 9, 13,
		13, 32, 32, 7, 0, 9, 10, 13, 13, 32, 32, 34, 34, 39, 39, 45, 45, 58, 58,
		3, 0, 65, 90, 95, 95, 97, 122, 1, 0, 45, 45, 2, 0, 10, 10, 13, 13, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 248, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0,
		0, 0, 16, 1, 0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0,
		0, 0, 0, 24, 1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1,
		0, 0, 0, 1, 32, 1, 0, 0, 0, 1, 34, 1, 0, 0, 0, 2, 36, 1, 0, 0, 0, 2, 38,
		1, 0, 0, 0, 3, 40, 1, 0, 0, 0, 3, 42, 1, 0, 0, 0, 3, 44, 1, 0, 0, 0, 3,
		46, 1, 0, 0, 0, 3, 48, 1, 0, 0, 0, 3, 50, 1, 0, 0, 0, 4, 53, 1, 0, 0, 0,
		6, 58, 1, 0, 0, 0, 8, 62, 1, 0, 0, 0, 10, 73, 1, 0, 0, 0, 12, 84, 1, 0,
		0, 0, 14, 90, 1, 0, 0, 0, 16, 113, 1, 0, 0, 0, 18, 120, 1, 0, 0, 0, 20,
		128, 1, 0, 0, 0, 22, 136, 1, 0, 0, 0, 24, 154, 1, 0, 0, 0, 26, 158, 1,
		0, 0, 0, 28, 161, 1, 0, 0, 0, 30, 178, 1, 0, 0, 0, 32, 180, 1, 0, 0, 0,
		34, 184, 1, 0, 0, 0, 36, 189, 1, 0, 0, 0, 38, 199, 1, 0, 0, 0, 40, 207,
		1, 0, 0, 0, 42, 214, 1, 0, 0, 0, 44, 216, 1, 0, 0, 0, 46, 218, 1, 0, 0,
		0, 48, 223, 1, 0, 0, 0, 50, 231, 1, 0, 0, 0, 52, 54, 7, 0, 0, 0, 53, 52,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0,
		56, 5, 1, 0, 0, 0, 57, 59, 5, 13, 0, 0, 58, 57, 1, 0, 0, 0, 58, 59, 1,
		0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 5, 10, 0, 0, 61, 7, 1, 0, 0, 0, 62,
		68, 5, 39, 0, 0, 63, 67, 8, 1, 0, 0, 64, 65, 5, 39, 0, 0, 65, 67, 5, 39,
		0, 0, 66, 63, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66,
		1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0,
		71, 72, 5, 39, 0, 0, 72, 9, 1, 0, 0, 0, 73, 79, 5, 34, 0, 0, 74, 75, 5,
		34, 0, 0, 75, 78, 5, 34, 0, 0, 76, 78, 8, 2, 0, 0, 77, 74, 1, 0, 0, 0,
		77, 76, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1,
		0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 83, 5, 34, 0, 0, 83,
		11, 1, 0, 0, 0, 84, 85, 5, 45, 0, 0, 85, 86, 5, 45, 0, 0, 86, 87, 1, 0,
		0, 0, 87, 88, 6, 4, 0, 0, 88, 89, 6, 4, 1, 0, 89, 13, 1, 0, 0, 0, 90, 91,
		5, 45, 0, 0, 91, 92, 5, 45, 0, 0, 92, 96, 1, 0, 0, 0, 93, 95, 3, 4, 0,
		0, 94, 93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97,
		1, 0, 0, 0, 97, 99, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 100, 5, 110, 0,
		0, 100, 101, 5, 97, 0, 0, 101, 102, 5, 109, 0, 0, 102, 103, 5, 101, 0,
		0, 103, 104, 5, 58, 0, 0, 104, 108, 1, 0, 0, 0, 105, 107, 3, 4, 0, 0, 106,
		105, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 111, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 112, 6, 5,
		2, 0, 112, 15, 1, 0, 0, 0, 113, 114, 5, 45, 0, 0, 114, 115, 5, 45, 0, 0,
		115, 116, 5, 119, 0, 0, 116, 117, 5, 104, 0, 0, 117, 118, 5, 101, 0, 0,
		118, 119, 5, 110, 0, 0, 119, 17, 1, 0, 0, 0, 120, 121, 5, 45, 0, 0, 121,
		122, 5, 45, 0, 0, 122, 123, 5, 60, 0, 0, 123, 124, 5, 119, 0, 0, 124, 125,
		5, 104, 0, 0, 125, 126, 5, 101, 0, 0, 126, 127, 5, 110, 0, 0, 127, 19,
		1, 0, 0, 0, 128, 129, 5, 45, 0, 0, 129, 130, 5, 45, 0, 0, 130, 131, 5,
		43, 0, 0, 131, 132, 5, 119, 0, 0, 132, 133, 5, 104, 0, 0, 133, 134, 5,
		101, 0, 0, 134, 135, 5, 110, 0, 0, 135, 21, 1, 0, 0, 0, 136, 137, 5, 45,
		0, 0, 137, 138, 5, 45, 0, 0, 138, 139, 5, 101, 0, 0, 139, 140, 5, 110,
		0, 0, 140, 141, 5, 100, 0, 0, 141, 142, 5, 119, 0, 0, 142, 143, 5, 104,
		0, 0, 143, 144, 5, 101, 0, 0, 144, 145, 5, 110, 0, 0, 145, 149, 1, 0, 0,
		0, 146, 148, 3, 4, 0, 0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149,
		147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 152, 153, 3, 6, 1, 0, 153, 23, 1, 0, 0, 0, 154, 155, 5, 58,
		0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 6, 10, 3, 0, 157, 25, 1, 0, 0, 0,
		158, 159, 3, 6, 1, 0, 159, 27, 1, 0, 0, 0, 160, 162, 7, 3, 0, 0, 161, 160,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0,
		0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 6, 12, 0, 0, 166, 29, 1, 0, 0, 0,
		167, 179, 3, 8, 2, 0, 168, 179, 3, 10, 3, 0, 169, 175, 8, 4, 0, 0, 170,
		171, 5, 58, 0, 0, 171, 175, 8, 5, 0, 0, 172, 173, 5, 45, 0, 0, 173, 175,
		8, 6, 0, 0, 174, 169, 1, 0, 0, 0, 174, 170, 1, 0, 0, 0, 174, 172, 1, 0,
		0, 0, 175, 176, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0,
		177, 179, 1, 0, 0, 0, 178, 167, 1, 0, 0, 0, 178, 168, 1, 0, 0, 0, 178,
		174, 1, 0, 0, 0, 179, 31, 1, 0, 0, 0, 180, 181, 8, 7, 0, 0, 181, 182, 1,
		0, 0, 0, 182, 183, 6, 14, 0, 0, 183, 33, 1, 0, 0, 0, 184, 185, 3, 6, 1,
		0, 185, 186, 1, 0, 0, 0, 186, 187, 6, 15, 0, 0, 187, 188, 6, 15, 4, 0,
		188, 35, 1, 0, 0, 0, 189, 193, 7, 5, 0, 0, 190, 192, 7, 8, 0, 0, 191, 190,
		1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0,
		0, 0, 194, 37, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 198, 3, 4, 0, 0,
		197, 196, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 203,
		3, 6, 1, 0, 203, 204, 1, 0, 0, 0, 204, 205, 6, 17, 0, 0, 205, 206, 6, 17,
		4, 0, 206, 39, 1, 0, 0, 0, 207, 211, 7, 5, 0, 0, 208, 210, 7, 8, 0, 0,
		209, 208, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211,
		212, 1, 0, 0, 0, 212, 41, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 214, 215, 5,
		46, 0, 0, 215, 43, 1, 0, 0, 0, 216, 217, 5, 33, 0, 0, 217, 45, 1, 0, 0,
		0, 218, 219, 3, 28, 12, 0, 219, 220, 1, 0, 0, 0, 220, 221, 6, 21, 0, 0,
		221, 222, 6, 21, 4, 0, 222, 47, 1, 0, 0, 0, 223, 224, 3, 6, 1, 0, 224,
		225, 1, 0, 0, 0, 225, 226, 6, 22, 5, 0, 226, 227, 6, 22, 4, 0, 227, 49,
		1, 0, 0, 0, 228, 232, 3, 8, 2, 0, 229, 232, 3, 10, 3, 0, 230, 232, 9, 0,
		0, 0, 231, 228, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 230, 1, 0, 0, 0,
		232, 233, 1, 0, 0, 0, 233, 234, 6, 23, 6, 0, 234, 235, 6, 23, 4, 0, 235,
		51, 1, 0, 0, 0, 21, 0, 1, 2, 3, 55, 58, 66, 68, 77, 79, 96, 108, 149, 163,
		174, 176, 178, 193, 199, 211, 231, 7, 6, 0, 0, 5, 1, 0, 5, 2, 0, 5, 3,
		0, 4, 0, 0, 7, 8, 0, 7, 10, 0,
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

// SQLQueryLexerInit initializes any static state used to implement SQLQueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSQLQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SQLQueryLexerInit() {
	staticData := &SQLQueryLexerLexerStaticData
	staticData.once.Do(sqlquerylexerLexerInit)
}

// NewSQLQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSQLQueryLexer(input antlr.CharStream) *SQLQueryLexer {
	SQLQueryLexerInit()
	l := new(SQLQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SQLQueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SQLQueryLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SQLQueryLexer tokens.
const (
	SQLQueryLexerSTARTCOMMENT   = 1
	SQLQueryLexerSTARTQUERY     = 2
	SQLQueryLexerSTARTWHEN      = 3
	SQLQueryLexerSTARTCHOMPWHEN = 4
	SQLQueryLexerSTARTJOINWHEN  = 5
	SQLQueryLexerENDWHEN        = 6
	SQLQueryLexerSTARTEXPR      = 7
	SQLQueryLexerEOL            = 8
	SQLQueryLexerWS             = 9
	SQLQueryLexerSQLTEXT        = 10
	SQLQueryLexerRAWCOMMENT     = 11
	SQLQueryLexerENDCOMMENT     = 12
	SQLQueryLexerQUERY_NAME     = 13
	SQLQueryLexerENDQUERY       = 14
	SQLQueryLexerIDENT          = 15
	SQLQueryLexerIDENT_SEP      = 16
	SQLQueryLexerIDENT_LITERAL  = 17
	SQLQueryLexerWS_EXPR        = 18
)

// SQLQueryLexer modes.
const (
	SQLQueryLexerComment = iota + 1
	SQLQueryLexerQuery
	SQLQueryLexerExpr
)
