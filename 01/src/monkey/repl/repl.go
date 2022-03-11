package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
	"os"
	"strconv"
)

const (
	PROMPT                  = ">> "
	ExitCommand CommandType = "exit"
)

type CommandType string

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, _ = fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		cmd, ok := lookupCommand(line)
		if ok {
			execCommand(cmd)
		}
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%s\n", tokenToString(tok))
		}
	}
}

func tokenToString(tok token.Token) []byte {
	if tok.Literal != "" {
		tok.Literal = strconv.QuoteToASCII(tok.Literal)
		tok.Literal = tok.Literal[1 : len(tok.Literal)-1]
	}
	tb, _ := json.Marshal(tok)
	return tb
}

func lookupCommand(line string) (CommandType, bool) {
	if line == string(ExitCommand) {
		return ExitCommand, true
	}
	return "", false
}

func execCommand(cmd CommandType) {
	if cmd == ExitCommand {
		os.Exit(0)
		return
	}
	return
}
