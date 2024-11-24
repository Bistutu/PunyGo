package ast

import (
	"bytes"

	"punyGo/pkg/token"
)

// Node 接口是所有 AST 节点的基础接口
type Node interface {
	TokenLiteral() string // 返回该节点对应的词法单元的字面量
	String() string       // 返回该节点的字符串表示
}

// Statement 接口表示所有语句节点
type Statement interface {
	Node
	statementNode() // 标识这是一个语句节点的方法
}

// Expression 接口表示所有表达式节点
type Expression interface {
	Node
	expressionNode() // 标识这是一个表达式节点的方法
}

// Program 程序根节点，包含一系列语句
type Program struct {
	Statements []Statement // 程序中的语句列表
}

// statementNode 实现 Statement 接口，用于标识 Program 是一个语句节点
func (p *Program) statementNode() {}

// TokenLiteral 返回程序中第一个语句的词法字面量，如果没有语句则返回空字符串
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String 返回程序中所有语句的字符串连接
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement 代表 let 语句节点
type LetStatement struct {
	Token token.Token // token.LET 词法单元
	Name  *Identifier // 变量名
	Value Expression  // 变量值表达式
}

// statementNode 实现 Statement 接口，用于标识 LetStatement 是一个语句节点
func (ls *LetStatement) statementNode() {}

// TokenLiteral 返回 let 语句的词法字面量
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String 返回 let 语句的字符串表示，例如 "let x = 5;"
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ") // 1. 写入 "let "
	out.WriteString(ls.Name.String())        // 2. 写入变量名
	out.WriteString(" = ")                   // 3. 写入 " = "

	if ls.Value != nil {
		out.WriteString(ls.Value.String()) // 4. 写入变量值的字符串表示
	}

	out.WriteString(";") // 5. 写入分号

	return out.String()
}

// ExpressionStatement 代表表达式语句节点
type ExpressionStatement struct {
	Token      token.Token // 表达式中的第一个词法单元
	Expression Expression  // 表达式
}

// statementNode 实现 Statement 接口，用于标识 ExpressionStatement 是一个语句节点
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral 返回表达式语句的词法字面量
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String 返回表达式语句的字符串表示
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String() // 1. 返回表达式的字符串表示
	}
	return ""
}

// Identifier 代表标识符节点
type Identifier struct {
	Token token.Token // token.IDENT 词法单元
	Value string      // 标识符的名称
}

// expressionNode 实现 Expression 接口，用于标识 Identifier 是一个表达式节点
func (i *Identifier) expressionNode() {}

// TokenLiteral 返回标识符的词法字面量
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String 返回标识符的名称
func (i *Identifier) String() string { return i.Value }

// IntegerLiteral 代表整数字面量节点
type IntegerLiteral struct {
	Token token.Token // 整数字面量的词法单元
	Value int64       // 整数的值
}

// expressionNode 实现 Expression 接口，用于标识 IntegerLiteral 是一个表达式节点
func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral 返回整数字面量的词法字面量
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// String 返回整数字面量的字符串表示
func (il *IntegerLiteral) String() string { return il.Token.Literal }

// PrefixExpression 代表前缀表达式节点，例如 !5 或 -a
type PrefixExpression struct {
	Token    token.Token // 操作符的词法单元，如 '!' 或 '-'
	Operator string      // 操作符的字符串表示
	Right    Expression  // 操作符右侧的表达式
}

// expressionNode 实现 Expression 接口，用于标识 PrefixExpression 是一个表达式节点
func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral 返回前缀表达式的操作符词法字面量
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String 返回前缀表达式的字符串表示，例如 "(!5)"
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")               // 1. 写入左括号
	out.WriteString(pe.Operator)       // 2. 写入操作符
	out.WriteString(pe.Right.String()) // 3. 写入右侧表达式的字符串表示
	out.WriteString(")")               // 4. 写入右括号

	return out.String()
}

// InfixExpression 代表中缀表达式节点，例如 5 + 5
type InfixExpression struct {
	Token    token.Token // 操作符的词法单元，如 '+'
	Left     Expression  // 操作符左侧的表达式
	Operator string      // 操作符的字符串表示
	Right    Expression  // 操作符右侧的表达式
}

// expressionNode 实现 Expression 接口，用于标识 InfixExpression 是一个表达式节点
func (ie *InfixExpression) expressionNode() {}

// TokenLiteral 返回中缀表达式的操作符词法字面量
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String 返回中缀表达式的字符串表示，例如 "(5 + 5)"
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")                     // 1. 写入左括号
	out.WriteString(ie.Left.String())        // 2. 写入左侧表达式的字符串表示
	out.WriteString(" " + ie.Operator + " ") // 3. 写入操作符，前后加空格
	out.WriteString(ie.Right.String())       // 4. 写入右侧表达式的字符串表示
	out.WriteString(")")                     // 5. 写入右括号

	return out.String()
}
