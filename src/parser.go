package main

type Node interface {
	Type() string
}
type StatementNode struct {
	children []Node
}

func (StatementNode) Type() string {
	return "Statement"
}

type KeywordNode struct {
	Value string
}

func (KeywordNode) Type() string {
	return "keyword"
}

type AssignemtNode struct {
	Value    string
	Variable Node
	Assignee Node
}

func (AssignemtNode) Type() string {
	return "Assignment"
}

type IdentifierNode struct {
	Value string
}

func (IdentifierNode) Type() string {
	return "Identifier"
}

type CompareNode struct {
	Value string
	com1  Node
	com2  Node
}

func (CompareNode) Type() string {
	return "Comparator"
}

type ArithmeticNode struct {
	Value string
	Var1  Node
	Var2  Node
}

func (ArithmeticNode) Type() string {
	return "Arithmetic"
}

type DeclarationNode struct {
	DataType string
	Variable Node
}

func (DeclarationNode) Type() string {
	return "Declaration"
}

type SymbolNode struct {
	Value string
}

func (SymbolNode) Type() string {
	return "Symbol"
}

type IntNode struct {
	Value string
}

func (IntNode) Type() string {
	return "Integer"
}

func Parse(tokens []Token) (StatementNode, error) {
	var program StatementNode
	i := 0
	for i < len(tokens) {
		i = parseStatement(tokens, i, &program)
	}
	return program, nil
}

func parseStatement(tokens []Token, i int, node *StatementNode) int {

	if tokens[i].Type == "DataType" && tokens[i+1].Type == "Identifier" {
		var declaration DeclarationNode
		declaration.DataType = tokens[i].Value
		var id IdentifierNode
		id.Value = tokens[i+1].Value
		declaration.Variable = id
		i += 2

		if tokens[i].Value == "=" {
			i++
			var assignment AssignemtNode
			assignment.Value = "="
			assignment.Variable = declaration
			if tokens[i+1].Value != ";" {
				assignment.Assignee, _ = parseArithmetic(tokens, i)
			} else {
				var int_lit IntNode
				int_lit.Value = tokens[i].Value
				assignment.Assignee = int_lit
				i++
			}
			node.children = append(node.children, assignment)
			return i + 1
		}
	} else if tokens[i].Type == "Identifier" && tokens[i+1].Type == "Assignment" {
		var assignment AssignemtNode
		assignment.Value = tokens[i+1].Value
		var id IdentifierNode
		id.Value = tokens[i].Value
		assignment.Variable = id

		i += 2
		if tokens[i+1].Value != ";" {
			assignment.Assignee, _ = parseArithmetic(tokens, i)
		} else {
			var int_lit IntNode
			int_lit.Value = tokens[i].Value
			assignment.Assignee = int_lit
			i++
		}
		node.children = append(node.children, assignment)
	}
	return i + 1
}

func parseArithmetic(tokens []Token, i int) (Node, error) {
	if tokens[i+1].Value == ";" {
		var num IntNode
		num.Value = tokens[i].Value
		i += 2
		return num, nil
	}
	var node ArithmeticNode
	var num IntNode
	num.Value = tokens[i].Value
	node.Value = tokens[i+1].Value
	node.Var1 = num
	i += 2
	node.Var2, _ = parseArithmetic(tokens, i)
	return node, nil
}

// func PrintParseTree(node Node) {
// 	fmt.Println(node);
// 	printNodes(node,2)
// }

// func printNodes(node Node,space int){
// 	for i := 0; i < space; i++ {
// 		fmt.Print(" ")
// 	}
// 	fmt.Println(node);
// 	switch node.Type() {
// 	case "Assignment":
// 		KeywordNode(node).Variable
// 	}
// }