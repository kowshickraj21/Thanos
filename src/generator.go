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

	buf.WriteString("mov rax, 0X0A3831 ;\n")
	buf.WriteString("push rax ;\n")
	buf.WriteString("mov rdx, 3 ;\n")
	buf.WriteString("call print \n")
	buf.WriteString("pop rax ;\n")

	buf.WriteString("mov rax, 60 ;\n")
	buf.WriteString("mov rdi, 20 ;\n")
	buf.WriteString("syscall ;\n")

	buf.WriteString("print: \n")
	buf.WriteString("mov rax, 1 ;\n")
	buf.WriteString("mov rdi, 1 ;\n")
	buf.WriteString("lea rsi, [rsp+8] ;\n")
	buf.WriteString("syscall ;\n")
	buf.WriteString("ret ;\n")

	file, _ := os.Create("out.asm")
	defer file.Close()

	file.Write(buf.Bytes())
}
