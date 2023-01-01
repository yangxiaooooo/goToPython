package main

import (
	"gobianyi/lexer"
	"gobianyi/parser"
)

func main() {
	source := "9-5+2"
	my_lexer := lexer.NewLexer(source)
	parser := simple_parser.NewSimpleParser(my_lexer)
	root, _ := parser.Parse()
	print(root.Attribute())
}
