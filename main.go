package main

import (
	"fmt"
	"time"
)

func main() {
	//ProbeerAll([]int{30, 3, 6, 9}, []int{48, 30, 87, 36, 18, 64, 60, 531})
	startTime := time.Now()
	ProbeerFull([]int{30, 3, 6, 9})
	fmt.Println("Duur:", time.Since(startTime))
	//antwoord := Probeer([]int{30, 3, 6, 9}, 531)
	//fmt.Printf("%s = %d\n", antwoord.String(), antwoord.Evaluate())
}

func ProbeerFull(getallen []int) {
	var nodes []Node
	for i := 0; i < len(getallen); i++ {
		nodes = append(nodes, L(getallen[i]))
	}

	antwoorden := ProbeerN(nodes)

	maxAnt := 1

	for i := 0; i < len(getallen); i++ {
		maxAnt *= getallen[i]
	}

	for antwoord := 0; antwoord <= maxAnt; antwoord++ {
		for _, a := range antwoorden {
			if a.Evaluate() == antwoord {
				fmt.Printf("%s = %d\n", a.String(), a.Evaluate())
				break
			}
		}
	}
}

func ProbeerAll(getallen []int, antwoorden []int) {
	for _, antwoord := range antwoorden {
		ants := Probeer(getallen, antwoord)
		if len(ants) == 0 {
			fmt.Printf("Niet gevonden voor %d\n", antwoord)
			continue
		}
		for _, ant := range ants {
			fmt.Printf("%s = %d\n", ant.String(), ant.Evaluate())
		}
	}
}

func Probeer(getallen []int, antwoord int) []Node {
	var nodes []Node
	for i := 0; i < len(getallen); i++ {
		nodes = append(nodes, L(getallen[i]))
	}

	var result []Node

	antwoorden := ProbeerN(nodes)
	for _, a := range antwoorden {
		if a.Evaluate() == antwoord {
			result = append(result, a)
		}
	}
	return result
}

func ProbeerN(nodes []Node) []Node {
	if len(nodes) == 1 {
		return nodes
	}

	var result []Node

	for i := 0; i < len(nodes)-1; i++ {
		left := ProbeerN(nodes[0 : i+1])
		for _, l := range left {
			right := ProbeerN(nodes[i+1:])
			for _, r := range right {
				result = append(result, Probeer2(l, r)...)
			}
		}
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
