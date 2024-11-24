// pkg/repl/repl.go

package repl

import (
	"bufio"
	"fmt"
	"io"

	"punyGo/pkg/evaluator"
	"punyGo/pkg/lexer"
	"punyGo/pkg/object"
	"punyGo/pkg/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment(nil)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text() // scan a line
		lex := lexer.New(line) // create a lexical analyzer
		par := parser.New(lex) // create a syntactic parser

		// start parsing
		program := par.ParseProgram()

		if len(par.Errors()) != 0 {
			printParserErrors(out, par.Errors())
			continue
		}

		// print AST
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		// evaluate AST
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
