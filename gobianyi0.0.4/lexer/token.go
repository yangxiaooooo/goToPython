package lexer

type Tag uint32

const (
	AND   Tag = iota + 256 // &&
	BREAK                  // break
	DO
	EQ               // ==
	FALSE            // flase
	GE               // >=
	LE               // <=
	ID               // IDENTIFIER
	IF               // if
	ELSE             // else
	MINUS            // -
	PLUS             // +
	NE               // !=
	NUM              //
	OR               // ||
	REAL             // 3.14
	STRING           // "abc"
	TRUE             // true
	WHILE            // while
	LEFT_BRACE       // {
	RIGHT_BRACE      // }
	AND_OPERATOR     // &
	OR_OPERATOR      // |
	ASSIGN_OPERATOR  // =
	NEGATE_OPERATOR  // !
	LESS_OPERATOR    // <
	GREATER_OPERATOR // >
	LEFT_BRACKET     // (
	RIGHT_BRACKET    // )
	EOF              // end of file
	ERROR            //
	SEMICOLON        // ;
	BOOL
	INT
	CHAR
	TYPE
)

var token_map = make(map[Tag]string)

func init() {
	token_map[AND] = "&&"
	token_map[DO] = "do"
	token_map[ELSE] = "else"
	token_map[EQ] = "=="
	token_map[FALSE] = "false"
	token_map[GE] = ">="
	token_map[LE] = "<="
	token_map[ID] = "identifier"
	token_map[IF] = "if"
	token_map[MINUS] = "-"
	token_map[PLUS] = "+"
	token_map[NE] = "NUM"
	token_map[OR] = "OR"
	token_map[REAL] = "REAL"
	token_map[WHILE] = "while"
	token_map[AND_OPERATOR] = "&"
	token_map[OR_OPERATOR] = "|"
	token_map[ASSIGN_OPERATOR] = "="
	token_map[NEGATE_OPERATOR] = "!"
	token_map[LESS_OPERATOR] = "<"
	token_map[GREATER_OPERATOR] = ">"
	token_map[LEFT_BRACE] = "{"
	token_map[RIGHT_BRACE] = "}"
	token_map[ERROR] = "ERROR"
	token_map[EOF] = "EOF"
	token_map[STRING] = "string"
	token_map[LEFT_BRACKET] = "("
	token_map[RIGHT_BRACKET] = "0"
	//token_map[BOOL] = "TYPE"
	//token_map[INT] = "TYPE"
	//token_map[CHAR] = "TYPE"

}

type Token struct {
	Tag Tag
}

func (t *Token) ToString() string {
	return token_map[t.Tag]
}

func NewToken(tag Tag) Token {
	return Token{
		Tag: tag,
	}
}
