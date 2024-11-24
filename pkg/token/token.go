package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 特殊标记

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 标识符 + 字面量

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 12345

	// 操作符

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// 分隔符

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent 根据标识符返回对应的关键字标识
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
