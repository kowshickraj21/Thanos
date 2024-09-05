package main

import (
	"fmt"
	"os"
)

func main() {
	bytes,_ := os.ReadFile("../examples/1.lang")
	code := string(bytes)
	fmt.Println(code)
}