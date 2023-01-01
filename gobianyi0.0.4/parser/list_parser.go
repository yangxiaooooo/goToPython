package simple_parser

import (
	"errors"
	"fmt"
	"gobianyi/lexer"
)

type SimpleParser struct {
	lexer lexer.Lexer
	top   *Env //当前作用域符号表
	saved *Env //进入下一个作用域前，保存当前作用域，方便链表连接

}

func NewSimpleParser(lexer lexer.Lexer) *SimpleParser {
	return &SimpleParser{
		lexer: lexer,
		top:   nil,
		saved: nil,
	}
}

func (s *SimpleParser) Parse() error {
	return s.program()
}

func (s *SimpleParser) program() error {
	s.top = nil
	return s.block()
}

func (s *SimpleParser) match(str string) error {
	if s.lexer.Lexme != str {
		err_s := fmt.Sprintf("match error ,expect %s got %s", str, s.lexer.Lexme)
		return errors.New(err_s)
	}

	return nil
}

func (s *SimpleParser) block() error {
	s.lexer.Scan()
	err := s.match("{")
	if err != nil {
		return err
	}

	s.saved = s.top
	s.top = NewEnv(s.top)
	fmt.Printf("{")

	err = s.decls()
	if err != nil {
		return err
	}

	err = s.stmts()
	if err != nil {
		return err
	}

	err = s.match("}")
	if err != nil {
		return err
	}

	fmt.Printf("}")
	return err
}

func (s *SimpleParser) decls() error {
	return s.decls_r()
}

func (s *SimpleParser) decls_r() error {
	var err error
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag == lexer.TYPE {
		s.lexer.ReverseScan() //?
		s.decl()
		if err != nil {
			return err
		}
		return s.decls_r()
	} else {
		s.lexer.ReverseScan()
	}

	return nil
}

func (s *SimpleParser) decl() error {
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag != lexer.TYPE {
		str := fmt.Sprintf("in decl, expect type but got :%s", s.lexer.Lexme)
		return errors.New(str)
	}

	type_str := s.lexer.Lexme

	tag, err = s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag != lexer.ID {
		str := fmt.Sprintf("in decl, expect id got %s", s.lexer.Lexme)
		return errors.New(str)
	}

	id_str := s.lexer.Lexme
	symbol := NewSymbol(id_str, type_str)
	s.top.Put(id_str, symbol)

	_, err = s.lexer.Scan()
	if err != nil {
		return err
	}

	err = s.match(";")

	return err
}

func (s *SimpleParser) stmts() error {
	return s.stmts_r()
}

func (s *SimpleParser) stmts_r() error {
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}

	if tag.Tag == lexer.ID || tag.Tag == lexer.LEFT_BRACE {
		s.lexer.ReverseScan()
		err = s.stmt()
		if err != nil {
			return err
		}
		err = s.stmts_r()
	} else if tag.Tag == lexer.SEMICOLON {
		return nil
	}
	return nil
}

func (s *SimpleParser) stmt() error {
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	} else if tag.Tag == lexer.ID {
		s.lexer.ReverseScan()
		err = s.factor()
		if err != nil {
			return err
		}
		s.lexer.Scan()
		err = s.match(";")
		if err == nil {
			fmt.Print(";")
		}
	} else if tag.Tag == lexer.LEFT_BRACE {
		s.lexer.ReverseScan()
		err = s.block()

	} else {
		err = errors.New("stmt parsing error")
	}
	return err
}

func (s *SimpleParser) factor() error {
	tag, err := s.lexer.Scan()
	if err != nil {
		return err
	}
	if tag.Tag != lexer.ID {
		str := fmt.Sprintf("except identifier , got %s", s.lexer.Lexme)
		return errors.New(str)
	}

	symbol := s.top.Get(s.lexer.Lexme)

	fmt.Printf(s.lexer.Lexme)
	fmt.Printf(":")
	fmt.Printf(symbol.Type)

	return err
}
