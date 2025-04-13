package ast

import (
	"bytes"
	"monkey/token"
)

// AST를 구성하는 모든 노드는 Node 인터페이스를 구현해야한다.
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// let <identifier> = <expression>;
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // 변수 바인딩 식별자. (`let x = 1 + 2;` 에서 `x`)
	Value Expression  // 값을 생성하는 표현식. (`let x = 1 + 2;` 에서 `1 + 2`)
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

// return <expression>;
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // 표현식의 첫번째 토큰
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
