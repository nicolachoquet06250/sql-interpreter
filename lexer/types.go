package lexer

type TokenType string

const (
	ADDITION    TokenType = "addition"
	SUBTRACTION TokenType = "subtraction"
	MULTIPLY    TokenType = "multiply"
	DIVIDE      TokenType = "divide"

	PARENTHESIS TokenType = "parenthesis"
	EQUAL       TokenType = "equal"
	QUOTE       TokenType = "quote"
	DB_QUOTE    TokenType = "db_quote"
	VIRGULE     TokenType = "virgule"

	NUMBER     TokenType = "number"
	IDENTIFIER TokenType = "identifier"
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}
