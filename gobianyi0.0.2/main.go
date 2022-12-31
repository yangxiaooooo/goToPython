package main

import (
	"gobianyi/lexer"
	"gobianyi/parser"
)

func main() {
	source := "(1+(2+3)"
	my_lexer := lexer.NewLexer(source)
	parser := simple_parser.NewSimpleParser(my_lexer)
	parser.Parse()
}
