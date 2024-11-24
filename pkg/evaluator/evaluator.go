package evaluator

import (
	"fmt"

	"punyGo/pkg/ast"
	"punyGo/pkg/object"
)

// Eval 函数是评估器的入口，根据节点类型调用相应的评估函数
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// 处理 Program 节点，评估程序中的所有语句
	case *ast.Program:
		return evalProgram(node.Statements, env)

	// 处理 ExpressionStatement 节点，评估其中的表达式
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	// 处理 IntegerLiteral 节点，返回对应的整数对象
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	// 处理 PrefixExpression 节点，评估前缀表达式
	case *ast.PrefixExpression:
		right := Eval(node.Right, env) // 1. 评估前缀表达式右侧的表达式
		if isError(right) {            // 2. 检查是否评估过程中产生错误
			return right // 3. 如果有错误，直接返回错误对象
		}
		return evalPrefixExpression(node.Operator, right) // 4. 根据操作符和右侧对象评估前缀表达式

	// 处理 InfixExpression 节点，评估中缀表达式
	case *ast.InfixExpression:
		left := Eval(node.Left, env) // 1. 评估中缀表达式左侧的表达式
		if isError(left) {           // 2. 检查是否评估过程中产生错误
			return left // 3. 如果有错误，直接返回错误对象
		}
		right := Eval(node.Right, env) // 4. 评估中缀表达式右侧的表达式
		if isError(right) {            // 5. 检查是否评估过程中产生错误
			return right // 6. 如果有错误，直接返回错误对象
		}
		return evalInfixExpression(node.Operator, left, right) // 7. 根据操作符、左侧和右侧对象评估中缀表达式

	// 处理 LetStatement 节点，评估变量声明和赋值
	case *ast.LetStatement:
		val := Eval(node.Value, env) // 1. 评估赋值表达式的值
		if isError(val) {            // 2. 检查是否评估过程中产生错误
			return val // 3. 如果有错误，直接返回错误对象
		}
		env.Set(node.Name.Value, val) // 4. 在环境中设置变量名和对应的值
		return nil                    // 5. 返回 nil

	// 处理 Identifier 节点，查找变量的值
	case *ast.Identifier:
		return evalIdentifier(node, env)

	// 其他未处理的节点类型
	default:
		return nil
	}
	return nil
}

// evalIdentifier 评估标识符节点，查找变量的值
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok { // 1. 在环境中查找标识符对应的值
		return val // 2. 如果找到，返回对应的对象
	}
	return newError("identifier not found: " + node.Value) // 3. 如果未找到，返回错误对象
}

// evalProgram 评估程序节点，依次评估所有语句
func evalProgram(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range stmts { // 1. 遍历所有语句
		result = Eval(statement, env)                            // 2. 评估当前语句
		if returnValue, ok := result.(*object.ReturnValue); ok { // 3. 如果是返回值对象，结束评估并返回值
			return returnValue.Value
		}
		if errObj, ok := result.(*object.Error); ok { // 4. 如果是错误对象，结束评估并返回错误
			return errObj
		}
	}

	return result // 5. 返回最后一个评估的对象
}

// evalPrefixExpression 评估前缀表达式，根据操作符调用相应的函数
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right) // 1. 处理 '!' 操作符
	case "-":
		return evalMinusPrefixOperatorExpression(right) // 2. 处理 '-' 操作符
	default:
		return newError("unknown operator: %s%s", operator, right.Type()) // 3. 未知操作符，返回错误对象
	}
}

// evalBangOperatorExpression 评估 '!' 操作符，目前仅支持整数类型
func evalBangOperatorExpression(right object.Object) object.Object {
	// 暂不处理布尔值，返回错误
	return newError("unknown operator: !%s", right.Type()) // 1. 返回未知操作符错误
}

// evalMinusPrefixOperatorExpression 评估 '-' 操作符，对整数取反
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ { // 1. 检查右侧对象是否为整数类型
		return newError("unknown operator: -%s", right.Type()) // 2. 如果不是，返回错误对象
	}

	value := right.(*object.Integer).Value // 3. 获取整数值
	return &object.Integer{Value: -value}  // 4. 返回取反后的整数对象
}

// evalInfixExpression 评估中缀表达式，根据操作符和操作数类型调用相应的函数
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right) // 1. 如果左右都是整数，调用整数中缀表达式评估
	default:
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type()) // 2. 类型不匹配，返回错误对象
	}
}

// evalIntegerInfixExpression 评估整数类型的中缀表达式，执行具体的算术操作
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value   // 1. 获取左侧整数值
	rightVal := right.(*object.Integer).Value // 2. 获取右侧整数值

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal} // 3. 执行加法
	case "-":
		return &object.Integer{Value: leftVal - rightVal} // 4. 执行减法
	case "*":
		return &object.Integer{Value: leftVal * rightVal} // 5. 执行乘法
	case "/":
		return &object.Integer{Value: leftVal / rightVal} // 6. 执行除法
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type()) // 7. 未知操作符，返回错误对象
	}
}

// newError 创建一个新的错误对象，包含格式化的错误消息
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)} // 1. 使用 fmt.Sprintf 格式化错误消息
}

// isError 检查一个对象是否为错误对象
func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ // 1. 如果对象不为 nil，检查其类型是否为 ERROR_OBJ
	}
	return false // 2. 如果对象为 nil，返回 false
}
