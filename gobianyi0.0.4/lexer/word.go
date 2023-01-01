package lexer

type Word struct {
	lexeme string
	tag    Token
}

func NewWordToken(s string, tag Tag) Word {
	return Word{
		lexeme: s,
		tag:    NewToken(tag),
	}
}

func GetKeyWords() []Word {
	key_words := []Word{}
	key_words = append(key_words, NewWordToken("&&", AND))
	key_words = append(key_words, NewWordToken("||", OR))
	key_words = append(key_words, NewWordToken("==", EQ))
	key_words = append(key_words, NewWordToken("!=", NE))
	key_words = append(key_words, NewWordToken("<=", LE))
	key_words = append(key_words, NewWordToken(">=", GE))
	key_words = append(key_words, NewWordToken("minus", MINUS))
	key_words = append(key_words, NewWordToken("true", TRUE))
	key_words = append(key_words, NewWordToken("false", FALSE))
	key_words = append(key_words, NewWordToken("+", PLUS))
	key_words = append(key_words, NewWordToken("if", IF))
	key_words = append(key_words, NewWordToken("else", ELSE))
	key_words = append(key_words, NewWordToken("int", TYPE))
	key_words = append(key_words, NewWordToken("bool", TYPE))
	key_words = append(key_words, NewWordToken("char", TYPE))
	key_words = append(key_words, NewWordToken(";", SEMICOLON))

	return key_words
}

func (w *Word) ToString() string {
	return w.lexeme
}
