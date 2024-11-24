# punyGo

> [中文](./README_ZH.md)|Engilsh

punyGo is a minimalist programming language implemented in Go. It includes a simple lexer, parser, Abstract Syntax Tree (AST), interpreter, and a REPL (Read-Eval-Print Loop). The language currently supports integer arithmetic operations, variable assignment, and the use of parentheses to control operation precedence.

This project is inspired by the book "[Writing An Interpreter In Go](https://interpreterbook.com/)" and aims to provide a learning resource for building interpreters and compilers.

---

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running punyGo](#running-punygo)
- [Usage Examples](#usage-examples)
  - [Sample Code](#sample-code)
- [Project Structure](#project-structure)
- [Project Case Study](#project-case-study)
- [Contribution Guidelines](#contribution-guidelines)
- [License](#license)

---

## Features

- **Integer Arithmetic Operations**: Supports `+`, `-`, `*`, `/` operators
- **Variable Assignment**: Use the `let` keyword for variable declaration and assignment
- **Parentheses Precedence**: Use parentheses to control the order of operations
- **REPL**: Provides an interactive programming environment
- **Simple Lexer and Parser**
- **Abstract Syntax Tree (AST) Representation**

## Getting Started

### Prerequisites

Go programming language (version 1.17 or higher) installed. You can download and install Go from the [official Go website](https://golang.org/dl/).

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/bistutu/punyGo.git
   cd punyGo
   ```

2. **Build the project:**

   Use the provided `Makefile` to build the project.

   ```bash
   make build
   ```

   This will compile the source code and generate the executable `punyGo` in the `bin/` directory.

### Running punyGo

To start the punyGo REPL, run:

```bash
make run
```

Or execute the binary directly:

```bash
./bin/punyGo
```

You will see the following prompt:

```bash
Hello yourusername! This is the punyGo programming language!
Feel free to type in commands
>>
```

## Usage Examples

In the command line, you can input punyGo code and see the results immediately.

### Sample Code

1. **Basic Arithmetic Operations:**

   ```plaintext
   >> 1 + 2 * 3;
   (1 + (2 * 3));
   7
   ```

2. **Using Parentheses to Control Precedence:**

   ```plaintext
   >> (1 + 2) * 3;
   ((1 + 2) * 3);
   9
   ```

3. **Variable Assignment:**

   ```plaintext
   >> let x = 5;
   let x = 5;
   >> x;
   x;
   5
   ```

4. **Variable Operations:**

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

5. **Complex Expressions:**

   ```plaintext
   >> let result = (x + y) * 2;
   let result = ((x + y) * 2);
   >> result;
   result;
   40
   ```

## Project Structure

```shell
punyGo/
├── Makefile                # Build and run commands
├── go.mod
├── pkg/
│   ├── ast/
│   │   └── ast.go          # Defines nodes of the Abstract Syntax Tree (AST)
│   ├── evaluator/
│   │   └── evaluator.go    # Interpreter for evaluating the AST
│   ├── lexer/
│   │   └── lexer.go        # Lexer that tokenizes the input code
│   ├── object/
│   │   └── object.go       # Defines the object system (integers, environment, etc.)
│   ├── parser/
│   │   └── parser.go       # Parser that builds the AST
│   ├── repl/
│   │   └── repl.go         # Implementation of the Read-Eval-Print Loop
│   └── token/
│       └── token.go        # Defines tokens used by the lexer and parser
└── main.go                 # Main entry point of the application
```

- **`Makefile`**: Contains commands to build and run the project, making it easy to perform common operations.
- **`main.go`**: The main entry point of the application, which starts the REPL and handles user input.
- **`pkg/`**: Contains the core code of punyGo.
  - **`ast/`**: Definitions of the Abstract Syntax Tree (AST), describing the structure of the program.
    - **`ast.go`**: Defines all AST node structs and methods, used to represent and traverse the syntax structure of the program.
  - **`evaluator/`**: The interpreter that evaluates the AST.
    - **`evaluator.go`**: Contains the core logic for evaluating expressions, handling variable assignment, and environment.
  - **`lexer/`**: The lexer that converts source code into tokens.
    - **`lexer.go`**: Implements scanning of source code, recognizing keywords, identifiers, numbers, operators, etc.
  - **`object/`**: Defines the object system, including integers, environment, etc.
    - **`object.go`**: Defines the object interface and concrete object types, such as integers, booleans, environment, and error objects.
  - **`parser/`**: The parser that parses token sequences into the AST.
    - **`parser.go`**: Implements a recursive descent parser, handling expressions, prefix and infix operators, etc.
  - **`repl/`**: The Read-Eval-Print Loop, implementing the interactive programming environment.
    - **`repl.go`**: Handles user input, invokes the lexer, parser, and interpreter, and outputs results.
  - **`token/`**: Defines token types used by the lexer and parser.
    - **`token.go`**: Defines token types, the token struct, and the mapping from keywords to token types.

## Project Case Study

When you input a statement (e.g., `let x = 5 + 5;`), the entire process involves multiple components such as the lexer, parser, abstract syntax tree (AST), evaluator, and object system. The lexer is responsible for breaking down the source code into tokens, the parser constructs the AST, the evaluator traverses the AST and executes the corresponding operations, ultimately achieving program execution and output. This modular design makes the interpreter highly maintainable and extensible, facilitating the support of more language features and complex syntax structures.

Below is a detailed flow description:

### 1. Input Source Code

The statement you input, `let x = 5 + 5;`, is passed to the interpreter as source code.

### 2. Lexical Analysis (Lexing)

**Component**: `Lexer` struct and related methods in the `lexer` package.

**Process**:

1. **Initialize Lexer**:
   - Call `lexer.New(input)` to create a new `Lexer` instance and initialize it with the input source code.
   - The `readChar` method is called to read the first character into `l.ch`.

2. **Tokenization**:
   - The `NextToken` method is repeatedly called to sequentially convert the input source code into a series of lexical tokens.
   - For `let x = 5 + 5;`, the lexer generates the following token sequence:
     - `LET` (keyword)
     - `IDENT` (identifier `x`)
     - `ASSIGN` (assignment operator `=`)
     - `INT` (integer `5`)
     - `PLUS` (plus sign `+`)
     - `INT` (integer `5`)
     - `SEMICOLON` (semicolon `;`)
     - `EOF` (end of file marker)

**Example**:

```go
lexer := lexer.New("let x = 5 + 5;")
tokens := []token.Token{}
for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
    tokens = append(tokens, tok)
}
```

### 3. Parsing

**Component**: `Parser` struct and related methods in the `parser` package.

**Process**:

1. **Initialize Parser**:
   - Create a new `Parser` instance using the tokens generated by the lexer: `parser.New(lexer)`.
   - Register prefix and infix parsing functions to parse different types of expressions and operators.

2. **Construct AST**:
   - Call the `ParseProgram` method to begin parsing the entire program.
   - The `parseStatement` method parses different types of statements based on the current token type.
   - For `let x = 5 + 5;`, the parser constructs the following AST structure:

```bash
Program
 └── LetStatement
      ├── Name: Identifier(x)
      └── Value: InfixExpression
           ├── Left: IntegerLiteral(5)
           ├── Operator: +
           └── Right: IntegerLiteral(5)
```

**Example**:

```go
parser := parser.New(lexer)
ast := parser.ParseProgram()
if len(parser.Errors()) != 0 {
    // Handle parsing errors
}
```

### 4. Evaluation

**Components**: `Evaluator` in the `evaluator` package and the object system in the `object` package.

**Process**:

1. **Initialize Environment**:
   - Create a new environment instance to store variables and their values: `env := object.NewEnvironment()`.

2. **Traverse and Evaluate AST**:
   - Call `evaluator.Eval(ast, env)` to start evaluating the AST.
   - The evaluation process is as follows:
     - **Program Node**:
       - Traverse its child node `LetStatement`.
     - **LetStatement Node**:
       - Evaluate the assignment expression `5 + 5`:
         - **InfixExpression Node**:
           - Evaluate the left operand `5`, resulting in an `Integer` object.
           - Evaluate the right operand `5`, resulting in an `Integer` object.
           - Perform the addition operation, resulting in an `Integer` object `10`.
       - Bind the variable `x` to the `Integer` object `10`, stored in the environment.

**Example**:

```go
result := evaluator.Eval(ast, env)
if errObj, ok := result.(*object.Error); ok {
    fmt.Println(errObj.Message)
} else {
    // Handle other types of results
}
```

### 5. Output Results

For a `let` statement, there is no direct output, as it is a variable declaration and assignment operation. However, the variable `x` has been stored in the environment and can be used in subsequent expressions or statements.

**Example**:

```go
// Suppose there is a subsequent expression `x + 10;`
lexer := lexer.New("let x = 5 + 5; x + 10;")
parser := parser.New(lexer)
ast := parser.ParseProgram()
env := object.NewEnvironment()
result := evaluator.Eval(ast, env)
if errObj, ok := result.(*object.Error); ok {
    fmt.Println(errObj.Message)
} else {
    fmt.Println(result.Inspect()) // Output: 20
}
```

### 6. Overall Flow Summary

1. **Input Source Code**:

   `let x = 5 + 5;`

2. **Lexical Analysis**:

   Tokenize into a sequence: `LET`, `IDENT(x)`, `ASSIGN`, `INT(5)`, `PLUS`, `INT(5)`, `SEMICOLON`, `EOF`.

3. **Parsing**:

   Construct AST:
   ```bash
   Program
    └── LetStatement
         ├── Name: Identifier(x)
         └── Value: InfixExpression(+)
              ├── Left: IntegerLiteral(5)
              ├── Operator: +
              └── Right: IntegerLiteral(5)
   ```

4. **Evaluation**:
   - Evaluate the expression `5 + 5`, resulting in `10`.
   - Bind `x` to `10`.

5. **Output**:
   - The `let` statement does not produce direct output, but the variable `x` can be used in subsequent operations.

### 7. Further Expansion

If there are more statements or expressions in the source code, the evaluator will continue to traverse the AST, evaluating each node in sequence and updating the environment as needed. For example:

```go
let x = 5 + 5;
x + 10;
```

When evaluating `x + 10;`, the evaluator retrieves the value of `x` (which is `10`) from the environment, then performs the addition operation, ultimately outputting `20`.

### 8. Error Handling

Throughout the process, if errors are encountered during lexical analysis, parsing, or evaluation (e.g., undefined variables, type mismatches, syntax errors), the corresponding error messages will be recorded and output, ensuring the robustness of the program and ease of debugging.

## Contribution Guidelines

Contributions are welcome! Please feel free to submit pull requests or open issues.

Areas you can contribute to include but are not limited to:

- **Error Handling**: Improve error messages in the lexer, parser, and interpreter.
- **New Features**:
  - Implement boolean types and operations.
  - Add conditional statements (`if-else`).
  - Support function definitions and calls, as well as scope management.
  - Extend the object system to include more data types (strings, arrays, etc.).
- **Optimization**: Enhance the performance of the interpreter.
- **Documentation**: Enrich documentation and add more examples.

## License

This project is licensed under the MIT License. For details, please refer to the [LICENSE](LICENSE) file.
