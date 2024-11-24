// pkg/object/object.go

package object

import "fmt"

// 定义对象类型的别名为字符串
type ObjectType string

// 常量定义不同的对象类型
const (
	INTEGER_OBJ      = "INTEGER"      // 整数对象
	RETURN_VALUE_OBJ = "RETURN_VALUE" // 返回值对象
	ERROR_OBJ        = "ERROR"        // 错误对象
)

// Object 接口定义了所有对象必须实现的方法
type Object interface {
	Type() ObjectType // 返回对象的类型
	Inspect() string  // 返回对象的字符串表示
}

// Integer 结构体表示整数对象
type Integer struct {
	Value int64 // 整数的值
}

// Type 方法返回对象的类型
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect 方法返回整数的字符串表示
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Error 结构体表示错误对象
type Error struct {
	Message string // 错误信息
}

// Type 方法返回对象的类型
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

// Inspect 方法返回错误的字符串表示
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Environment 结构体表示变量环境，支持嵌套
type Environment struct {
	store map[string]Object // 存储变量名到对象的映射
	outer *Environment      // 外层环境，支持嵌套作用域
}

// NewEnvironment 创建一个新的环境实例
func NewEnvironment(env *Environment) *Environment {
	s := make(map[string]Object)              // 1. 创建一个新的映射用于存储变量
	return &Environment{store: s, outer: env} // 2. 返回包含新映射和无外层环境的Environment实例
}

// Get 方法根据变量名获取对应的对象
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]   // 1. 尝试在当前环境中查找变量
	if !ok && e.outer != nil { // 2. 如果未找到且存在外层环境
		obj, ok = e.outer.Get(name) // 3. 递归在外层环境中查找变量
	}
	return obj, ok // 4. 返回找到的对象和查找状态
}

// Set 方法在环境中设置一个变量
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val // 1. 在当前环境的存储中设置变量名和对应的对象
	return val          // 2. 返回设置的对象
}

// ReturnValue 结构体表示返回值对象
type ReturnValue struct {
	Value Object // 返回的值
}

// Type 方法返回对象的类型
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Inspect 方法返回返回值的字符串表示
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
