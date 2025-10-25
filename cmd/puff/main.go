package main

import (
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./tmp/main.puff")
	source := string(bytes)

}