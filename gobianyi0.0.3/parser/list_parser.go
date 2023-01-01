package simple_parser

import (
	"errors"
	"gobianyi/lexer"
	"io"
)

type SimpleParser struct {
	lexer lexer.Lexer
}

func NewSimpleParser(lexer lexer.Lexer) *SimpleParser {
	return &SimpleParser{
		lexer: lexer,
	}
}

func (s *SimpleParser) list() (*SyntaxNode, error) {
	token, err := s.lexer.Scan()
	if err != nil {
		return nil, err
	}

	current_list_node := NewSyntaxNode()

	if token.Tag == lexer.LEFT_BRACKET {
		print("(")
		child_list_node, err := s.list()
		if err != nil {
			return nil, err
		}
		if child_list_node != nil {
			current_list_node.AddChild(child_list_node)
		}
		token, err = s.lexer.Scan()
		if token.Tag != lexer.RIGHT_BRACKET {
			err := errors.New("Missing of right bracket")
			return nil, err
		} else {
			print(")")
		}
	}

	if token.Tag == lexer.NUM {
		child_list_node := NewSyntaxNode()
		child_num_node, err := s.number()
		if err != nil {
			return nil, err
		}
		if child_num_node != nil {
			child_list_node.AddChild(child_num_node)
			current_list_node.AddChild(child_list_node)
		}
		print("NUM")
	}
	token, err = s.lexer.Scan()
	if err != nil {
		if err == io.EOF {
			return current_list_node, nil
		}
	}
	if token.Tag == lexer.PLUS || token.Tag == lexer.MINUS {
		current_list_node.T = s.lexer.Lexme
		print("+/-")
		child_list_node, err := s.list()
		if child_list_node != nil {
			current_list_node.AddChild(child_list_node)
		}
		if err != nil {
			if err == io.EOF {
				return current_list_node, nil
			}
			return nil, err
		}
	} else {
		s.lexer.ReverseScan()
	}
	return current_list_node, nil
}

func (s *SimpleParser) number() (*SyntaxNode, error) {
	current_list_node := NewSyntaxNode()
	current_list_node.T = s.lexer.Lexme
	return current_list_node, nil
}

func (s *SimpleParser) Parse() (*SyntaxNode, error) {
	return s.list()
}
