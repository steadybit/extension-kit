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
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"AND", "OR", "NOT", "IS", "IN", "PRESENT", "COUNT", "LPAREN", "RPAREN",
		"COMMA", "OP_EQUAL", "OP_NOT_EQUAL", "OP_EQUAL_IGNORE_CASE", "OP_NOT_EQUAL_IGNORE_CASE",
		"OP_TILDE", "OP_NOT_TILDE", "OP_TILDE_IGNORE_CASE", "OP_NOT_TILDE_IGNORE_CASE",
		"OP_GREATER_THAN", "OP_GREATER_THAN_EQUAL", "OP_LESS_THAN", "OP_LESS_THAN_EQUAL",
		"INTEGER", "QUOTED", "VARIABLE", "PLACEHOLDER", "TERM", "WHITESPACE",
		"LINE_COMMENT", "QUOTED_CHAR", "TERM_START_CHAR", "TERM_CHAR", "MARKER_CHAR",
		"ESCAPED_CHAR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 259, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 3, 0, 78, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		86, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 95, 8, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 101, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 107,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 123, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 135, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 5, 22, 180, 8, 22, 10, 22,
		12, 22, 183, 9, 22, 3, 22, 185, 8, 22, 1, 23, 1, 23, 5, 23, 189, 8, 23,
		10, 23, 12, 23, 192, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 4,
		24, 200, 8, 24, 11, 24, 12, 24, 201, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 4, 25, 211, 8, 25, 11, 25, 12, 25, 212, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 5, 26, 220, 8, 26, 10, 26, 12, 26, 223, 9, 26, 1, 27,
		4, 27, 226, 8, 27, 11, 27, 12, 27, 227, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 5, 28, 236, 8, 28, 10, 28, 12, 28, 239, 9, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 3, 29, 245, 8, 29, 1, 30, 1, 30, 3, 30, 249, 8, 30, 1, 31,
		1, 31, 3, 31, 253, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 0, 0, 34,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 0, 61, 0, 63, 0, 65, 0, 67, 0, 1, 0, 8, 1, 0, 49, 57, 1, 0, 48, 57,
		4, 0, 9, 10, 13, 13, 32, 32, 12288, 12288, 2, 0, 10, 10, 13, 13, 2, 0,
		34, 34, 92, 92, 13, 0, 9, 10, 13, 13, 32, 34, 40, 41, 43, 45, 47, 47, 58,
		58, 60, 62, 64, 64, 91, 94, 123, 123, 125, 126, 12288, 12288, 2, 0, 43,
		43, 45, 45, 5, 0, 45, 46, 48, 57, 65, 90, 95, 95, 97, 122, 274, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 77, 1, 0, 0, 0, 3, 85, 1, 0, 0, 0, 5, 94,
		1, 0, 0, 0, 7, 100, 1, 0, 0, 0, 9, 106, 1, 0, 0, 0, 11, 122, 1, 0, 0, 0,
		13, 134, 1, 0, 0, 0, 15, 136, 1, 0, 0, 0, 17, 138, 1, 0, 0, 0, 19, 140,
		1, 0, 0, 0, 21, 142, 1, 0, 0, 0, 23, 144, 1, 0, 0, 0, 25, 147, 1, 0, 0,
		0, 27, 150, 1, 0, 0, 0, 29, 154, 1, 0, 0, 0, 31, 156, 1, 0, 0, 0, 33, 159,
		1, 0, 0, 0, 35, 162, 1, 0, 0, 0, 37, 166, 1, 0, 0, 0, 39, 168, 1, 0, 0,
		0, 41, 171, 1, 0, 0, 0, 43, 173, 1, 0, 0, 0, 45, 184, 1, 0, 0, 0, 47, 186,
		1, 0, 0, 0, 49, 195, 1, 0, 0, 0, 51, 206, 1, 0, 0, 0, 53, 217, 1, 0, 0,
		0, 55, 225, 1, 0, 0, 0, 57, 231, 1, 0, 0, 0, 59, 244, 1, 0, 0, 0, 61, 248,
		1, 0, 0, 0, 63, 252, 1, 0, 0, 0, 65, 254, 1, 0, 0, 0, 67, 256, 1, 0, 0,
		0, 69, 70, 5, 65, 0, 0, 70, 71, 5, 78, 0, 0, 71, 78, 5, 68, 0, 0, 72, 73,
		5, 97, 0, 0, 73, 74, 5, 110, 0, 0, 74, 78, 5, 100, 0, 0, 75, 76, 5, 38,
		0, 0, 76, 78, 5, 38, 0, 0, 77, 69, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 77,
		75, 1, 0, 0, 0, 78, 2, 1, 0, 0, 0, 79, 80, 5, 79, 0, 0, 80, 86, 5, 82,
		0, 0, 81, 82, 5, 111, 0, 0, 82, 86, 5, 114, 0, 0, 83, 84, 5, 124, 0, 0,
		84, 86, 5, 124, 0, 0, 85, 79, 1, 0, 0, 0, 85, 81, 1, 0, 0, 0, 85, 83, 1,
		0, 0, 0, 86, 4, 1, 0, 0, 0, 87, 88, 5, 78, 0, 0, 88, 89, 5, 79, 0, 0, 89,
		95, 5, 84, 0, 0, 90, 91, 5, 110, 0, 0, 91, 92, 5, 111, 0, 0, 92, 95, 5,
		116, 0, 0, 93, 95, 5, 33, 0, 0, 94, 87, 1, 0, 0, 0, 94, 90, 1, 0, 0, 0,
		94, 93, 1, 0, 0, 0, 95, 6, 1, 0, 0, 0, 96, 97, 5, 73, 0, 0, 97, 101, 5,
		83, 0, 0, 98, 99, 5, 105, 0, 0, 99, 101, 5, 115, 0, 0, 100, 96, 1, 0, 0,
		0, 100, 98, 1, 0, 0, 0, 101, 8, 1, 0, 0, 0, 102, 103, 5, 73, 0, 0, 103,
		107, 5, 78, 0, 0, 104, 105, 5, 105, 0, 0, 105, 107, 5, 110, 0, 0, 106,
		102, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 10, 1, 0, 0, 0, 108, 109, 5,
		80, 0, 0, 109, 110, 5, 82, 0, 0, 110, 111, 5, 69, 0, 0, 111, 112, 5, 83,
		0, 0, 112, 113, 5, 69, 0, 0, 113, 114, 5, 78, 0, 0, 114, 123, 5, 84, 0,
		0, 115, 116, 5, 112, 0, 0, 116, 117, 5, 114, 0, 0, 117, 118, 5, 101, 0,
		0, 118, 119, 5, 115, 0, 0, 119, 120, 5, 101, 0, 0, 120, 121, 5, 110, 0,
		0, 121, 123, 5, 116, 0, 0, 122, 108, 1, 0, 0, 0, 122, 115, 1, 0, 0, 0,
		123, 12, 1, 0, 0, 0, 124, 125, 5, 67, 0, 0, 125, 126, 5, 79, 0, 0, 126,
		127, 5, 85, 0, 0, 127, 128, 5, 78, 0, 0, 128, 135, 5, 84, 0, 0, 129, 130,
		5, 99, 0, 0, 130, 131, 5, 111, 0, 0, 131, 132, 5, 117, 0, 0, 132, 133,
		5, 110, 0, 0, 133, 135, 5, 116, 0, 0, 134, 124, 1, 0, 0, 0, 134, 129, 1,
		0, 0, 0, 135, 14, 1, 0, 0, 0, 136, 137, 5, 40, 0, 0, 137, 16, 1, 0, 0,
		0, 138, 139, 5, 41, 0, 0, 139, 18, 1, 0, 0, 0, 140, 141, 5, 44, 0, 0, 141,
		20, 1, 0, 0, 0, 142, 143, 5, 61, 0, 0, 143, 22, 1, 0, 0, 0, 144, 145, 5,
		33, 0, 0, 145, 146, 5, 61, 0, 0, 146, 24, 1, 0, 0, 0, 147, 148, 5, 61,
		0, 0, 148, 149, 5, 42, 0, 0, 149, 26, 1, 0, 0, 0, 150, 151, 5, 33, 0, 0,
		151, 152, 5, 61, 0, 0, 152, 153, 5, 42, 0, 0, 153, 28, 1, 0, 0, 0, 154,
		155, 5, 126, 0, 0, 155, 30, 1, 0, 0, 0, 156, 157, 5, 33, 0, 0, 157, 158,
		5, 126, 0, 0, 158, 32, 1, 0, 0, 0, 159, 160, 5, 126, 0, 0, 160, 161, 5,
		42, 0, 0, 161, 34, 1, 0, 0, 0, 162, 163, 5, 33, 0, 0, 163, 164, 5, 126,
		0, 0, 164, 165, 5, 42, 0, 0, 165, 36, 1, 0, 0, 0, 166, 167, 5, 62, 0, 0,
		167, 38, 1, 0, 0, 0, 168, 169, 5, 62, 0, 0, 169, 170, 5, 61, 0, 0, 170,
		40, 1, 0, 0, 0, 171, 172, 5, 60, 0, 0, 172, 42, 1, 0, 0, 0, 173, 174, 5,
		60, 0, 0, 174, 175, 5, 61, 0, 0, 175, 44, 1, 0, 0, 0, 176, 185, 5, 48,
		0, 0, 177, 181, 7, 0, 0, 0, 178, 180, 7, 1, 0, 0, 179, 178, 1, 0, 0, 0,
		180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182,
		185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 176, 1, 0, 0, 0, 184, 177,
		1, 0, 0, 0, 185, 46, 1, 0, 0, 0, 186, 190, 5, 34, 0, 0, 187, 189, 3, 59,
		29, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0,
		190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193,
		194, 5, 34, 0, 0, 194, 48, 1, 0, 0, 0, 195, 196, 5, 123, 0, 0, 196, 197,
		5, 123, 0, 0, 197, 199, 1, 0, 0, 0, 198, 200, 3, 65, 32, 0, 199, 198, 1,
		0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0,
		0, 202, 203, 1, 0, 0, 0, 203, 204, 5, 125, 0, 0, 204, 205, 5, 125, 0, 0,
		205, 50, 1, 0, 0, 0, 206, 207, 5, 91, 0, 0, 207, 208, 5, 91, 0, 0, 208,
		210, 1, 0, 0, 0, 209, 211, 3, 65, 32, 0, 210, 209, 1, 0, 0, 0, 211, 212,
		1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0,
		0, 0, 214, 215, 5, 93, 0, 0, 215, 216, 5, 93, 0, 0, 216, 52, 1, 0, 0, 0,
		217, 221, 3, 61, 30, 0, 218, 220, 3, 63, 31, 0, 219, 218, 1, 0, 0, 0, 220,
		223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 54, 1,
		0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 226, 7, 2, 0, 0, 225, 224, 1, 0, 0,
		0, 226, 227, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228,
		229, 1, 0, 0, 0, 229, 230, 6, 27, 0, 0, 230, 56, 1, 0, 0, 0, 231, 232,
		5, 47, 0, 0, 232, 233, 5, 47, 0, 0, 233, 237, 1, 0, 0, 0, 234, 236, 8,
		3, 0, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0, 0,
		0, 237, 238, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 240,
		241, 6, 28, 0, 0, 241, 58, 1, 0, 0, 0, 242, 245, 8, 4, 0, 0, 243, 245,
		3, 67, 33, 0, 244, 242, 1, 0, 0, 0, 244, 243, 1, 0, 0, 0, 245, 60, 1, 0,
		0, 0, 246, 249, 8, 5, 0, 0, 247, 249, 3, 67, 33, 0, 248, 246, 1, 0, 0,
		0, 248, 247, 1, 0, 0, 0, 249, 62, 1, 0, 0, 0, 250, 253, 3, 61, 30, 0, 251,
		253, 7, 6, 0, 0, 252, 250, 1, 0, 0, 0, 252, 251, 1, 0, 0, 0, 253, 64, 1,
		0, 0, 0, 254, 255, 7, 7, 0, 0, 255, 66, 1, 0, 0, 0, 256, 257, 5, 92, 0,
		0, 257, 258, 9, 0, 0, 0, 258, 68, 1, 0, 0, 0, 19, 0, 77, 85, 94, 100, 106,
		122, 134, 181, 184, 190, 201, 212, 221, 227, 237, 244, 248, 252, 1, 0,
		1, 0,
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
	QueryLanguageLexerLINE_COMMENT             = 29
)
