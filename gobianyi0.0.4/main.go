package main

import (
	"fmt"
	"gobianyi/lexer"
	"gobianyi/parser"
	"io"
)

func main() {
	source := "{int x; char y; {bool y; x; y;} x; y;}"
	my_lexer := lexer.NewLexer(source)
	parser := simple_parser.NewSimpleParser(my_lexer)
	err := parser.Parse()
	if err == io.EOF || err == nil {
		fmt.Sprintf("\nparsing success")
	} else {
		print(err.Error())
	}
}
