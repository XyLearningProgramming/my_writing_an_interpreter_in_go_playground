package my_repl

import (
	"fmt"
	"io"
	lexer "monkey/my_lexer"
	"os"

	evaluator "monkey/my_evaluator"
	parser "monkey/my_parser"

	// "monkey/evaluator"
	// "monkey/parser"
	object "monkey/my_object"

	console "github.com/xingshuo/console/src"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	env := object.NewEnvironment()
	con := console.NewConsole()
	con.SetKeyDownHook('\x03', func(c *console.Console, s string) {
		fmt.Fprintln(out, "keyboard interupt.")
		os.Exit(0)
	})
	con.Init(func(c *console.Console, line string) {
		switch line {
		case "exit":
			fallthrough
		case "exit()":
			fmt.Fprintln(out)
			os.Exit(0)
		default:
			l := lexer.New(line)
			p := parser.New(l)

			program := p.Parse()
			if p.Error() != nil {
				printParserErrors(out, []string{p.Error().Error()})
				return
			}
			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				io.WriteString(out, "\n")
				io.WriteString(out, evaluated.String())
			}
			return
		}
	})
	con.LoopCmd()
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
