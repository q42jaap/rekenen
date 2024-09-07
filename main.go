package main

import "fmt"

func main() {
	antwoord := Probeer([]int{600, 40, 20, 8}, 610)
	fmt.Printf("%s = %d\n", antwoord.String(), antwoord.Evaluate())
}

func Probeer(getallen []int, antwoord int) Node {
	var nodes []Node
	for i := 0; i < len(getallen); i++ {
		nodes = append(nodes, L(getallen[i]))
	}

	antwoorden := ProbeerN(nodes)
	for _, a := range antwoorden {
		if a.Evaluate() == antwoord {
			return a
		}
	}
	panic("niet gevonden")
}

func ProbeerN(nodes []Node) []Node {
	if len(nodes) == 1 {
		return nodes
	}
	if len(nodes) == 2 {
		return Probeer2(nodes[0], nodes[1])
	}

	var result []Node
	// A: combine first two nodes
	for _, a := range Probeer2(nodes[0], nodes[1]) {
		// B: combine result of A with the rest of the nodes
		for _, b := range ProbeerN(nodes[2:]) {
			result = append(result, Probeer2(a, b)...)
		}
	}

	// C: don't combine the first two, but keep the first node standalone
	for _, c := range ProbeerN(nodes[1:]) {
		result = append(result, Probeer2(nodes[0], c)...)
	}

	return result
}

func Probeer2(a, b Node) []Node {
	var result []Node

	// (a + b)
	result = append(result, BinaryOperation{
		left:     a,
		operator: OperatorAdd,
		right:    b,
	})

	// (a - b)
	result = append(result, BinaryOperation{
		left:     a,
		operator: OperatorSubtract,
		right:    b,
	})

	// (a * b)
	result = append(result, BinaryOperation{
		left:     a,
		operator: OperatorMultiply,
		right:    b,
	})

	// (a / b)
	aValue := a.Evaluate()
	bValue := b.Evaluate()
	if bValue != 0 && aValue >= bValue && aValue%bValue == 0 {
		result = append(result, BinaryOperation{
			left:     a,
			operator: OperatorDivide,
			right:    b,
		})
	}

	return result
}

type Node interface {
	Evaluate() int
	String() string
}

type Literal int

func L(value int) Literal {
	return Literal(value)
}

func (l Literal) Evaluate() int {
	return int(l)
}

func (l Literal) String() string {
	return fmt.Sprintf("%d", l)
}

type BinaryOperation struct {
	left     Node
	operator Operator
	right    Node
}

func (b BinaryOperation) Evaluate() int {
	switch b.operator {
	case OperatorAdd:
		return b.left.Evaluate() + b.right.Evaluate()
	case OperatorSubtract:
		return b.left.Evaluate() - b.right.Evaluate()
	case OperatorMultiply:
		return b.left.Evaluate() * b.right.Evaluate()
	case OperatorDivide:
		return b.left.Evaluate() / b.right.Evaluate()
	}
	panic("unknown operator")
}

func (b BinaryOperation) String() string {
	return fmt.Sprintf("(%s %s %s)", b.left.String(), b.operator, b.right.String())
}

type Operator string

const (
	OperatorAdd      Operator = "+"
	OperatorSubtract          = "-"
	OperatorMultiply          = "*"
	OperatorDivide            = "/"
)
