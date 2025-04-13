package ast

import "monkey/token"

// AST를 구성하는 모든 노드는 Node 인터페이스를 구현해야한다.
type Node interface {
	TokenLiteral() string
}

// 명령문 노드
type Statement interface {
	Node
	statementNode()
}

// 표현식 노드
type Expression interface {
	Node
	expressionNode()
}

// parser가 생성하는 모든 AST의 루트 노드
type Program struct {
	Statements []Statement
}

// Program은 노드이므로 Node를 구현해야한다.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let <identifier> = <expression>;
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // 변수 바인딩 식별자. (`let x = 1 + 2;` 에서 `x`)
	Value Expression  // 값을 생성하는 표현식. (`let x = 1 + 2;` 에서 `1 + 2`)
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// return <expression>;
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
