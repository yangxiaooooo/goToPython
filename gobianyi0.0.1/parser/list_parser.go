package simple_parser

import (
	"errors"
	"gobianyi/lexer"
)

type SimpleParser struct {
	lexer lexer.Lexer
}

func NewSimpleParser(lexer lexer.Lexer) *SimpleParser {
	return &SimpleParser{
		lexer: lexer,
	}
}

func (s *SimpleParser) list() error {
	token, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if token.Tag == lexer.LEFT_BRACKET {
		print("(")
		s.list()
		token, err = s.lexer.Scan()
		if token.Tag != lexer.RIGHT_BRACKET {
			err := errors.New("Missing of right bracket")
			return err
		} else {
			print(")")
		}
	}

	if token.Tag == lexer.NUM {
		print("NUM")
	}
	token, err = s.lexer.Scan()
	if token.Tag == lexer.PLUS || token.Tag == lexer.MINUS {
		print("+/-")
		s.list()
	} else {
		s.lexer.ReverseScan()
	}
	return err
}

func (s *SimpleParser) Parse() error {
	return s.list()
}
