package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/c-bata/go-prompt"
	parser "github.com/li-zeyuan/antlr-example/visitor-parser"
)

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("calc"),
	)

	p.Run()
}

func executor(in string) {
	if len(in) == 0 {
		return
	}

	fmt.Printf("Answer: %d\n", calc(in))
}

func completer(in prompt.Document) []prompt.Suggest {
	var ret []prompt.Suggest
	return ret
}

func calc(input string) int {

	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(tokens)

	v := NewVisitor()
	//Start is rule name of Calc.g4
	p.Start_().Accept(v)
	return v.pop()
}
