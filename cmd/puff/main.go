package main

import (
	"os"
	"github.com/outroninja/pufftmp/pkg/compiler/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./tmp/main.puff")
	source := string(bytes)

}