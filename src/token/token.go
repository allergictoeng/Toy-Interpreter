package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Control
	ILLEGAL = "ILLEGAL" // When we find token/character we don't know
	EOF     = "EOF"     // End of file
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
