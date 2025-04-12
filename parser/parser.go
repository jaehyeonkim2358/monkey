package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer // 현재의 Lexer 인스턴스를 가리키는 포인터
	curToken  token.Token  // 현재 토큰
	peekToken token.Token  // 현재 토큰의 다음 토큰
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 토큰을 2개 읽어서 curToken과 peekToken을 세팅
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
