package parser

type ExpressionType string

const (
	LITERAL ExpressionType = "literal"
	BINARY  ExpressionType = "binary"
	UNARY   ExpressionType = "unary"
	SELECT  ExpressionType = "select"
	FROM    ExpressionType = "from"
	WHERE   ExpressionType = "where"
)

type Expression struct {
}
