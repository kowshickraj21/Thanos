package main

import (
	"fmt"
	"os"
)

func main() {
	bytes,_ := os.ReadFile("../examples/1.lang")
	code := string(bytes)
	tokens := Tokenize(code)
	fmt.Println(tokens[0].Type)
}