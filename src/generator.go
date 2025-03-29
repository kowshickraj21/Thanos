package main

import (
	"bytes"
	"os"
)

func Generate(node Node) {
	var buf bytes.Buffer
	buf.WriteString("section .text \n")
	buf.WriteString("global _start \n")
	buf.WriteString("_start: \n")
	buf.WriteString("mov rax, 60 ;\n")
	buf.WriteString("mov rdi, 20 ;\n")
	buf.WriteString("syscall ;\n")
	file, _ := os.Create("out.asm")
	defer file.Close()

	file.Write(buf.Bytes())
}
