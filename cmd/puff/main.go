package main

import (
	"os"
	"github.com/outroninja/pufftmp/pkg/compiler/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./tmp/main.puff")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}