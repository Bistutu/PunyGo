package lexer

import (
	"punyGo/pkg/token"
)

// Lexer 结构体定义了词法分析器的状态
type Lexer struct {
	input        string // 输入的源代码
	position     int    // 当前字符的位置（当前读取的字符）
	readPosition int    // 下一个字符的位置（即将读取的字符）
	ch           byte   // 当前读取的字符
}

// New 创建并返回一个新的 Lexer 实例
func New(input string) *Lexer {
	l := &Lexer{input: input} // 1. 初始化 Lexer 结构体，设置输入源代码
	l.readChar()              // 2. 读取第一个字符，初始化 ch、position 和 readPosition
	return l                  // 3. 返回初始化后的 Lexer 实例
}

// readChar 读取下一个字符并更新 Lexer 的状态
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // 1. 判断是否已读取到输入的末尾
		l.ch = 0 // 2. 如果是，设置当前字符为 0，表示文件结束（EOF）
	} else {
		l.ch = l.input[l.readPosition] // 3. 否则，读取下一个字符
	}
	l.position = l.readPosition // 4. 更新当前字符的位置
	l.readPosition++            // 5. 更新下一个字符的位置
}

// NextToken 解析输入并返回下一个 Token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // 1. 跳过所有空白字符

	switch l.ch { // 2. 根据当前字符决定 Token 类型
	case '=':
		if l.peekChar() == '=' { // 2.1 如果下一个字符也是 '=', 则是等于比较操作符
			ch := l.ch                                          // 2.1.1 保存当前字符
			l.readChar()                                        // 2.1.2 读取下一个字符
			literal := string(ch) + string(l.ch)                // 2.1.3 组合成 "=="
			tok = token.Token{Type: token.EQ, Literal: literal} // 2.1.4 创建 EQ Token
		} else {
			tok = newToken(token.ASSIGN, l.ch) // 2.2 否则，创建赋值操作符 Token
		}
	case '+':
		tok = newToken(token.PLUS, l.ch) // 3. 处理 '+' 操作符
	case '-':
		tok = newToken(token.MINUS, l.ch) // 4. 处理 '-' 操作符
	case '!':
		if l.peekChar() == '=' { // 5.1 如果下一个字符是 '=', 则是非等于操作符
			ch := l.ch                                              // 5.1.1 保存当前字符
			l.readChar()                                            // 5.1.2 读取下一个字符
			literal := string(ch) + string(l.ch)                    // 5.1.3 组合成 "!="
			tok = token.Token{Type: token.NOT_EQ, Literal: literal} // 5.1.4 创建 NOT_EQ Token
		} else {
			tok = newToken(token.BANG, l.ch) // 5.2 否则，创建 '!' 操作符 Token
		}
	case '/':
		tok = newToken(token.SLASH, l.ch) // 6. 处理 '/' 操作符
	case '*':
		tok = newToken(token.ASTERISK, l.ch) // 7. 处理 '*' 操作符
	case '<':
		tok = newToken(token.LT, l.ch) // 8. 处理 '<' 操作符
	case '>':
		tok = newToken(token.GT, l.ch) // 9. 处理 '>' 操作符
	case ';':
		tok = newToken(token.SEMICOLON, l.ch) // 10. 处理分号 ';'
	case ',':
		tok = newToken(token.COMMA, l.ch) // 11. 处理逗号 ','
	case '(':
		tok = newToken(token.LPAREN, l.ch) // 12. 处理左括号 '('
	case ')':
		tok = newToken(token.RPAREN, l.ch) // 13. 处理右括号 ')'
	case '{':
		tok = newToken(token.LBRACE, l.ch) // 14. 处理左大括号 '{'
	case '}':
		tok = newToken(token.RBRACE, l.ch) // 15. 处理右大括号 '}'
	case 0:
		tok.Literal = ""     // 16. 如果是 EOF，设置空字符串
		tok.Type = token.EOF // 17. 设置 Token 类型为 EOF
	default:
		if isLetter(l.ch) { // 18. 如果当前字符是字母，读取整个标识符
			tok.Literal = l.readIdentifier()          // 18.1 读取标识符
			tok.Type = token.LookupIdent(tok.Literal) // 18.2 确定标识符的 Token 类型
			return tok                                // 18.3 返回标识符 Token
		} else if isDigit(l.ch) { // 19. 如果当前字符是数字，读取整个数字
			tok.Literal = l.readNumber() // 19.1 读取数字
			tok.Type = token.INT         // 19.2 设置 Token 类型为 INT
			return tok                   // 19.3 返回数字 Token
		} else {
			tok = newToken(token.ILLEGAL, l.ch) // 20. 否则，创建非法字符 Token
		}
	}

	l.readChar() // 21. 读取下一个字符，为下一次调用做准备
	return tok   // 22. 返回当前 Token
}

// newToken 辅助函数，根据类型和字符创建一个新的 Token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier 读取一个标识符，并返回其字符串
func (l *Lexer) readIdentifier() string {
	position := l.position // 1. 记录标识符的起始位置
	for isLetter(l.ch) {   // 2. 当当前字符是字母时，继续读取
		l.readChar() // 3. 读取下一个字符
	}
	return l.input[position:l.position] // 4. 返回标识符的字符串
}

// readNumber 读取一个数字，并返回其字符串
func (l *Lexer) readNumber() string {
	position := l.position // 1. 记录数字的起始位置
	for isDigit(l.ch) {    // 2. 当当前字符是数字时，继续读取
		l.readChar() // 3. 读取下一个字符
	}
	return l.input[position:l.position] // 4. 返回数字的字符串
}

// skipWhitespace 跳过所有空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' { // 1. 判断当前字符是否为空白字符
		l.readChar() // 2. 如果是，读取下一个字符
	}
}

// peekChar 查看下一个字符，但不移动位置
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) { // 1. 判断下一个位置是否超出输入长度
		return 0 // 2. 如果是，返回 0，表示 EOF
	} else {
		return l.input[l.readPosition] // 3. 否则，返回下一个字符
	}
}

// isLetter 判断一个字符是否是字母或下划线
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// isDigit 判断一个字符是否是数字
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
