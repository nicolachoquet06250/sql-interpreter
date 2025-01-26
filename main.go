package main

import (
	"fmt"
	"sql/lexer"
)

func main() {
	var lex = lexer.Lex(`SELECT 'id_user' as id, 'stripe_id' as stripe FROM user_daft WHERE id = 1`)
	fmt.Printf("%v, %d", lex, len(lex))
}
