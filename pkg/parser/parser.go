package parser

import (
	"fmt"
	"strconv"

	"punyGo/pkg/ast"
	"punyGo/pkg/lexer"
	"punyGo/pkg/token"
)

// 定义操作符优先级
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > 或 <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X 或 !X
	CALL        // myFunction(X)
)

// 定义每个Token类型对应的优先级
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

// 定义前缀解析函数类型
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Parser 结构体定义了解析器的状态和方法
type Parser struct {
	lex    *lexer.Lexer // 词法分析器
	errors []string     // 解析过程中产生的错误

	curToken  token.Token // 当前处理的Token
	peekToken token.Token // 下一个Token

	prefixParseFns map[token.TokenType]prefixParseFn // 前缀解析函数映射
	infixParseFns  map[token.TokenType]infixParseFn  // 中缀解析函数映射
}

// New 创建并返回一个新的 Parser 实例
func New(lex *lexer.Lexer) *Parser {
	p := &Parser{
		lex:    lex,
		errors: []string{},
	}

	// 初始化前缀解析函数映射
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)         // 1. 注册标识符解析函数
	p.registerPrefix(token.INT, p.parseIntegerLiteral)       // 2. 注册整数字面量解析函数
	p.registerPrefix(token.BANG, p.parsePrefixExpression)    // 3. 注册逻辑非解析函数
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)   // 4. 注册负号解析函数
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression) // 5. 注册分组表达式解析函数

	// 初始化中缀解析函数映射
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)     // 1. 注册加法解析函数
	p.registerInfix(token.MINUS, p.parseInfixExpression)    // 2. 注册减法解析函数
	p.registerInfix(token.SLASH, p.parseInfixExpression)    // 3. 注册除法解析函数
	p.registerInfix(token.ASTERISK, p.parseInfixExpression) // 4. 注册乘法解析函数
	p.registerInfix(token.EQ, p.parseInfixExpression)       // 5. 注册等于比较解析函数
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)   // 6. 注册不等于比较解析函数
	p.registerInfix(token.LT, p.parseInfixExpression)       // 7. 注册小于比较解析函数
	p.registerInfix(token.GT, p.parseInfixExpression)       // 8. 注册大于比较解析函数

	// 读取两个Token，初始化curToken和peekToken
	p.nextToken() // 1. 读取第一个Token
	p.nextToken() // 2. 读取第二个Token

	return p // 3. 返回初始化后的Parser实例
}

// ParseProgram 开始解析程序，返回AST的根节点Program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}              // 1. 创建一个新的Program节点
	program.Statements = []ast.Statement{} // 2. 初始化语句列表

	for p.curToken.Type != token.EOF { // 3. 当当前Token不是EOF时
		stmt := p.parseStatement() // 3.1. 解析当前语句
		if stmt != nil {           // 3.2. 如果解析成功
			program.Statements = append(program.Statements, stmt) // 3.2.1. 将语句添加到Program节点
		}
		p.nextToken() // 3.3. 前进到下一个Token
	}

	return program // 4. 返回Program节点
}

// Errors 返回解析过程中产生的错误
func (p *Parser) Errors() []string {
	return p.errors
}

// nextToken 前进到下一个Token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken        // 1. 将peekToken赋值给curToken
	p.peekToken = p.lex.NextToken() // 2. 从Lexer获取下一个Token，赋值给peekToken
}

// curTokenIs 检查当前Token的类型是否与给定类型匹配
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs 检查下一个Token的类型是否与给定类型匹配
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek 检查下一个Token的类型是否与给定类型匹配，若匹配则前进到下一个Token
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) { // 1. 如果下一个Token类型匹配
		p.nextToken() // 1.1. 前进到下一个Token
		return true   // 1.2. 返回真
	} else {
		p.peekError(t) // 2. 如果不匹配，记录错误
		return false   // 3. 返回假
	}
}

// peekError 记录期待的Token类型与实际不匹配的错误
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type) // 1. 构建错误消息
	p.errors = append(p.errors, msg)                                                        // 2. 将错误消息添加到错误列表
}

// noPrefixParseFnError 记录缺少前缀解析函数的错误
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t) // 1. 构建错误消息
	p.errors = append(p.errors, msg)                               // 2. 将错误消息添加到错误列表
}

// registerPrefix 注册前缀解析函数
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn // 1. 将前缀解析函数与Token类型关联
}

// registerInfix 注册中缀解析函数
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn // 1. 将中缀解析函数与Token类型关联
}

// parseStatement 根据当前Token解析不同类型的语句
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET: // 1. 如果是let语句
		return p.parseLetStatement() // 1.1. 解析let语句
	default: // 2. 默认解析为表达式语句
		return p.parseExpressionStatement()
	}
}

// parseLetStatement 解析let语句
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken} // 1. 创建一个新的LetStatement节点，记录当前Token

	if !p.expectPeek(token.IDENT) { // 2.1. 期待下一个Token是标识符
		return nil // 2.2. 如果不是，返回nil
	}

	stmt.Name = &ast.Identifier{ // 3. 设置变量名
		Token: p.curToken,         // 3.1. 当前Token
		Value: p.curToken.Literal, // 3.2. 变量名的字面量
	}

	if !p.expectPeek(token.ASSIGN) { // 4.1. 期待下一个Token是赋值操作符
		return nil // 4.2. 如果不是，返回nil
	}

	p.nextToken() // 5. 前进到下一个Token，开始解析赋值表达式

	stmt.Value = p.parseExpression(LOWEST) // 6. 解析赋值表达式，优先级最低

	if p.peekTokenIs(token.SEMICOLON) { // 7. 如果下一个Token是分号
		p.nextToken() // 7.1. 前进到分号
	}

	return stmt // 8. 返回解析后的LetStatement节点
}

// parseExpressionStatement 解析表达式语句
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken} // 1. 创建一个新的ExpressionStatement节点，记录当前Token

	stmt.Expression = p.parseExpression(LOWEST) // 2. 解析表达式，优先级最低

	if p.peekTokenIs(token.SEMICOLON) { // 3. 如果下一个Token是分号
		p.nextToken() // 3.1. 前进到分号
	}

	return stmt // 4. 返回解析后的ExpressionStatement节点
}

// parseExpression 解析表达式，根据当前Token的优先级决定解析顺序
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type] // 1. 获取当前Token对应的前缀解析函数
	if prefix == nil {                          // 2. 如果没有对应的前缀解析函数
		p.noPrefixParseFnError(p.curToken.Type) // 2.1. 记录错误
		return nil                              // 2.2. 返回nil
	}
	leftExp := prefix() // 3. 调用前缀解析函数，获取左表达式

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() { // 4. 循环处理中缀表达式
		infix := p.infixParseFns[p.peekToken.Type] // 4.1. 获取下一个Token对应的中缀解析函数
		if infix == nil {                          // 4.2. 如果没有对应的中缀解析函数
			return leftExp // 4.2.1. 返回当前的左表达式
		}

		p.nextToken() // 5. 前进到下一个Token，准备解析中缀表达式

		leftExp = infix(leftExp) // 6. 解析中缀表达式，更新左表达式
	}

	return leftExp // 7. 返回解析后的表达式
}

// parseIdentifier 解析标识符，返回Identifier节点
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,         // 1. 当前Token
		Value: p.curToken.Literal, // 2. 标识符的名称
	}
}

// parseIntegerLiteral 解析整数字面量，返回IntegerLiteral节点
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken} // 1. 创建一个新的IntegerLiteral节点，记录当前Token

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64) // 2. 将字面量转换为整数
	if err != nil {                                           // 3. 如果转换失败
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal) // 3.1. 构建错误消息
		p.errors = append(p.errors, msg)                                        // 3.2. 记录错误
		return nil                                                              // 3.3. 返回nil
	}

	lit.Value = value // 4. 设置整数值
	return lit        // 5. 返回解析后的IntegerLiteral节点
}

// parsePrefixExpression 解析前缀表达式，返回PrefixExpression节点
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,         // 1. 当前Token
		Operator: p.curToken.Literal, // 2. 操作符
	}

	p.nextToken() // 3. 前进到下一个Token，解析右侧表达式

	expression.Right = p.parseExpression(PREFIX) // 4. 解析右侧表达式，优先级为PREFIX

	return expression // 5. 返回解析后的PrefixExpression节点
}

// parseInfixExpression 解析中缀表达式，返回InfixExpression节点
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,         // 1. 当前Token
		Operator: p.curToken.Literal, // 2. 操作符
		Left:     left,               // 3. 左侧表达式
	}

	precedence := p.curPrecedence()                  // 4. 获取当前操作符的优先级
	p.nextToken()                                    // 5. 前进到下一个Token，解析右侧表达式
	expression.Right = p.parseExpression(precedence) // 6. 解析右侧表达式，优先级为当前操作符的优先级

	return expression // 7. 返回解析后的InfixExpression节点
}

// parseGroupedExpression 解析分组表达式（括号内的表达式），返回解析后的表达式节点
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken() // 1. 前进到下一个Token，解析括号内的表达式

	exp := p.parseExpression(LOWEST) // 2. 解析括号内的表达式，优先级最低

	if !p.expectPeek(token.RPAREN) { // 3. 期待下一个Token是右括号
		return nil // 4. 如果不是，返回nil
	}

	return exp // 5. 返回解析后的表达式
}

// peekPrecedence 获取下一个Token的优先级
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok { // 1. 如果下一个Token类型有定义优先级
		return p // 1.1. 返回对应的优先级
	}
	return LOWEST // 2. 否则，返回最低优先级
}

// curPrecedence 获取当前Token的优先级
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok { // 1. 如果当前Token类型有定义优先级
		return p // 1.1. 返回对应的优先级
	}
	return LOWEST // 2. 否则，返回最低优先级
}
