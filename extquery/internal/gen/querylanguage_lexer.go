// Code generated from QueryLanguageLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen

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

type QueryLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var QueryLanguageLexerLexerStaticData struct {
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

func querylanguagelexerLexerInit() {
	staticData := &QueryLanguageLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"AND", "OR", "NOT", "IS", "IN", "PRESENT", "COUNT", "LPAREN", "RPAREN",
		"COMMA", "OP_EQUAL", "OP_NOT_EQUAL", "OP_EQUAL_IGNORE_CASE", "OP_NOT_EQUAL_IGNORE_CASE",
		"OP_TILDE", "OP_NOT_TILDE", "OP_TILDE_IGNORE_CASE", "OP_NOT_TILDE_IGNORE_CASE",
		"OP_GREATER_THAN", "OP_GREATER_THAN_EQUAL", "OP_LESS_THAN", "OP_LESS_THAN_EQUAL",
		"INTEGER", "QUOTED", "VARIABLE", "PLACEHOLDER", "TERM", "WHITESPACE",
		"QUOTED_CHAR", "TERM_START_CHAR", "TERM_CHAR", "MARKER_CHAR", "ESCAPED_CHAR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 246, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		3, 0, 76, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 84, 8, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 93, 8, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 99, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 105, 8, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 121, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 133, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 5, 22, 178, 8, 22, 10, 22, 12, 22, 181, 9,
		22, 3, 22, 183, 8, 22, 1, 23, 1, 23, 5, 23, 187, 8, 23, 10, 23, 12, 23,
		190, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 4, 24, 198, 8, 24,
		11, 24, 12, 24, 199, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 4,
		25, 209, 8, 25, 11, 25, 12, 25, 210, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		5, 26, 218, 8, 26, 10, 26, 12, 26, 221, 9, 26, 1, 27, 4, 27, 224, 8, 27,
		11, 27, 12, 27, 225, 1, 27, 1, 27, 1, 28, 1, 28, 3, 28, 232, 8, 28, 1,
		29, 1, 29, 3, 29, 236, 8, 29, 1, 30, 1, 30, 3, 30, 240, 8, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 0, 0, 33, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 28, 57, 0, 59, 0, 61, 0, 63, 0, 65, 0, 1, 0, 7,
		1, 0, 49, 57, 1, 0, 48, 57, 4, 0, 9, 10, 13, 13, 32, 32, 12288, 12288,
		2, 0, 34, 34, 92, 92, 13, 0, 9, 10, 13, 13, 32, 34, 40, 41, 43, 45, 47,
		47, 58, 58, 60, 62, 64, 64, 91, 94, 123, 123, 125, 126, 12288, 12288, 2,
		0, 43, 43, 45, 45, 5, 0, 45, 46, 48, 57, 65, 90, 95, 95, 97, 122, 260,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 1, 75, 1, 0, 0, 0, 3, 83, 1, 0, 0, 0, 5, 92, 1, 0, 0,
		0, 7, 98, 1, 0, 0, 0, 9, 104, 1, 0, 0, 0, 11, 120, 1, 0, 0, 0, 13, 132,
		1, 0, 0, 0, 15, 134, 1, 0, 0, 0, 17, 136, 1, 0, 0, 0, 19, 138, 1, 0, 0,
		0, 21, 140, 1, 0, 0, 0, 23, 142, 1, 0, 0, 0, 25, 145, 1, 0, 0, 0, 27, 148,
		1, 0, 0, 0, 29, 152, 1, 0, 0, 0, 31, 154, 1, 0, 0, 0, 33, 157, 1, 0, 0,
		0, 35, 160, 1, 0, 0, 0, 37, 164, 1, 0, 0, 0, 39, 166, 1, 0, 0, 0, 41, 169,
		1, 0, 0, 0, 43, 171, 1, 0, 0, 0, 45, 182, 1, 0, 0, 0, 47, 184, 1, 0, 0,
		0, 49, 193, 1, 0, 0, 0, 51, 204, 1, 0, 0, 0, 53, 215, 1, 0, 0, 0, 55, 223,
		1, 0, 0, 0, 57, 231, 1, 0, 0, 0, 59, 235, 1, 0, 0, 0, 61, 239, 1, 0, 0,
		0, 63, 241, 1, 0, 0, 0, 65, 243, 1, 0, 0, 0, 67, 68, 5, 65, 0, 0, 68, 69,
		5, 78, 0, 0, 69, 76, 5, 68, 0, 0, 70, 71, 5, 97, 0, 0, 71, 72, 5, 110,
		0, 0, 72, 76, 5, 100, 0, 0, 73, 74, 5, 38, 0, 0, 74, 76, 5, 38, 0, 0, 75,
		67, 1, 0, 0, 0, 75, 70, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 2, 1, 0, 0,
		0, 77, 78, 5, 79, 0, 0, 78, 84, 5, 82, 0, 0, 79, 80, 5, 111, 0, 0, 80,
		84, 5, 114, 0, 0, 81, 82, 5, 124, 0, 0, 82, 84, 5, 124, 0, 0, 83, 77, 1,
		0, 0, 0, 83, 79, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 4, 1, 0, 0, 0, 85,
		86, 5, 78, 0, 0, 86, 87, 5, 79, 0, 0, 87, 93, 5, 84, 0, 0, 88, 89, 5, 110,
		0, 0, 89, 90, 5, 111, 0, 0, 90, 93, 5, 116, 0, 0, 91, 93, 5, 33, 0, 0,
		92, 85, 1, 0, 0, 0, 92, 88, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 6, 1, 0,
		0, 0, 94, 95, 5, 73, 0, 0, 95, 99, 5, 83, 0, 0, 96, 97, 5, 105, 0, 0, 97,
		99, 5, 115, 0, 0, 98, 94, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 8, 1, 0,
		0, 0, 100, 101, 5, 73, 0, 0, 101, 105, 5, 78, 0, 0, 102, 103, 5, 105, 0,
		0, 103, 105, 5, 110, 0, 0, 104, 100, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0,
		105, 10, 1, 0, 0, 0, 106, 107, 5, 80, 0, 0, 107, 108, 5, 82, 0, 0, 108,
		109, 5, 69, 0, 0, 109, 110, 5, 83, 0, 0, 110, 111, 5, 69, 0, 0, 111, 112,
		5, 78, 0, 0, 112, 121, 5, 84, 0, 0, 113, 114, 5, 112, 0, 0, 114, 115, 5,
		114, 0, 0, 115, 116, 5, 101, 0, 0, 116, 117, 5, 115, 0, 0, 117, 118, 5,
		101, 0, 0, 118, 119, 5, 110, 0, 0, 119, 121, 5, 116, 0, 0, 120, 106, 1,
		0, 0, 0, 120, 113, 1, 0, 0, 0, 121, 12, 1, 0, 0, 0, 122, 123, 5, 67, 0,
		0, 123, 124, 5, 79, 0, 0, 124, 125, 5, 85, 0, 0, 125, 126, 5, 78, 0, 0,
		126, 133, 5, 84, 0, 0, 127, 128, 5, 99, 0, 0, 128, 129, 5, 111, 0, 0, 129,
		130, 5, 117, 0, 0, 130, 131, 5, 110, 0, 0, 131, 133, 5, 116, 0, 0, 132,
		122, 1, 0, 0, 0, 132, 127, 1, 0, 0, 0, 133, 14, 1, 0, 0, 0, 134, 135, 5,
		40, 0, 0, 135, 16, 1, 0, 0, 0, 136, 137, 5, 41, 0, 0, 137, 18, 1, 0, 0,
		0, 138, 139, 5, 44, 0, 0, 139, 20, 1, 0, 0, 0, 140, 141, 5, 61, 0, 0, 141,
		22, 1, 0, 0, 0, 142, 143, 5, 33, 0, 0, 143, 144, 5, 61, 0, 0, 144, 24,
		1, 0, 0, 0, 145, 146, 5, 61, 0, 0, 146, 147, 5, 42, 0, 0, 147, 26, 1, 0,
		0, 0, 148, 149, 5, 33, 0, 0, 149, 150, 5, 61, 0, 0, 150, 151, 5, 42, 0,
		0, 151, 28, 1, 0, 0, 0, 152, 153, 5, 126, 0, 0, 153, 30, 1, 0, 0, 0, 154,
		155, 5, 33, 0, 0, 155, 156, 5, 126, 0, 0, 156, 32, 1, 0, 0, 0, 157, 158,
		5, 126, 0, 0, 158, 159, 5, 42, 0, 0, 159, 34, 1, 0, 0, 0, 160, 161, 5,
		33, 0, 0, 161, 162, 5, 126, 0, 0, 162, 163, 5, 42, 0, 0, 163, 36, 1, 0,
		0, 0, 164, 165, 5, 62, 0, 0, 165, 38, 1, 0, 0, 0, 166, 167, 5, 62, 0, 0,
		167, 168, 5, 61, 0, 0, 168, 40, 1, 0, 0, 0, 169, 170, 5, 60, 0, 0, 170,
		42, 1, 0, 0, 0, 171, 172, 5, 60, 0, 0, 172, 173, 5, 61, 0, 0, 173, 44,
		1, 0, 0, 0, 174, 183, 5, 48, 0, 0, 175, 179, 7, 0, 0, 0, 176, 178, 7, 1,
		0, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0,
		179, 180, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182,
		174, 1, 0, 0, 0, 182, 175, 1, 0, 0, 0, 183, 46, 1, 0, 0, 0, 184, 188, 5,
		34, 0, 0, 185, 187, 3, 57, 28, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0,
		0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0,
		190, 188, 1, 0, 0, 0, 191, 192, 5, 34, 0, 0, 192, 48, 1, 0, 0, 0, 193,
		194, 5, 123, 0, 0, 194, 195, 5, 123, 0, 0, 195, 197, 1, 0, 0, 0, 196, 198,
		3, 63, 31, 0, 197, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 197, 1,
		0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 5, 125,
		0, 0, 202, 203, 5, 125, 0, 0, 203, 50, 1, 0, 0, 0, 204, 205, 5, 91, 0,
		0, 205, 206, 5, 91, 0, 0, 206, 208, 1, 0, 0, 0, 207, 209, 3, 63, 31, 0,
		208, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210,
		211, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 5, 93, 0, 0, 213, 214,
		5, 93, 0, 0, 214, 52, 1, 0, 0, 0, 215, 219, 3, 59, 29, 0, 216, 218, 3,
		61, 30, 0, 217, 216, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0,
		0, 0, 219, 220, 1, 0, 0, 0, 220, 54, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0,
		222, 224, 7, 2, 0, 0, 223, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225,
		223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228,
		6, 27, 0, 0, 228, 56, 1, 0, 0, 0, 229, 232, 8, 3, 0, 0, 230, 232, 3, 65,
		32, 0, 231, 229, 1, 0, 0, 0, 231, 230, 1, 0, 0, 0, 232, 58, 1, 0, 0, 0,
		233, 236, 8, 4, 0, 0, 234, 236, 3, 65, 32, 0, 235, 233, 1, 0, 0, 0, 235,
		234, 1, 0, 0, 0, 236, 60, 1, 0, 0, 0, 237, 240, 3, 59, 29, 0, 238, 240,
		7, 5, 0, 0, 239, 237, 1, 0, 0, 0, 239, 238, 1, 0, 0, 0, 240, 62, 1, 0,
		0, 0, 241, 242, 7, 6, 0, 0, 242, 64, 1, 0, 0, 0, 243, 244, 5, 92, 0, 0,
		244, 245, 9, 0, 0, 0, 245, 66, 1, 0, 0, 0, 18, 0, 75, 83, 92, 98, 104,
		120, 132, 179, 182, 188, 199, 210, 219, 225, 231, 235, 239, 1, 0, 1, 0,
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

// QueryLanguageLexerInit initializes any static state used to implement QueryLanguageLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQueryLanguageLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLanguageLexerInit() {
	staticData := &QueryLanguageLexerLexerStaticData
	staticData.once.Do(querylanguagelexerLexerInit)
}

// NewQueryLanguageLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQueryLanguageLexer(input antlr.CharStream) *QueryLanguageLexer {
	QueryLanguageLexerInit()
	l := new(QueryLanguageLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &QueryLanguageLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "QueryLanguageLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLanguageLexer tokens.
const (
	QueryLanguageLexerAND                      = 1
	QueryLanguageLexerOR                       = 2
	QueryLanguageLexerNOT                      = 3
	QueryLanguageLexerIS                       = 4
	QueryLanguageLexerIN                       = 5
	QueryLanguageLexerPRESENT                  = 6
	QueryLanguageLexerCOUNT                    = 7
	QueryLanguageLexerLPAREN                   = 8
	QueryLanguageLexerRPAREN                   = 9
	QueryLanguageLexerCOMMA                    = 10
	QueryLanguageLexerOP_EQUAL                 = 11
	QueryLanguageLexerOP_NOT_EQUAL             = 12
	QueryLanguageLexerOP_EQUAL_IGNORE_CASE     = 13
	QueryLanguageLexerOP_NOT_EQUAL_IGNORE_CASE = 14
	QueryLanguageLexerOP_TILDE                 = 15
	QueryLanguageLexerOP_NOT_TILDE             = 16
	QueryLanguageLexerOP_TILDE_IGNORE_CASE     = 17
	QueryLanguageLexerOP_NOT_TILDE_IGNORE_CASE = 18
	QueryLanguageLexerOP_GREATER_THAN          = 19
	QueryLanguageLexerOP_GREATER_THAN_EQUAL    = 20
	QueryLanguageLexerOP_LESS_THAN             = 21
	QueryLanguageLexerOP_LESS_THAN_EQUAL       = 22
	QueryLanguageLexerINTEGER                  = 23
	QueryLanguageLexerQUOTED                   = 24
	QueryLanguageLexerVARIABLE                 = 25
	QueryLanguageLexerPLACEHOLDER              = 26
	QueryLanguageLexerTERM                     = 27
	QueryLanguageLexerWHITESPACE               = 28
)
