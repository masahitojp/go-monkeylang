package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vishen/go-monkeylang/eval"
	"github.com/vishen/go-monkeylang/lexer"
	"github.com/vishen/go-monkeylang/object"
	"github.com/vishen/go-monkeylang/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// Keep the environment around so we can use variables
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, "[DEBUG] ")
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		evaluated := eval.Eval(program, env)
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
