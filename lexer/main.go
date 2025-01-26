package lexer

func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

func isLetter(c string) bool {
	return c >= "a" && c <= "z" || c >= "A" && c <= "Z" || c == "_"
}

func Lex(code string) (tokens []Token) {
	var cmp = 0
	var line = 1
	var column = 1

	var number = func(start int) {
		for {
			cmp++
			column++

			if cmp >= len(code) || !isDigit(string(code[cmp])) {
				break
			}
		}

		if cmp < len(code) && string(code[cmp]) == "." && isDigit(string(code[cmp+1])) {
			cmp++
			column++

			for {
				cmp++
				column++

				if cmp >= len(code) || !isDigit(string(code[cmp])) {
					break
				}
			}
		}

		tokens = append(tokens, Token{
			Type:   NUMBER,
			Value:  code[start:cmp],
			Line:   line,
			Column: column,
		})
	}

	var letter = func(start int) {
		for {
			cmp++
			column++

			if cmp >= len(code) || (!isLetter(string(code[cmp])) && !isDigit(string(code[cmp]))) {
				break
			}
		}

		identifier := code[start : cmp+1]

		startIndex := 0
		endIndex := len(identifier)
		if identifier[0:1] == "'" {
			tokens = append(tokens, Token{
				Type:   QUOTE,
				Value:  "'",
				Line:   line,
				Column: column,
			})
			startIndex = 1
		}

		var quoteAfter = false
		var afterVirgule = false
		if identifier[len(identifier)-1:] == "," {
			endIndex = len(identifier) - 1
			afterVirgule = true

			if identifier[len(identifier)-2:] == "'" {
				endIndex = len(identifier) - 2
				quoteAfter = true
			}
		} else if identifier[len(identifier)-1:] == "'" {
			endIndex = len(identifier) - 1
			quoteAfter = true
		}

		tokens = append(tokens, Token{
			Type:   IDENTIFIER,
			Value:  identifier[startIndex:endIndex],
			Line:   line,
			Column: column,
		})

		if quoteAfter {
			tokens = append(tokens, Token{
				Type:   QUOTE,
				Value:  "'",
				Line:   line,
				Column: column,
			})
		}

		if afterVirgule {
			tokens = append(tokens, Token{
				Type:   VIRGULE,
				Value:  ",",
				Line:   line,
				Column: column,
			})
		}
	}

	for {
		if code[cmp] == '\n' {
			line++
			column = 1
		} else {
			column++
		}

		var c = string(code[cmp])
		//println(c)

		//var re = regexp.MustCompile(`(?m)'`)
		//for _, match := range re.FindAllString(c, -1) {
		//	fmt.Println(match + " test")
		//}
		switch c {
		case " ":
		case "\n":
		case "\r":
		case "\t":
			break
		case ",":
			tokens = append(tokens, Token{
				Type:   VIRGULE,
				Value:  c,
				Line:   line,
				Column: column,
			})
			break
		case "(":
		case ")":
			tokens = append(tokens, Token{
				Type:   PARENTHESIS,
				Value:  c,
				Line:   line,
				Column: column,
			})
			break
		case "=":
			tokens = append(tokens, Token{
				Type:   EQUAL,
				Value:  c,
				Line:   line,
				Column: column,
			})
			break
		/*case "+":
			tokens = append(tokens, Token{
				Type:   ADDITION,
				Value:  c,
				Line:   line,
				Column: column,
			})
			break
		case "-":
			tokens = append(tokens, Token{
				Type:   SUBTRACTION,
				Value:  c,
				Line:   line,
				Column: column,
			})
		case "*":
			tokens = append(tokens, Token{
				Type:   MULTIPLY,
				Value:  c,
				Line:   line,
				Column: column,
			})
			break
		case "/":
			tokens = append(tokens, Token{
				Type:   DIVIDE,
				Value:  c,
				Line:   line,
				Column: column,
			})
			break*/
		default:
			if isDigit(c) {
				number(cmp)
			} else if isLetter(c) || c == "'" || c == "\"" {
				letter(cmp)
			} else {
				panic("unexpected token " + c)
			}
		}

		cmp++

		if cmp >= len(code) {
			break
		}
	}
	return
}
