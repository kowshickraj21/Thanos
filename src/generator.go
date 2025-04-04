package main

import (
	"bytes"
	"fmt"
	"os"
)

var variables = make(map[string]int)
var sp = 0
var buf bytes.Buffer

func Generate(node StatementNode) {
	
	buf.WriteString("section .text \n")
	buf.WriteString("global _start \n")
	buf.WriteString("_start: \n")

	for _, i := range node.children {
		generateStatement(i)
	}

	// buf.WriteString("mov rax, 0X0A3831 ;\n")
	// buf.WriteString("push rax ;\n")
	// buf.WriteString("mov rdx, 3 ;\n")
	// buf.WriteString("call print \n")
	// buf.WriteString("pop rax ;\n")

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

func generateStatement(node Node) {
	if node.Type() == "Assignment" {
		generateAssignment(node.(AssignmentNode))
	}
}

func generateAssignment(node AssignmentNode) {

	var variable string
	var value int
	if node.Variable.Type() == "Declaration" {
		declaration,_ := node.Variable.(DeclarationNode)
		variable = generateDeclaration(declaration)
	} else {
		identifier,_ := node.Variable.(IdentifierNode)
		variable = identifier.Value
	}

	intNode,_ := node.Assignee.(IntNode)
	value = intNode.Value
	
	if node.Value == "+=" {
		buf.WriteString(fmt.Sprintf("mov rax, [rsp+%d] ;\n", variables[variable]*8))
		buf.WriteString(fmt.Sprintf("add rax, %d ;\n", value))
		buf.WriteString(fmt.Sprintf("mov [rsp+%d], rax ;\n", variables[variable]*8))
	} else {
		buf.WriteString(fmt.Sprintf("add rax, %d ;\n", value))
		buf.WriteString(fmt.Sprintf("mov [rsp+%d], rax ;\n", variables[variable]*8))
	}
}

func generateDeclaration(node DeclarationNode) string {
	buf.WriteString("mov rax, 0 ;\n")
	buf.WriteString("push rax ;\n")
	sp++
	variables[node.Variable.Value] = sp
	return node.Variable.Value;
}
