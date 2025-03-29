package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("../examples/1.lang")
	code := string(bytes)
	tokens, err := Tokenize(code)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(tokens)
	node, _ := Parse(tokens)
	// PrintParseTree(node)
	Generate(node)

	fmt.Println(node)
}

