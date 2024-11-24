# PunyGo

> 中文|[Engilsh](./README.md)

PunyGo 是一个用 Go 语言实现的极简编程语言，用于方便开发者理解编译器的基本原理和实现过程。它包含了一个简单的词法分析器、语法分析器、抽象语法树（AST）、解释器以及 REPL（读取-求值-输出循环）。该语言目前支持整数算术运算、变量赋值以及使用括号控制运算优先级。

本项目受《[Writing An Interpreter In Go](https://interpreterbook.com/)》一书的启发，旨在为构建解释器和编译器提供一个学习资源。

---

## 目录

- [功能特性](#功能特性)
- [快速开始](#快速开始)
    - [前置条件](#前置条件)
    - [安装](#安装)
    - [运行 punyGo](#运行-punygo)
- [使用示例](#使用示例)
    - [示例代码](#示例代码)
- [项目结构](#项目结构)
- [贡献指南](#贡献指南)
- [许可证](#许可证)

---

## 功能特性

- **整数算术运算**：支持 `+`、`-`、`*`、`/` 操作符
- **变量赋值**：使用 `let` 关键字进行变量声明和赋值
- **括号优先级**：使用括号控制运算顺序
- **REPL**：提供交互式编程环境
- **简单的词法分析器和语法分析器**
- **抽象语法树（AST）表示**

## 快速开始

### 前置条件

安装了 Go 编程语言（版本 1.17 或更高），您可以从 [Go 官方网站](https://golang.org/dl/) 下载并安装 Go。

### 安装

1. **克隆仓库：**

   ```bash
   git clone https://github.com/yourusername/punyGo.git
   cd punyGo
   ```

2. **构建项目：**

   使用提供的 `Makefile` 来构建项目。

   ```bash
   make build
   ```

   这将编译源代码，并在 `bin/` 目录下生成可执行文件 `punyGo`。

### 运行 punyGo

要启动 punyGo REPL，请运行：

```bash
make run
```

或者直接执行二进制文件：

```bash
./bin/punyGo
```

您将看到以下提示：

```bash
Hello yourusername! This is the punyGo programming language!
Feel free to type in commands
>>
```

## 使用示例

在命令行中，您可以输入 punyGo 代码并立即看到结果。

### 示例代码

1. **基本算术运算：**

   ```plaintext
   >> 1 + 2 * 3;
   (1 + (2 * 3));
   7
   ```

2. **使用括号控制优先级：**

   ```plaintext
   >> (1 + 2) * 3;
   ((1 + 2) * 3);
   9
   ```

3. **变量赋值：**

   ```plaintext
   >> let x = 5;
   let x = 5;
   >> x;
   x;
   5
   ```

4. **变量运算：**

   ```plaintext
   >> let y = x + 10;
   let y = (x + 10);
   >> y;
   y;
   15
   >> x + y;
   (x + y);
   20
   ```

5. **复杂表达式：**

   ```plaintext
   >> let result = (x + y) * 2;
   let result = ((x + y) * 2);
   >> result;
   result;
   40
   ```

## 项目结构

```shell
punyGo/
├── Makefile                # 构建和运行命令
├── go.mod
├── pkg/
│   ├── ast/
│   │   └── ast.go          # 定义抽象语法树（AST）的节点
│   ├── evaluator/
│   │   └── evaluator.go    # 解释器，用于评估 AST
│   ├── lexer/
│   │   └── lexer.go        # 词法分析器，将输入代码分解为标记
│   ├── object/
│   │   └── object.go       # 定义对象系统（整数、环境等）
│   ├── parser/
│   │   └── parser.go       # 语法分析器，构建 AST
│   ├── repl/
│   │   └── repl.go         # 读取-求值-输出循环的实现
│   └── token/
│       └── token.go         # 定义词法分析器和语法分析器使用的标记
└─── main.go		# 应用程序的主入口点
```

- **`Makefile`**：包含构建和运行项目的命令，方便执行常用操作。
- **`main.go`**：应用程序的主入口点，启动 REPL 并处理用户输入。
- **`pkg/`**：包含 punyGo 的核心代码。
  - **`ast/`**：抽象语法树（AST）的定义，描述程序的结构。
    - **`ast.go`**：定义了所有 AST 节点的结构体和方法，用于表示和遍历程序的语法结构。
  - **`evaluator/`**：解释器，实现了对 AST 的评估。
    - **`evaluator.go`**：包含评估表达式、处理变量赋值和环境的核心逻辑。
  - **`lexer/`**：词法分析器，将源代码转换为标记（tokens）。
    - **`lexer.go`**：实现了扫描源代码、识别关键字、标识符、数字和操作符等。
  - **`object/`**：定义对象系统，包括整数、环境等。
    - **`object.go`**：定义了对象接口和具体对象类型，如整数、布尔值、环境和错误对象。
  - **`parser/`**：语法分析器，将标记序列解析为 AST。
    - **`parser.go`**：实现了递归下降解析器，处理表达式、前缀和中缀运算符等。
  - **`repl/`**：读取-求值-输出循环，实现交互式编程环境。
    - **`repl.go`**：处理用户输入，调用词法分析器、语法分析器和解释器，并输出结果。
  - **`token/`**：定义词法分析器和语法分析器使用的标记类型。
    - **`token.go`**：定义了标记类型、标记结构体，以及关键字到标记类型的映射。

## 项目案例

当你输入一条语句（例如 `let x = 5 + 5;`）时，整个流程涉及词法分析器（Lexer）、语法解析器（Parser）、抽象语法树（AST）、评估器（Evaluator）以及对象系统（Object System）等多个组件。词法分析器负责将源代码分解为 Token，语法解析器构建 AST，评估器遍历 AST 并执行相应的操作，最终实现程序的执行和输出。这种模块化的设计使得解释器具有良好的可维护性和扩展性，便于支持更多的语言特性和复杂的语法结构。

以下是详细的流程说明：

### 1. 输入源代码

你输入的语句 `let x = 5 + 5;` 作为源代码传递给解释器。

### 2. 词法分析（Lexing）

**组件**：`lexer` 包中的 `Lexer` 结构体和相关方法。

**流程**：

1. **初始化 Lexer**：
   - 调用 `lexer.New(input)` 创建一个新的 `Lexer` 实例，并初始化输入源代码。
   - `readChar` 方法被调用，读取第一个字符 `l.ch`。

2. **分解成 Token**：
   - `NextToken` 方法被反复调用，逐个字符地将输入源代码分解成一系列的词法单元（Token）。
   - 对于 `let x = 5 + 5;`，词法分析器将生成以下 Token 序列：
     - `LET`（关键字）
     - `IDENT`（标识符 `x`）
     - `ASSIGN`（赋值操作符 `=`）
     - `INT`（整数 `5`）
     - `PLUS`（加号 `+`）
     - `INT`（整数 `5`）
     - `SEMICOLON`（分号 `;`）
     - `EOF`（文件结束标记）

**示例**：

```go
lexer := lexer.New("let x = 5 + 5;")
tokens := []token.Token{}
for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
    tokens = append(tokens, tok)
}
```

### 3. 解析（Parsing）

**组件**：`parser` 包中的 `Parser` 结构体和相关方法。

**流程**：

1. **初始化 Parser**：
   - 使用词法分析器生成的 Token 序列创建一个新的 `Parser` 实例：`parser.New(lexer)`.
   - 注册前缀和中缀解析函数，以便解析不同类型的表达式和操作符。

2. **构建 AST**：
   - 调用 `ParseProgram` 方法，开始解析整个程序。
   - `parseStatement` 方法根据当前 Token 类型解析不同类型的语句。
   - 对于 `let x = 5 + 5;`，解析器将构建以下 AST 结构：

```bash
Program
 └── LetStatement
      ├── Name: Identifier(x)
      └── Value: InfixExpression
           ├── Left: IntegerLiteral(5)
           ├── Operator: +
           └── Right: IntegerLiteral(5)
```

**示例**：

```go
parser := parser.New(lexer)
ast := parser.ParseProgram()
if len(parser.Errors()) != 0 {
    // 处理解析错误
}
```

### 4. 评估（Evaluation）

**组件**：`evaluator` 包中的 `Evaluator` 以及 `object` 包中的对象系统。

**流程**：

1. **初始化环境**：
   - 创建一个新的环境实例，用于存储变量和它们的值：`env := object.NewEnvironment()`。

2. **遍历 AST 并评估**：
   - 调用 `evaluator.Eval(ast, env)` 开始评估 AST。
   - 评估过程如下：
     - **Program 节点**：
       - 遍历其子节点 `LetStatement`。
     - **LetStatement 节点**：
       - 评估赋值表达式 `5 + 5`：
         - **InfixExpression 节点**：
           - 评估左操作数 `5`，得到 `Integer` 对象。
           - 评估右操作数 `5`，得到 `Integer` 对象。
           - 执行加法操作，得到 `Integer` 对象 `10`。
       - 将变量 `x` 绑定到 `Integer` 对象 `10` 中，存储在环境中。

**示例**：

```go
result := evaluator.Eval(ast, env)
if errObj, ok := result.(*object.Error); ok {
    fmt.Println(errObj.Message)
} else {
    // 处理其他类型的结果
}
```

### 5. 输出结果

对于 `let` 语句，本身不产生直接的输出，因为它是变量声明和赋值操作。然而，变量 `x` 已经被存储在环境中，并且可以在后续的表达式或语句中使用。

**示例**：

```go
// 假设后续有表达式 `x + 10;`
lexer := lexer.New("let x = 5 + 5; x + 10;")
parser := parser.New(lexer)
ast := parser.ParseProgram()
env := object.NewEnvironment()
result := evaluator.Eval(ast, env)
if errObj, ok := result.(*object.Error); ok {
    fmt.Println(errObj.Message)
} else {
    fmt.Println(result.Inspect()) // 输出: 15
}
```

### 6. 整体流程总结

1. **输入源代码**：

   `let x = 5 + 5;`

2. **词法分析**：

   分解成 Token 序列：`LET`, `IDENT(x)`, `ASSIGN`, `INT(5)`, `PLUS`, `INT(5)`, `SEMICOLON`, `EOF`.

3. **语法分析**：

   构建 AST
   ```bash
   Program
    └── LetStatement
         ├── Name: Identifier(x)
         └── Value: InfixExpression(+)
              ├── Left: IntegerLiteral(5)
              ├── Operator: +
              └── Right: IntegerLiteral(5)
   ```

4. **评估**：
   - 评估表达式 `5 + 5`，得到 `10`。
   - 将 `x` 绑定到 `10`。

5. **输出**：
   - `let` 语句不直接输出，但变量 `x` 可用于后续操作。

### 7. 进一步扩展

如果在源代码中有更多的语句或表达式，评估器会继续遍历 AST，依次评估每个节点，并根据需要更新环境。例如：

```go
let x = 5 + 5;
x + 10;
```

在评估 `x + 10;` 时，评估器会从环境中获取 `x` 的值 `10`，然后执行加法操作，最终输出 `20`。

### 8. 错误处理

在整个流程中，如果遇到词法分析、解析或评估过程中的错误（例如，未定义的变量、类型不匹配、语法错误等），相应的错误信息会被记录并输出，确保程序的健壮性和易于调试。

## 贡献指南

欢迎贡献！请随时提交拉取请求或打开问题（issue）。

可贡献的领域包括但不限于：

- **错误处理**：改进词法分析器、语法分析器和解释器的错误信息。
- **新功能**：
    - 实现布尔类型和操作。
    - 添加条件语句（`if-else`）。
    - 支持函数定义和调用，以及作用域管理。
    - 扩展对象系统以包含更多数据类型（字符串、数组等）。
- **优化**：提升解释器的性能。
- **文档**：增强文档并添加更多示例。

## 许可证

本项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE) 文件。