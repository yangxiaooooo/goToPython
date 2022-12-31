package main

import (
	"fmt"
	"gobianyi/lexer"
)

func main() {
	source := "if a1 >= 100.34"
	my_lexer := lexer.NewLexer(source)
	for {
		token, err := my_lexer.Scan()
		if err != nil {
			fmt.Println("lexer error", err)
			break
		}

		if token.Tag == lexer.EOF {
			break
		} else {
			fmt.Println("read token: ", token)
		}
	}
}
