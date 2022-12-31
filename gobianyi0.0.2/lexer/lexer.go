package lexer

import (
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	lexme     string
	peek      byte             //读入的字符
	line      int              //当前字符串处于第几行
	reader    *bufio.Reader    //用于读取字节流
	key_words map[string]Token //存储关键字
}

func NewLexer(source string) Lexer {
	str := strings.NewReader(source)
	source_reader := bufio.NewReaderSize(str, len(source))
	lexer := Lexer{
		line:      1,
		reader:    source_reader,
		key_words: make(map[string]Token),
	}
	lexer.reserve()
	return lexer
}

func (l *Lexer) ReverseScan() {
	back_len := len(l.lexme)
	for i := 0; i < back_len; i++ {
		l.reader.UnreadByte()
	}
}

func (l *Lexer) reserve() {
	key_words := GetKeyWords()
	for _, key_word := range key_words {
		l.key_words[key_word.ToString()] = key_word.tag
	}
}

func (l *Lexer) Readch() error {
	char, err := l.reader.ReadByte()
	l.peek = char
	return err
}

func (l *Lexer) ReadCharacter(c byte) (bool, error) {
	chars, err := l.reader.Peek(1)
	if err != nil {
		return false, err
	}
	peekChar := chars[0]
	if peekChar != c {
		return false, nil
	}
	return true, nil
}

func (l *Lexer) UnRead() error {
	return l.reader.UnreadByte()
}

func (l *Lexer) Scan() (Token, error) {
	for {
		err := l.Readch()
		if err != nil {
			return NewToken(ERROR), err
		}

		if l.peek == ' ' || l.peek == '\t' {
			continue
		} else if l.peek == '\n' {
			break
		} else {
			break
		}
	}

	l.lexme = ""

	switch l.peek {
	case '{':
		l.lexme = "{"
		return NewToken(LEFT_BRACE), nil
	case '}':
		l.lexme = "}"
		return NewToken(RIGHT_BRACE), nil
	case '+':
		l.lexme = "+"
		return NewToken(PLUS), nil
	case '-':
		l.lexme = "-"
		return NewToken(MINUS), nil
	case '(':
		l.lexme = "("
		return NewToken(LEFT_BRACKET), nil
	case ')':
		l.lexme = ")"
		return NewToken(RIGHT_BRACKET), nil
	case '&':
		if ok, _ := l.ReadCharacter('&'); ok {
			l.lexme = "&&"
			word := NewToken(AND)
			return word, nil
		} else {
			l.lexme = "&"
			return NewToken(AND_OPERATOR), nil
		}
	case '|':
		if ok, _ := l.ReadCharacter('|'); ok {
			l.lexme = "||"
			word := NewToken(OR)
			return word, nil
		} else {
			l.lexme = "|"
			return NewToken(OR_OPERATOR), nil
		}
	case '=':
		if ok, _ := l.ReadCharacter('='); ok {
			l.lexme = "=="
			word := NewToken(EQ)
			return word, nil
		} else {
			l.lexme = "="
			return NewToken(ASSIGN_OPERATOR), nil
		}
	case '!':
		if ok, _ := l.ReadCharacter('='); ok {
			l.lexme = "!="
			word := NewToken(NE)
			return word, nil
		} else {
			l.lexme = "!"
			return NewToken(NEGATE_OPERATOR), nil
		}
	case '<':
		if ok, _ := l.ReadCharacter('='); ok {
			l.lexme = "<="
			word := NewToken(LE)
			return word, nil
		} else {
			l.lexme = "<"
			return NewToken(LESS_OPERATOR), nil
		}
	case '>':
		if ok, _ := l.ReadCharacter('='); ok {
			l.lexme = ">="
			word := NewToken(GE)
			return word, nil
		} else {
			l.lexme = ">"
			return NewToken(GREATER_OPERATOR), nil
		}
	}

	if unicode.IsNumber(rune(l.peek)) {
		var v int
		var err error
		for {
			num, err := strconv.Atoi(string(l.peek))
			if err != nil {
				l.UnRead()
				break
			}
			v = v*10 + num
			l.Readch()

			l.lexme += string(l.peek)
		}
		if l.peek != '.' {
			return NewToken(NUM), err
		}

		x := float64(v)
		d := float64(10)
		for {
			l.Readch()
			num, err := strconv.Atoi(string(l.peek))
			if err != nil {
				l.UnRead()
				break
			}

			x = x + float64(num)/d
			d = d * 10
			l.lexme += string(l.peek)
		}

		return NewToken(REAL), err
	}
	if unicode.IsLetter(rune(l.peek)) {
		var buffer []byte
		for {
			buffer = append(buffer, l.peek)
			l.Readch()
			if !unicode.IsLetter(rune(l.peek)) && !unicode.IsNumber(rune(l.peek)) {
				l.UnRead()
				break
			}
		}
		s := string(buffer)
		token, ok := l.key_words[s]
		if ok {
			return token, nil
		}
		return NewToken(ID), nil
	}
	return NewToken(EOF), nil
}
