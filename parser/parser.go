// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2

//line parser.go.y:28
type yySymType struct {
	yys   int
	token Token
	expr  Node
}

const TDo = 57346
const TLocal = 57347
const TElseIf = 57348
const TThen = 57349
const TEnd = 57350
const TBreak = 57351
const TElse = 57352
const TFor = 57353
const TWhile = 57354
const TFunc = 57355
const TIf = 57356
const TReturn = 57357
const TReturnVoid = 57358
const TRepeat = 57359
const TUntil = 57360
const TNot = 57361
const TLabel = 57362
const TGoto = 57363
const TIn = 57364
const TOr = 57365
const TAnd = 57366
const TEqeq = 57367
const TNeq = 57368
const TLte = 57369
const TGte = 57370
const TIdent = 57371
const TNumber = 57372
const TString = 57373
const TIDiv = 57374
const TAddEq = 57375
const TSubEq = 57376
const TMulEq = 57377
const TDivEq = 57378
const TModEq = 57379
const TSquare = 57380
const TDotDot = 57381
const ASSIGN = 57382
const FUNC = 57383
const UNARY = 57384

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TDo",
	"TLocal",
	"TElseIf",
	"TThen",
	"TEnd",
	"TBreak",
	"TElse",
	"TFor",
	"TWhile",
	"TFunc",
	"TIf",
	"TReturn",
	"TReturnVoid",
	"TRepeat",
	"TUntil",
	"TNot",
	"TLabel",
	"TGoto",
	"TIn",
	"TOr",
	"TAnd",
	"TEqeq",
	"TNeq",
	"TLte",
	"TGte",
	"TIdent",
	"TNumber",
	"TString",
	"'{'",
	"'['",
	"'('",
	"'='",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'#'",
	"'.'",
	"'&'",
	"TIDiv",
	"TAddEq",
	"TSubEq",
	"TMulEq",
	"TDivEq",
	"TModEq",
	"TSquare",
	"TDotDot",
	"'T'",
	"ASSIGN",
	"FUNC",
	"UNARY",
	"';'",
	"']'",
	"','",
	"')'",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:365

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	35, 55,
	61, 55,
	-2, 81,
	-1, 99,
	35, 56,
	61, 56,
	-2, 81,
}

const yyPrivate = 57344

const yyLast = 1038

var yyAct = [...]int{

	33, 36, 47, 16, 174, 91, 173, 32, 59, 37,
	113, 162, 162, 182, 181, 163, 163, 97, 169, 94,
	34, 35, 53, 95, 26, 57, 56, 58, 49, 38,
	67, 16, 198, 54, 39, 40, 96, 25, 86, 87,
	88, 141, 142, 89, 192, 161, 140, 139, 183, 97,
	179, 143, 93, 100, 92, 112, 16, 102, 98, 41,
	116, 149, 97, 67, 26, 25, 50, 148, 117, 118,
	119, 120, 121, 122, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 48, 99, 49, 107,
	25, 4, 51, 79, 80, 82, 144, 146, 66, 138,
	81, 106, 49, 3, 145, 150, 152, 76, 78, 79,
	80, 82, 83, 154, 135, 44, 81, 84, 46, 4,
	108, 155, 5, 77, 60, 61, 62, 63, 64, 85,
	90, 3, 55, 52, 31, 147, 30, 151, 159, 160,
	101, 187, 117, 153, 165, 158, 6, 164, 156, 16,
	5, 37, 1, 16, 44, 16, 42, 46, 16, 177,
	29, 94, 34, 35, 180, 95, 26, 20, 43, 15,
	16, 38, 14, 185, 17, 45, 39, 189, 19, 104,
	190, 2, 168, 25, 193, 0, 0, 25, 16, 25,
	16, 0, 25, 0, 0, 0, 200, 0, 0, 202,
	0, 16, 0, 16, 25, 0, 205, 186, 16, 188,
	76, 78, 79, 80, 82, 83, 0, 0, 0, 81,
	0, 0, 25, 0, 25, 0, 0, 0, 199, 0,
	0, 201, 0, 0, 0, 25, 0, 25, 206, 0,
	0, 0, 25, 68, 69, 74, 75, 73, 72, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 76, 78,
	79, 80, 82, 83, 0, 0, 0, 81, 0, 0,
	0, 0, 0, 0, 77, 0, 0, 0, 0, 0,
	157, 0, 0, 137, 68, 69, 74, 75, 73, 72,
	0, 0, 0, 0, 0, 0, 0, 70, 71, 76,
	78, 79, 80, 82, 83, 0, 194, 0, 81, 0,
	0, 0, 0, 0, 0, 77, 0, 0, 0, 0,
	0, 136, 0, 0, 137, 68, 69, 74, 75, 73,
	72, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	76, 78, 79, 80, 82, 83, 0, 0, 0, 81,
	0, 0, 0, 0, 0, 0, 77, 68, 69, 74,
	75, 73, 72, 195, 0, 0, 0, 0, 0, 0,
	70, 71, 76, 78, 79, 80, 82, 83, 0, 0,
	0, 81, 0, 0, 0, 0, 0, 0, 77, 68,
	69, 74, 75, 73, 72, 0, 115, 0, 0, 0,
	0, 0, 70, 71, 76, 78, 79, 80, 82, 83,
	0, 0, 0, 81, 0, 0, 0, 0, 0, 0,
	77, 68, 69, 74, 75, 73, 72, 172, 0, 0,
	0, 0, 0, 0, 70, 71, 76, 78, 79, 80,
	82, 83, 0, 0, 0, 81, 68, 69, 74, 75,
	73, 72, 77, 0, 0, 0, 0, 0, 191, 70,
	71, 76, 78, 79, 80, 82, 83, 0, 0, 0,
	81, 68, 69, 74, 75, 73, 72, 77, 0, 0,
	0, 0, 0, 178, 70, 71, 76, 78, 79, 80,
	82, 83, 0, 0, 0, 81, 0, 0, 0, 109,
	111, 176, 77, 0, 9, 175, 23, 21, 166, 24,
	13, 12, 22, 0, 0, 11, 10, 0, 0, 0,
	0, 0, 0, 0, 27, 0, 7, 18, 0, 26,
	65, 9, 0, 23, 21, 28, 24, 13, 12, 22,
	0, 0, 11, 10, 0, 0, 0, 0, 0, 0,
	0, 27, 0, 0, 110, 0, 26, 109, 111, 0,
	0, 207, 9, 0, 23, 21, 0, 24, 13, 12,
	22, 0, 0, 11, 10, 0, 0, 0, 0, 0,
	0, 8, 27, 0, 109, 111, 0, 26, 203, 9,
	0, 23, 21, 37, 24, 13, 12, 22, 0, 0,
	11, 10, 0, 41, 34, 35, 0, 0, 26, 27,
	0, 0, 110, 38, 26, 109, 111, 0, 39, 196,
	9, 0, 23, 21, 0, 24, 13, 12, 22, 0,
	0, 11, 10, 0, 0, 0, 0, 0, 0, 110,
	27, 0, 109, 111, 0, 26, 184, 9, 0, 23,
	21, 0, 24, 13, 12, 22, 0, 0, 11, 10,
	0, 0, 0, 0, 0, 0, 0, 27, 0, 0,
	110, 0, 26, 109, 111, 0, 0, 171, 9, 0,
	23, 21, 0, 24, 13, 12, 22, 0, 0, 11,
	10, 0, 0, 0, 0, 0, 0, 110, 27, 0,
	109, 111, 0, 26, 170, 9, 0, 23, 21, 0,
	24, 13, 12, 22, 0, 0, 11, 10, 0, 0,
	0, 0, 0, 0, 0, 27, 0, 0, 110, 0,
	26, 109, 111, 0, 0, 167, 9, 0, 23, 21,
	0, 24, 13, 12, 22, 0, 0, 11, 10, 0,
	0, 0, 109, 111, 0, 110, 27, 9, 0, 23,
	21, 26, 24, 13, 12, 22, 105, 0, 11, 10,
	0, 0, 0, 0, 0, 0, 0, 27, 0, 0,
	7, 18, 26, 0, 0, 9, 110, 23, 21, 28,
	24, 13, 12, 22, 0, 0, 11, 10, 0, 0,
	0, 109, 111, 0, 0, 27, 9, 110, 23, 21,
	26, 24, 13, 12, 22, 0, 0, 11, 10, 0,
	0, 0, 0, 0, 0, 204, 27, 0, 0, 0,
	0, 26, 0, 0, 0, 8, 0, 0, 0, 0,
	0, 0, 0, 0, 68, 69, 74, 75, 73, 72,
	0, 0, 0, 197, 0, 0, 110, 70, 71, 76,
	78, 79, 80, 82, 83, 0, 0, 0, 81, 68,
	69, 74, 75, 73, 72, 77, 0, 0, 114, 0,
	0, 0, 70, 71, 76, 78, 79, 80, 82, 83,
	0, 0, 0, 81, 68, 69, 74, 75, 73, 72,
	77, 0, 0, 0, 0, 0, 0, 70, 71, 76,
	78, 79, 80, 82, 83, 103, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 68, 69, 74, 75, 73, 72,
	0, 0, 0, 0, 0, 0, 0, 70, 71, 76,
	78, 79, 80, 82, 83, 0, 0, 0, 81, 68,
	69, 74, 75, 73, 72, 77, 0, 0, 0, 0,
	0, 0, 70, 71, 76, 78, 79, 80, 82, 83,
	0, 0, 0, 81, 69, 74, 75, 73, 72, 0,
	77, 0, 0, 0, 0, 0, 70, 71, 76, 78,
	79, 80, 82, 83, 0, 0, 0, 81, 74, 75,
	73, 72, 0, 0, 77, 0, 0, 0, 0, 70,
	71, 76, 78, 79, 80, 82, 83, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 77,
}
var yyPact = [...]int{

	-1000, 776, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	107, 105, -1000, 574, -1000, -1000, 123, -1000, 73, 31,
	104, 574, -1000, 103, 574, -1000, 574, 76, -1000, 522,
	-1000, 78, -31, 936, -1000, -1000, 84, 574, 574, 574,
	-1000, -1000, 574, 101, -1000, -1000, -10, 1, -1000, -1000,
	574, 30, 23, 911, 748, 20, -12, 871, 334, 574,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 574, 574, 574,
	574, 574, 574, 574, 574, 574, 574, 574, 574, 574,
	574, 574, 574, 574, 574, 85, -1000, -1000, -1000, 261,
	76, -15, -20, -1000, 16, 574, 574, 68, -31, -1000,
	84, 36, -1, -1000, -1000, 574, -1000, -1000, -1000, -1000,
	-1000, 59, 574, 574, -1000, -1000, 936, 936, 960, 983,
	69, 69, 69, 69, 69, 69, 53, 172, 53, -1000,
	-1000, -1000, -1000, 53, 220, -1000, 76, 574, 574, -1000,
	-17, 132, -1000, 574, 448, -31, -1000, 727, -1000, -1000,
	-44, 696, 936, 669, 366, 2, 495, -1000, 574, 423,
	936, -1000, 15, 574, -48, 936, 13, -1000, 638, -1000,
	-1000, -1000, 574, -1000, 133, -1000, 574, 936, -1000, 574,
	398, -1000, -18, 574, -1000, 302, 611, -1000, 797, 846,
	936, -3, -1000, 936, -1000, 574, -1000, -1000, 574, 580,
	821, 495, 936, -1000, -1000, -1000, 553, -1000,
}
var yyPgo = [...]int{

	0, 152, 33, 181, 179, 35, 178, 2, 7, 175,
	5, 0, 174, 8, 1, 120, 172, 169, 4, 101,
	89, 167, 146, 140,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 4, 4, 4, 4, 4, 20, 20, 13, 13,
	13, 13, 13, 15, 15, 15, 15, 15, 12, 12,
	12, 16, 16, 16, 16, 16, 17, 18, 18, 18,
	21, 21, 22, 22, 23, 23, 19, 19, 19, 19,
	19, 5, 5, 5, 5, 6, 6, 7, 7, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 14, 14, 14, 14, 14, 14, 14, 14, 8,
	8, 9, 9, 10, 10, 10, 10,
}
var yyR2 = [...]int{

	0, 0, 2, 0, 2, 1, 1, 1, 1, 3,
	1, 1, 1, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 4, 3, 3, 6,
	5, 5, 4, 9, 7, 11, 6, 0, 2, 5,
	1, 2, 5, 6, 2, 3, 1, 2, 3, 1,
	2, 1, 4, 6, 3, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 1, 2, 2, 3, 4, 5, 6, 7, 1,
	3, 2, 3, 3, 5, 5, 7,
}
var yyChk = [...]int{

	-1000, -1, -3, -19, -20, -15, -22, 4, 59, 9,
	21, 20, 16, 15, -16, -17, -14, -12, 5, -6,
	-21, 12, 17, 11, 14, -5, 34, 29, 13, -1,
	29, 29, -8, -11, 30, 31, -14, 19, 39, 44,
	-5, 29, 33, 45, 31, -9, 34, -7, 13, 29,
	35, 61, 29, -11, -2, 29, -7, -11, -11, -13,
	48, 49, 50, 51, 52, 8, 20, 61, 23, 24,
	36, 37, 28, 27, 25, 26, 38, 54, 39, 40,
	41, 47, 42, 43, 33, 45, -11, -11, -11, -11,
	29, -10, -8, 62, 29, 33, 35, 61, -8, -5,
	-14, -23, 34, 4, -4, 18, -19, -20, -15, 4,
	59, 5, 35, 22, 7, 62, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, 29, 60, 63, -13, 62,
	61, 61, 62, 35, -11, -8, 29, -2, 31, 62,
	-7, -2, -11, -2, -11, -8, -2, 60, -13, -11,
	-11, 62, 29, 33, -10, -11, 60, 8, -2, 62,
	8, 8, 61, 4, -18, 10, 6, -11, 60, 35,
	-11, 62, 61, 35, 8, -11, -2, 8, -2, -11,
	-11, 60, 62, -11, 4, 61, 8, 7, 35, -2,
	-11, -2, -11, 8, 4, -18, -2, 8,
}
var yyDef = [...]int{

	1, -2, 2, 5, 6, 7, 8, 1, 10, 46,
	0, 0, 49, 0, 16, 17, 23, 24, 0, 0,
	0, 0, 3, 0, 0, -2, 0, 51, 40, 0,
	47, 0, 50, 89, 59, 60, 61, 0, 0, 0,
	81, 51, 0, 0, 82, 83, 0, 25, 41, 57,
	0, 0, 0, 0, 0, 57, 0, 0, 0, 0,
	18, 19, 20, 21, 22, 9, 48, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 78, 79, 80, 0,
	54, 0, 0, 91, 51, 0, 0, 0, 27, -2,
	0, 3, 0, 3, 4, 0, 11, 12, 13, 3,
	15, 0, 0, 0, 3, 84, 28, 90, 62, 63,
	64, 65, 66, 67, 68, 69, 70, 71, 72, 73,
	74, 75, 76, 77, 0, 54, 52, 0, 0, 85,
	0, 0, 92, 0, 0, 26, 58, 0, 3, 44,
	0, 0, 32, 0, 0, 0, 37, 52, 0, 0,
	30, 86, 0, 0, 0, 93, 0, 42, 0, 45,
	31, 14, 0, 3, 0, 3, 0, 29, 53, 0,
	0, 87, 0, 0, 43, 0, 0, 36, 38, 0,
	95, 0, 88, 94, 3, 0, 34, 3, 0, 0,
	0, 37, 96, 33, 3, 39, 0, 35,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 44, 3, 42, 46, 3,
	34, 62, 40, 38, 61, 39, 45, 41, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 63, 59,
	37, 35, 36, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 55, 3, 3, 3, 3, 3,
	3, 33, 3, 60, 43, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 32,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	47, 48, 49, 50, 51, 52, 53, 54, 56, 57,
	58,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:58
		{
			yyVAL.expr = __chain()
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:64
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[2].expr)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:72
		{
			yyVAL.expr = __chain()
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:75
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[2].expr)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:80
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:81
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:82
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:83
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:84
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:85
		{
			yyVAL.expr = emptyNode
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:88
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:89
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:90
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:91
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:92
		{
			yyVAL.expr = emptyNode
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:95
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:96
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:99
		{
			yyVAL.expr = NewSymbol(AAdd).SetPos(yyDollar[1].token.Pos)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:100
		{
			yyVAL.expr = NewSymbol(ASub).SetPos(yyDollar[1].token.Pos)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:101
		{
			yyVAL.expr = NewSymbol(AMul).SetPos(yyDollar[1].token.Pos)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:102
		{
			yyVAL.expr = NewSymbol(ADiv).SetPos(yyDollar[1].token.Pos)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:103
		{
			yyVAL.expr = NewSymbol(AMod).SetPos(yyDollar[1].token.Pos)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:106
		{
			if yyDollar[1].expr.isCallStat() {
				// Single call statement, clear env.V to avoid side effects
				yyVAL.expr = __chain(yyDollar[1].expr, popvClearNode)
			} else {
				yyVAL.expr = yyDollar[1].expr
			}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:114
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:117
		{
			yyVAL.expr = __chain()
			for _, v := range yyDollar[2].expr.Nodes {
				yyVAL.expr = yyVAL.expr.append(__set(v, NewSymbol(ANil)).SetPos(yyDollar[1].token.Pos))
			}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:123
		{
			yyVAL.expr = __local(yyDollar[2].expr.Nodes, yyDollar[4].expr.Nodes, yyDollar[1].token.Pos)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:126
		{
			yyVAL.expr = __moveMulti(yyDollar[1].expr.Nodes, yyDollar[3].expr.Nodes, yyDollar[2].token.Pos)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:131
		{
			if yyDollar[2].expr.SymbolValue() == AAdd && yyDollar[3].expr.IsNumber() {
				yyVAL.expr = __inc(NewSymbolFromToken(yyDollar[1].token), yyDollar[3].expr).SetPos(yyDollar[2].expr.Pos())
			} else if yyDollar[2].expr.SymbolValue() == ASub && yyDollar[3].expr.IsNumber() {
				f, i, isInt := yyDollar[3].expr.NumberValue()
				if isInt {
					yyVAL.expr = __inc(NewSymbolFromToken(yyDollar[1].token), NewNumberFromInt(-i)).SetPos(yyDollar[2].expr.Pos())
				} else {
					yyVAL.expr = __inc(NewSymbolFromToken(yyDollar[1].token), NewNumberFromFloat(-f)).SetPos(yyDollar[2].expr.Pos())
				}
			} else {
				yyVAL.expr = __move(NewSymbolFromToken(yyDollar[1].token), NewComplex(yyDollar[2].expr, NewSymbolFromToken(yyDollar[1].token), yyDollar[3].expr)).SetPos(yyDollar[2].expr.Pos())
			}
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:145
		{
			yyVAL.expr = __store(yyDollar[1].expr, yyDollar[3].expr, NewComplex(yyDollar[5].expr, __load(yyDollar[1].expr, yyDollar[3].expr), yyDollar[6].expr).SetPos(yyDollar[5].expr.Pos()))
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:148
		{
			i := NewString(yyDollar[3].token.Str)
			yyVAL.expr = __store(yyDollar[1].expr, i, NewComplex(yyDollar[4].expr, __load(yyDollar[1].expr, i), yyDollar[5].expr).SetPos(yyDollar[4].expr.Pos()))
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:154
		{
			yyVAL.expr = __loop(__if(yyDollar[2].expr, yyDollar[4].expr, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos)
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:157
		{
			yyVAL.expr = __loop(
				__chain(
					yyDollar[2].expr,
					__if(yyDollar[4].expr, breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			).SetPos(yyDollar[1].token.Pos)
		}
	case 33:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:165
		{
			forVar, forEnd := NewSymbolFromToken(yyDollar[2].token), randomVarname()
			yyVAL.expr = __do(
				__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
				__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
				__loop(
					__if(
						__lessEq(forVar, forEnd),
						__chain(yyDollar[8].expr, __inc(forVar, oneNode).SetPos(yyDollar[1].token.Pos)),
						breakNode,
					).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			)
		}
	case 34:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:179
		{
			f := randomVarname()
			bsDDD, bs := randomDDDVarname()
			if len(yyDollar[4].expr.Nodes) == 1 {
				// for b1, ..., bn in expr do
				vsDDD, vs := randomDDDVarname()
				yyVAL.expr = __do(
					// local f, ...vs = expr
					__local([]Node{f, vsDDD}, yyDollar[4].expr.Nodes, yyDollar[1].token.Pos),
					// local b1, ..., bn, ...bs = f(vs)
					__local(append(yyDollar[2].expr.Nodes, bsDDD), []Node{__call(f, NewComplex(vs)).SetPos(yyDollar[1].token.Pos)}, yyDollar[1].token.Pos))
			} else {
				// for b1, ..., bn in expr, v1, ..., vn do
				yyVAL.expr = __do(
					// local f = expr
					__set(f, yyDollar[4].expr.Nodes[0]).SetPos(yyDollar[1].token.Pos),
					// local b1, ..., bn, ...bs = f(v1, ..., vn)
					__local(append(yyDollar[2].expr.Nodes, bsDDD), []Node{__call(f, NewComplex(yyDollar[4].expr.Nodes[1:]...)).SetPos(yyDollar[1].token.Pos)}, yyDollar[1].token.Pos))
			}
			yyVAL.expr = yyVAL.expr.append(
				__loop(
					__chain(
						__if(
							NewComplex(NewSymbol(AEq), NewSymbol("nil"), yyDollar[2].expr.Nodes[0]).SetPos(yyDollar[1].token.Pos),
							breakNode,
							__chain(yyDollar[6].expr),
						).SetPos(yyDollar[1].token.Pos),
						// b1, ..., bn, ...bs = f(b1, ..., bn, bs)
						__moveMulti(append(yyDollar[2].expr.Nodes, bsDDD), []Node{
							__call(f, NewComplex(append(yyDollar[2].expr.DuplicateNodes(), bs)...)).SetPos(yyDollar[1].token.Pos),
						}, yyDollar[1].token.Pos),
					),
				).SetPos(yyDollar[1].token.Pos),
			)
		}
	case 35:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:214
		{
			forVar, forEnd := NewSymbolFromToken(yyDollar[2].token), randomVarname()
			if yyDollar[8].expr.IsNumber() { // step is a static number, easy case
				var cond Node
				if yyDollar[8].expr.IsNegativeNumber() {
					cond = __lessEq(forEnd, forVar)
				} else {
					cond = __lessEq(forVar, forEnd)
				}
				yyVAL.expr = __do(
					__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
					__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
					__loop(
						__chain(
							__if(
								cond,
								__chain(yyDollar[10].expr, __inc(forVar, yyDollar[8].expr)),
								breakNode,
							).SetPos(yyDollar[1].token.Pos),
						),
					).SetPos(yyDollar[1].token.Pos),
				)
			} else {
				forStep := randomVarname()
				yyVAL.expr = __do(
					__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
					__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
					__set(forStep, yyDollar[8].expr).SetPos(yyDollar[1].token.Pos),
					__loop(
						__chain(
							__if(
								__less(zeroNode, forStep).SetPos(yyDollar[1].token.Pos),
								// +step
								__if(__less(forEnd, forVar), breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
								// -step
								__if(__less(forVar, forEnd), breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
							).SetPos(yyDollar[1].token.Pos),
							yyDollar[10].expr,
							__inc(forVar, forStep),
						),
					).SetPos(yyDollar[1].token.Pos),
				)
			}

		}
	case 36:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:261
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:266
		{
			yyVAL.expr = NewComplex()
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:269
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:272
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:277
		{
			yyVAL.expr = NewSymbol(AMove).SetPos(yyDollar[1].token.Pos)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:278
		{
			yyVAL.expr = NewSymbol(ASet).SetPos(yyDollar[1].token.Pos)
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:281
		{
			yyVAL.expr = __func(yyDollar[1].expr, yyDollar[2].token, yyDollar[3].expr, "", yyDollar[4].expr)
		}
	case 43:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:282
		{
			yyVAL.expr = __func(yyDollar[1].expr, yyDollar[2].token, yyDollar[3].expr, yyDollar[4].token.Str, yyDollar[5].expr)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:285
		{
			yyVAL.expr = NewComplex()
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:286
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:289
		{
			yyVAL.expr = NewComplex(NewSymbol(ABreak)).SetPos(yyDollar[1].token.Pos)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:290
		{
			yyVAL.expr = NewComplex(NewSymbol(AGoto), NewSymbolFromToken(yyDollar[2].token)).SetPos(yyDollar[1].token.Pos)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:291
		{
			yyVAL.expr = NewComplex(NewSymbol(ALabel), NewSymbolFromToken(yyDollar[2].token))
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:292
		{
			yyVAL.expr = NewComplex(NewSymbol(AReturn), emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:293
		{
			if len(yyDollar[2].expr.Nodes) == 1 {
				x := yyDollar[2].expr.Nodes[0]
				if len(x.Nodes) == 3 && x.Nodes[0].SymbolValue() == ACall {
					x.Nodes[0].strSym = ATailCall
				}
			}
			yyVAL.expr = NewComplex(NewSymbol(AReturn), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:304
		{
			yyVAL.expr = NewSymbolFromToken(yyDollar[1].token)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:305
		{
			yyVAL.expr = __load(yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:306
		{
			yyVAL.expr = NewComplex(NewSymbol(ASlice), yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:307
		{
			yyVAL.expr = __load(yyDollar[1].expr, NewString(yyDollar[3].token.Str)).SetPos(yyDollar[2].token.Pos)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:310
		{
			yyVAL.expr = NewComplex(yyDollar[1].expr)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:311
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[3].expr)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:314
		{
			yyVAL.expr = NewComplex(NewSymbolFromToken(yyDollar[1].token))
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:315
		{
			yyVAL.expr = yyDollar[1].expr.append(NewSymbolFromToken(yyDollar[3].token))
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:318
		{
			yyVAL.expr = NewNumberFromString(yyDollar[1].token.Str)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:319
		{
			yyVAL.expr = NewString(yyDollar[1].token.Str)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:320
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:321
		{
			yyVAL.expr = NewComplex(NewSymbol(AOr), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:322
		{
			yyVAL.expr = NewComplex(NewSymbol(AAnd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:323
		{
			yyVAL.expr = NewComplex(NewSymbol(ALess), yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:324
		{
			yyVAL.expr = NewComplex(NewSymbol(ALess), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:325
		{
			yyVAL.expr = NewComplex(NewSymbol(ALessEq), yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:326
		{
			yyVAL.expr = NewComplex(NewSymbol(ALessEq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:327
		{
			yyVAL.expr = NewComplex(NewSymbol(AEq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:328
		{
			yyVAL.expr = NewComplex(NewSymbol(ANeq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:329
		{
			yyVAL.expr = NewComplex(NewSymbol(AAdd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:330
		{
			yyVAL.expr = NewComplex(NewSymbol(AConcat), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:331
		{
			yyVAL.expr = NewComplex(NewSymbol(ASub), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:332
		{
			yyVAL.expr = NewComplex(NewSymbol(AMul), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:333
		{
			yyVAL.expr = NewComplex(NewSymbol(ADiv), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:334
		{
			yyVAL.expr = NewComplex(NewSymbol(AIDiv), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:335
		{
			yyVAL.expr = NewComplex(NewSymbol(AMod), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:336
		{
			yyVAL.expr = NewComplex(NewSymbol(APow), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:337
		{
			yyVAL.expr = NewComplex(NewSymbol(ANot), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:338
		{
			yyVAL.expr = NewComplex(NewSymbol(ASub), zeroNode, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:339
		{
			yyVAL.expr = NewComplex(NewSymbol(ALen), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:342
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:343
		{
			yyVAL.expr = __call(yyDollar[1].expr, NewComplex(NewString(yyDollar[2].token.Str))).SetPos(yyDollar[1].expr.Pos())
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:344
		{
			yyVAL.expr = __call(yyDollar[1].expr, yyDollar[2].expr).SetPos(yyDollar[1].expr.Pos())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:345
		{
			yyVAL.expr = __chain(yyDollar[2].expr)
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:346
		{
			yyVAL.expr = __callMap(yyDollar[1].expr, emptyNode, yyDollar[3].expr).SetPos(yyDollar[1].expr.Pos())
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:347
		{
			yyVAL.expr = __callMap(yyDollar[1].expr, emptyNode, yyDollar[3].expr).SetPos(yyDollar[1].expr.Pos())
		}
	case 87:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:348
		{
			yyVAL.expr = __callMap(yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr).SetPos(yyDollar[1].expr.Pos())
		}
	case 88:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:349
		{
			yyVAL.expr = __callMap(yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr).SetPos(yyDollar[1].expr.Pos())
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:352
		{
			yyVAL.expr = NewComplex(yyDollar[1].expr)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:353
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[3].expr)
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:356
		{
			yyVAL.expr = NewComplex()
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:357
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:360
		{
			yyVAL.expr = NewComplex(NewString(yyDollar[1].token.Str), yyDollar[3].expr)
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:361
		{
			yyVAL.expr = NewComplex(yyDollar[2].expr, yyDollar[5].expr)
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:362
		{
			yyVAL.expr = yyDollar[1].expr.append(NewString(yyDollar[3].token.Str)).append(yyDollar[5].expr)
		}
	case 96:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:363
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[4].expr).append(yyDollar[7].expr)
		}
	}
	goto yystack /* stack new state and value */
}
