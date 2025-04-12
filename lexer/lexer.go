package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 입력에서 현재 위치(현재 문자를 가리킴)
	readPosition int  // 입력에서 현재 읽는 위치(현재 문자의 다음을 가리킴)
	ch           byte // 현재 조사하고 있는 문자
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	return token.Token{}
}
