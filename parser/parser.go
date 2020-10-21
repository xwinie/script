// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2

//line parser.go.y:26
type yySymType struct {
	yys   int
	token Token
	expr  Node
	atom  Symbol
}

const TDo = 57346
const TIn = 57347
const TLocal = 57348
const TElseIf = 57349
const TThen = 57350
const TEnd = 57351
const TBreak = 57352
const TElse = 57353
const TFor = 57354
const TWhile = 57355
const TFunc = 57356
const TIf = 57357
const TLen = 57358
const TReturn = 57359
const TReturnVoid = 57360
const TImport = 57361
const TYield = 57362
const TYieldVoid = 57363
const TRepeat = 57364
const TUntil = 57365
const TNot = 57366
const TLabel = 57367
const TGoto = 57368
const TOr = 57369
const TAnd = 57370
const TEqeq = 57371
const TNeq = 57372
const TLte = 57373
const TGte = 57374
const TIdent = 57375
const TNumber = 57376
const TString = 57377
const TAddEq = 57378
const TSubEq = 57379
const TMulEq = 57380
const TDivEq = 57381
const TModEq = 57382
const TSquare = 57383
const TDotDot = 57384
const ASSIGN = 57385
const FUNC = 57386
const UNARY = 57387

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TDo",
	"TIn",
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
	"TLen",
	"TReturn",
	"TReturnVoid",
	"TImport",
	"TYield",
	"TYieldVoid",
	"TRepeat",
	"TUntil",
	"TNot",
	"TLabel",
	"TGoto",
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
	"':'",
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
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:376

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	39, 50,
	64, 50,
	-2, 77,
	-1, 107,
	39, 51,
	64, 51,
	-2, 77,
}

const yyPrivate = 57344

const yyLast = 885

var yyAct = [...]int{

	34, 39, 68, 19, 111, 97, 157, 1, 105, 151,
	56, 75, 140, 58, 32, 33, 75, 87, 88, 89,
	112, 104, 59, 98, 61, 91, 53, 65, 92, 55,
	67, 50, 62, 44, 19, 28, 63, 155, 47, 46,
	93, 94, 95, 96, 29, 136, 105, 60, 142, 134,
	101, 57, 100, 84, 86, 87, 88, 89, 90, 66,
	64, 49, 108, 105, 19, 43, 28, 27, 85, 116,
	58, 103, 114, 48, 99, 106, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 57, 133, 107, 164, 28, 75, 27, 69,
	70, 71, 72, 73, 139, 135, 45, 37, 18, 137,
	4, 144, 3, 146, 31, 17, 16, 143, 5, 36,
	141, 147, 148, 20, 41, 54, 22, 145, 27, 2,
	0, 0, 0, 46, 35, 38, 0, 19, 29, 0,
	153, 152, 0, 40, 53, 19, 51, 55, 42, 19,
	19, 0, 0, 161, 0, 0, 0, 163, 52, 0,
	166, 0, 0, 162, 19, 102, 165, 19, 0, 28,
	172, 0, 0, 19, 0, 19, 171, 28, 173, 19,
	176, 28, 28, 177, 84, 86, 87, 88, 89, 90,
	45, 0, 0, 0, 0, 0, 28, 0, 31, 28,
	0, 27, 0, 36, 0, 28, 0, 28, 41, 27,
	0, 28, 0, 27, 27, 0, 0, 46, 35, 38,
	0, 0, 29, 0, 0, 0, 0, 40, 27, 0,
	0, 27, 42, 0, 0, 0, 0, 27, 6, 27,
	21, 159, 0, 27, 10, 158, 25, 23, 31, 26,
	0, 15, 14, 11, 8, 9, 24, 0, 0, 13,
	12, 0, 0, 0, 0, 0, 0, 30, 6, 0,
	21, 0, 29, 178, 10, 0, 25, 23, 31, 26,
	0, 15, 14, 11, 8, 9, 24, 0, 0, 13,
	12, 0, 0, 0, 0, 0, 7, 30, 6, 0,
	21, 0, 29, 174, 10, 0, 25, 23, 31, 26,
	168, 15, 14, 11, 8, 9, 24, 0, 0, 13,
	12, 0, 0, 0, 0, 0, 7, 30, 0, 0,
	0, 0, 29, 76, 77, 82, 83, 81, 80, 0,
	0, 0, 0, 0, 0, 0, 78, 79, 84, 86,
	87, 88, 89, 90, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 85, 0, 0, 6, 0, 21, 0,
	169, 167, 10, 0, 25, 23, 31, 26, 0, 15,
	14, 11, 8, 9, 24, 0, 0, 13, 12, 0,
	0, 0, 0, 0, 0, 30, 6, 0, 21, 0,
	29, 160, 10, 0, 25, 23, 31, 26, 0, 15,
	14, 11, 8, 9, 24, 0, 0, 13, 12, 0,
	0, 0, 0, 0, 7, 30, 6, 0, 21, 0,
	29, 154, 10, 0, 25, 23, 31, 26, 0, 15,
	14, 11, 8, 9, 24, 0, 0, 13, 12, 0,
	0, 0, 0, 0, 7, 30, 6, 0, 21, 0,
	29, 150, 10, 0, 25, 23, 31, 26, 0, 15,
	14, 11, 8, 9, 24, 0, 0, 13, 12, 0,
	0, 0, 0, 0, 7, 30, 0, 0, 0, 0,
	29, 76, 77, 82, 83, 81, 80, 0, 0, 0,
	0, 0, 0, 0, 78, 79, 84, 86, 87, 88,
	89, 90, 0, 0, 7, 0, 0, 6, 0, 21,
	0, 85, 0, 10, 0, 25, 23, 31, 26, 115,
	15, 14, 11, 8, 9, 24, 110, 0, 13, 12,
	0, 0, 0, 0, 0, 0, 30, 6, 0, 21,
	0, 29, 74, 10, 0, 25, 23, 31, 26, 0,
	15, 14, 11, 8, 9, 24, 0, 0, 13, 12,
	0, 0, 0, 0, 0, 7, 30, 6, 0, 21,
	0, 29, 0, 10, 0, 25, 23, 31, 26, 0,
	15, 14, 11, 8, 9, 24, 0, 0, 13, 12,
	0, 0, 0, 0, 0, 7, 30, 0, 0, 0,
	0, 29, 76, 77, 82, 83, 81, 80, 0, 0,
	0, 0, 0, 0, 0, 78, 79, 84, 86, 87,
	88, 89, 90, 0, 0, 7, 0, 0, 0, 0,
	0, 0, 85, 76, 77, 82, 83, 81, 80, 156,
	0, 0, 0, 0, 0, 0, 78, 79, 84, 86,
	87, 88, 89, 90, 0, 0, 0, 76, 77, 82,
	83, 81, 80, 85, 0, 0, 0, 0, 0, 149,
	78, 79, 84, 86, 87, 88, 89, 90, 175, 0,
	0, 0, 0, 0, 0, 0, 0, 85, 0, 0,
	0, 0, 0, 138, 0, 0, 0, 0, 0, 0,
	0, 76, 77, 82, 83, 81, 80, 0, 0, 0,
	0, 0, 0, 170, 78, 79, 84, 86, 87, 88,
	89, 90, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 85, 76, 77, 82, 83, 81, 80, 0, 0,
	0, 0, 0, 0, 113, 78, 79, 84, 86, 87,
	88, 89, 90, 0, 0, 0, 0, 0, 0, 0,
	0, 109, 85, 76, 77, 82, 83, 81, 80, 0,
	0, 0, 0, 0, 0, 0, 78, 79, 84, 86,
	87, 88, 89, 90, 76, 77, 82, 83, 81, 80,
	0, 0, 0, 85, 0, 0, 0, 78, 79, 84,
	86, 87, 88, 89, 90, 76, 77, 82, 83, 81,
	80, 0, 0, 0, 85, 0, 0, 0, 78, 79,
	84, 86, 87, 88, 89, 90, 77, 82, 83, 81,
	80, 0, 0, 0, 0, 85, 0, 0, 78, 79,
	84, 86, 87, 88, 89, 90, 82, 83, 81, 80,
	0, 0, 0, 0, 0, 85, 0, 78, 79, 84,
	86, 87, 88, 89, 90, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 85,
}
var yyPact = [...]int{

	-1000, 573, -1000, -1000, -1000, -1000, -1000, -1000, 184, -1000,
	-1000, 3, 40, 28, -1000, 184, -1000, -1000, -1000, 109,
	-1000, 37, -17, 184, -1000, 27, 184, 26, -1000, 184,
	48, -1000, 543, -48, 788, -1000, -10, -1000, -1000, -9,
	184, 184, 184, -15, -1000, 78, -1000, -1000, -1000, 49,
	-48, 184, 17, -1000, -1000, 100, -18, -1000, -1000, 184,
	6, 767, 513, -1, -19, 746, -15, 464, 184, -1000,
	-1000, -1000, -1000, -1000, -1000, 184, 184, 184, 184, 184,
	184, 184, 184, 184, 184, 184, 184, 184, 184, 184,
	184, -1000, 184, 16, -1000, -1000, -1000, -1000, -20, -1000,
	640, 48, -1000, -53, 184, 15, -48, -1000, -9, -1000,
	184, 184, 184, -1000, -1000, -1000, 788, 788, 808, 827,
	11, 11, 11, 11, 11, 11, -27, 142, -27, -1000,
	-1000, -1000, -27, 616, -1000, 452, -1000, -56, 48, 184,
	-1000, -48, -1000, 422, 788, 33, 585, 234, 392, -1000,
	-1000, -1000, 184, 788, -1000, -1000, 184, 86, -1000, 184,
	-1000, 788, 362, 306, -1000, 573, 715, -1000, -1000, 184,
	-1000, 294, 684, 234, -1000, -1000, -1000, 264, -1000,
}
var yyPgo = [...]int{

	0, 7, 129, 33, 126, 10, 15, 125, 0, 123,
	2, 1, 118, 116, 115, 6, 112, 110, 65, 108,
	107, 5,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 17, 17,
	17, 10, 10, 10, 10, 10, 12, 12, 12, 12,
	12, 9, 9, 9, 13, 13, 13, 13, 13, 14,
	15, 15, 15, 18, 18, 19, 20, 21, 21, 16,
	16, 16, 16, 16, 16, 16, 16, 3, 3, 3,
	4, 4, 5, 5, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 11, 11, 11,
	11, 6, 6, 7, 7,
}
var yyR2 = [...]int{

	0, 0, 2, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	3, 3, 6, 5, 5, 4, 7, 9, 11, 6,
	0, 2, 5, 1, 2, 5, 4, 2, 3, 2,
	1, 1, 2, 2, 3, 1, 2, 1, 4, 3,
	1, 3, 1, 3, 1, 2, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 1, 2, 2,
	3, 1, 3, 2, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, -16, -17, -12, 4, 62, 20, 21,
	10, 19, 26, 25, 18, 17, -13, -14, -19, -11,
	-9, 6, -4, 13, 22, 12, 15, -18, -3, 38,
	33, 14, -1, -6, -8, 34, 19, -20, 35, -11,
	43, 24, 48, -18, -3, 6, 33, 35, 33, 33,
	-6, 37, 49, 35, -7, 38, -5, 14, 33, 39,
	64, -8, -1, -5, 33, -8, 33, -8, -10, 51,
	52, 53, 54, 55, 9, 64, 27, 28, 40, 41,
	32, 31, 29, 30, 42, 57, 43, 44, 45, 46,
	47, 35, 37, 49, -8, -8, -8, -21, 38, 25,
	-8, 33, 65, -6, 39, 64, -6, -3, -11, 4,
	23, 5, 39, 8, -21, 65, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, 33, -1, 65, -5, 63, -10,
	65, -6, 33, -1, -8, -6, -8, -1, -1, 63,
	9, 65, -10, -8, 9, 4, 64, -15, 11, 7,
	9, -8, -1, -8, 9, -1, -8, 9, 4, 64,
	8, -1, -8, -1, 9, 4, -15, -1, 9,
}
var yyDef = [...]int{

	1, -2, 2, 3, 4, 5, 1, 7, 0, 40,
	41, 0, 0, 0, 45, 0, 8, 9, 10, 16,
	17, 0, 0, 0, 1, 0, 0, 0, -2, 0,
	47, 33, 0, 39, 81, 54, 0, 56, 57, 58,
	0, 0, 0, 0, 77, 0, 47, 42, 43, 0,
	46, 0, 0, 78, 79, 0, 18, 34, 52, 0,
	0, 0, 0, 0, 52, 0, 0, 0, 0, 11,
	12, 13, 14, 15, 6, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 0, 0, 74, 75, 76, 1, 0, 44,
	0, 49, 83, 0, 0, 0, 20, -2, 0, 1,
	0, 0, 0, 1, 1, 80, 21, 82, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 0, 49, 0, 37, 0, 48, 0,
	84, 19, 53, 0, 25, 0, 0, 30, 0, 48,
	36, 38, 0, 23, 24, 1, 0, 0, 1, 0,
	35, 22, 0, 0, 29, 31, 0, 26, 1, 0,
	1, 0, 0, 30, 27, 1, 32, 0, 28,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 48, 3, 46, 3, 3,
	38, 65, 44, 42, 64, 43, 49, 45, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 50, 62,
	41, 39, 40, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 58, 3, 3, 3, 3, 3,
	3, 37, 3, 63, 47, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 36,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 51, 52, 53, 54, 55, 56,
	57, 59, 60, 61,
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
			yyVAL.expr = yyDollar[1].expr.CplAppend(yyDollar[2].expr)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:72
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:73
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:74
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:75
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:76
		{
			yyVAL.expr = emptyNode
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:79
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:80
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:81
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:84
		{
			yyVAL.expr = Node{AAdd.SetPos(yyDollar[1].token.Pos)}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:85
		{
			yyVAL.expr = Node{ASub.SetPos(yyDollar[1].token.Pos)}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:86
		{
			yyVAL.expr = Node{AMul.SetPos(yyDollar[1].token.Pos)}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:87
		{
			yyVAL.expr = Node{ADiv.SetPos(yyDollar[1].token.Pos)}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:88
		{
			yyVAL.expr = Node{AMod.SetPos(yyDollar[1].token.Pos)}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:91
		{
			if yyDollar[1].expr.isCallStat() {
				// Single call statement, clear env.V to avoid side effects
				yyVAL.expr = __chain(yyDollar[1].expr, popvClearNode)
			} else {
				yyVAL.expr = yyDollar[1].expr
			}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:99
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:102
		{
			yyVAL.expr = __chain()
			for _, v := range yyDollar[2].expr.Cpl() {
				yyVAL.expr = yyVAL.expr.CplAppend(__set(v, Node{ANil}).SetPos(yyDollar[1].token.Pos))
			}
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:108
		{
			m, n := len(yyDollar[2].expr.Cpl()), len(yyDollar[4].expr.Cpl())
			for i, count := 0, m-n; i < count; i++ {
				if i == count-1 {
					yyDollar[4].expr = yyDollar[4].expr.CplAppend(popvEndNode)
				} else {
					yyDollar[4].expr = yyDollar[4].expr.CplAppend(popvNode)
				}
			}

			yyVAL.expr = __chain()
			for i, v := range yyDollar[2].expr.Cpl() {
				if v.SymDDD() {
					yyVAL.expr = yyVAL.expr.CplAppend(__set(v, __popvAll(i, yyDollar[4].expr.CplIndex(i))).SetPos(yyDollar[1].token.Pos))
				} else {
					yyVAL.expr = yyVAL.expr.CplAppend(__set(v, yyDollar[4].expr.CplIndex(i)).SetPos(yyDollar[1].token.Pos))
				}
			}

			if m == 1 && n == 1 && yyDollar[4].expr.CplIndex(0).isCallStat() {
				// Single call statement with single assignment, clear env.V to avoid side effects
				yyVAL.expr = yyVAL.expr.CplAppend(popvClearNode)
			}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:132
		{
			nodes := yyDollar[1].expr.Cpl()
			m, n := len(nodes), len(yyDollar[3].expr.Cpl())
			for i, count := 0, m-n; i < count; i++ {
				if i == count-1 {
					yyDollar[3].expr = yyDollar[3].expr.CplAppend(popvEndNode)
				} else {
					yyDollar[3].expr = yyDollar[3].expr.CplAppend(popvNode)
				}
			}

			if head := nodes[0]; len(nodes) == 1 && !nodes[0].SymDDD() {
				// a0 = b0
				// if a, s, ok := $3.CplIndex(0).isSimpleAddSub(); ok && a.Equals(head.Sym()) {
				//    $$ = __inc(head, Num(s)).SetPos($2.Pos)
				// } else {
				yyVAL.expr = head.moveLoadStore(__move, yyDollar[3].expr.CplIndex(0)).SetPos(yyDollar[2].token.Pos)
				// }
			} else {
				// a0, ..., an = b0, ..., bn
				yyVAL.expr = __chain()
				names, retaddr := []Node{}, Cpl(Node{ARetAddr})
				for i := range nodes {
					names = append(names, randomVarname())
					retaddr = retaddr.CplAppend(names[i])
					if nodes[i].SymDDD() {
						yyVAL.expr = yyVAL.expr.CplAppend(__set(names[i], __popvAll(i, yyDollar[3].expr.CplIndex(i))).SetPos(yyDollar[2].token.Pos))
					} else {
						yyVAL.expr = yyVAL.expr.CplAppend(__set(names[i], yyDollar[3].expr.CplIndex(i)).SetPos(yyDollar[2].token.Pos))
					}
				}
				for i, v := range nodes {
					yyVAL.expr = yyVAL.expr.CplAppend(v.moveLoadStore(__move, names[i]).SetPos(yyDollar[2].token.Pos))
				}
				yyVAL.expr = yyVAL.expr.CplAppend(retaddr)
			}

			if m == 1 && n == 1 && yyDollar[3].expr.CplIndex(0).isCallStat() {
				// Single call statement with single assignment, clear env.V to avoid side effects
				yyVAL.expr = __chain(yyVAL.expr, popvClearNode)
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:176
		{
			yyVAL.expr = __move(SymTok(yyDollar[1].token), Cpl(yyDollar[2].expr, SymTok(yyDollar[1].token), yyDollar[3].expr)).SetPos(yyDollar[2].expr.Pos())
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:179
		{
			yyVAL.expr = __store(yyDollar[1].expr, yyDollar[3].expr, Cpl(yyDollar[5].expr, __load(yyDollar[1].expr, yyDollar[3].expr), yyDollar[6].expr).SetPos(yyDollar[5].expr.Pos()))
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:182
		{
			i := Node{yyDollar[3].token.Str}
			yyVAL.expr = __store(yyDollar[1].expr, i, Cpl(yyDollar[4].expr, __load(yyDollar[1].expr, i), yyDollar[5].expr).SetPos(yyDollar[4].expr.Pos()))
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:188
		{
			yyVAL.expr = __loop(__if(yyDollar[2].expr, yyDollar[4].expr, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:191
		{
			yyVAL.expr = __loop(
				__chain(
					yyDollar[2].expr,
					__if(yyDollar[4].expr, breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			).SetPos(yyDollar[1].token.Pos)
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:199
		{
			yyVAL.expr = forLoop(yyDollar[1].token.Pos, yyDollar[2].expr.Cpl(), yyDollar[4].expr.Cpl(), yyDollar[6].expr)
		}
	case 27:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:202
		{
			forVar, forEnd := SymTok(yyDollar[2].token), randomVarname()
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
	case 28:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:216
		{
			forVar, forEnd := SymTok(yyDollar[2].token), randomVarname()
			if yyDollar[8].expr.Type() == NUM { // step is a static number, easy case
				var cond Node
				if f, i := yyDollar[8].expr.Num(); f < 0 || i < 0 {
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
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:263
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:268
		{
			yyVAL.expr = Cpl()
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:271
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:274
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:279
		{
			yyVAL.expr = Node{AMove}.SetPos(yyDollar[1].token.Pos)
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:280
		{
			yyVAL.expr = Node{ASet}.SetPos(yyDollar[1].token.Pos)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:283
		{
			funcname := SymTok(yyDollar[2].token)
			x := __move
			if yyDollar[1].expr.Sym().Equals(ASet) {
				x = __set
			}
			yyVAL.expr = __chain(
				x(funcname, Node{ANil}).SetPos(yyDollar[1].expr.Pos()),
				__move(funcname, __func(funcname, yyDollar[3].expr, yyDollar[4].expr).SetPos(yyDollar[1].expr.Pos())).SetPos(yyDollar[1].expr.Pos()),
			)
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:296
		{
			yyVAL.expr = __func(emptyNode, yyDollar[2].expr, yyDollar[3].expr).SetPos(yyDollar[1].expr.Pos()).SetPos(yyDollar[1].expr.Pos())
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:301
		{
			yyVAL.expr = Cpl()
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:302
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:305
		{
			yyVAL.expr = Cpl(Node{AYield}, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:306
		{
			yyVAL.expr = Cpl(Node{AYield}, emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:307
		{
			yyVAL.expr = Cpl(Node{ABreak}).SetPos(yyDollar[1].token.Pos)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:308
		{
			yyVAL.expr = __move(Sym(moduleNameFromPath(yyDollar[2].token.Str)), yylex.(*Lexer).loadFile(joinSourcePath(yyDollar[1].token.Pos.Source, yyDollar[2].token.Str))).SetPos(yyDollar[1].token.Pos)
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:309
		{
			yyVAL.expr = Cpl(Node{AGoto}, SymTok(yyDollar[2].token)).SetPos(yyDollar[1].token.Pos)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:310
		{
			yyVAL.expr = Cpl(Node{ALabel}, SymTok(yyDollar[2].token))
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:311
		{
			yyVAL.expr = Cpl(Node{AReturn}, emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:312
		{
			if len(yyDollar[2].expr.Cpl()) == 1 {
				x := yyDollar[2].expr.CplIndex(0)
				if len(x.Cpl()) == 3 && x.CplIndex(0).Sym().Equals(ACall) {
					tc := x.CplIndex(0).Sym()
					tc.Text = ATailCall.Text
					x.Value.([]Node)[0] = Node{tc}
				}
			}
			yyVAL.expr = Cpl(Node{AReturn}, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:325
		{
			yyVAL.expr = SymTok(yyDollar[1].token)
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:326
		{
			yyVAL.expr = __load(yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos) /* (10)[0] is valid if number has metamethod */
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:327
		{
			yyVAL.expr = __load(yyDollar[1].expr, Node{yyDollar[3].token.Str}).SetPos(yyDollar[2].token.Pos)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:330
		{
			yyVAL.expr = Cpl(yyDollar[1].expr)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:331
		{
			yyVAL.expr = yyDollar[1].expr.CplAppend(yyDollar[3].expr)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:334
		{
			yyVAL.expr = Cpl(SymTok(yyDollar[1].token))
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:335
		{
			yyVAL.expr = yyDollar[1].expr.CplAppend(SymTok(yyDollar[3].token))
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:338
		{
			yyVAL.expr = Num(yyDollar[1].token.Str)
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:339
		{
			yyVAL.expr = yylex.(*Lexer).loadFile(joinSourcePath(yyDollar[1].token.Pos.Source, yyDollar[2].token.Str))
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:340
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:341
		{
			yyVAL.expr = Node{yyDollar[1].token.Str}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:342
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:343
		{
			yyVAL.expr = Cpl(Node{AOr}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:344
		{
			yyVAL.expr = Cpl(Node{AAnd}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:345
		{
			yyVAL.expr = Cpl(Node{ALess}, yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:346
		{
			yyVAL.expr = Cpl(Node{ALess}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:347
		{
			yyVAL.expr = Cpl(Node{ALessEq}, yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:348
		{
			yyVAL.expr = Cpl(Node{ALessEq}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:349
		{
			yyVAL.expr = Cpl(Node{AEq}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:350
		{
			yyVAL.expr = Cpl(Node{ANeq}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:351
		{
			yyVAL.expr = Cpl(Node{AAdd}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:352
		{
			yyVAL.expr = Cpl(Node{AConcat}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:353
		{
			yyVAL.expr = Cpl(Node{ASub}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:354
		{
			yyVAL.expr = Cpl(Node{AMul}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:355
		{
			yyVAL.expr = Cpl(Node{ADiv}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:356
		{
			yyVAL.expr = Cpl(Node{AMod}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:357
		{
			yyVAL.expr = Cpl(Node{APow}, yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:358
		{
			yyVAL.expr = Cpl(Node{AUnm}, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:359
		{
			yyVAL.expr = Cpl(Node{ANot}, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:360
		{
			yyVAL.expr = Cpl(Node{ALen}, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:363
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:364
		{
			yyVAL.expr = __call(yyDollar[1].expr, Cpl(Node{yyDollar[2].token.Str})).SetPos(yyDollar[1].expr.Pos())
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:365
		{
			yyVAL.expr = __call(yyDollar[1].expr, yyDollar[2].expr).SetPos(yyDollar[1].expr.Pos())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:366
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:369
		{
			yyVAL.expr = Cpl(yyDollar[1].expr)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:370
		{
			yyVAL.expr = yyDollar[1].expr.CplAppend(yyDollar[3].expr)
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:373
		{
			yyVAL.expr = Cpl()
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:374
		{
			yyVAL.expr = yyDollar[2].expr
		}
	}
	goto yystack /* stack new state and value */
}
