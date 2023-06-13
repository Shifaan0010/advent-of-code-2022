package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Direction uint8

const (
	Left Direction = iota
	Right
)

func (dir Direction) String() string {
	if dir == Left {
		return "L"
	} else if dir == Right {
		return "R"
	} else {
		return "X"
	}
}

type Monkey interface {
	Yell() int
	Find(num *Number) []Direction
	SolveFor(numDirections []Direction, currVal int) int
}

type Number struct {
	val int
}

func (number Number) Yell() int {
	return number.val
}

func (number *Number) Find(num *Number) []Direction {
	if number == num {
		return []Direction{}
	} else {
		return nil
	}
}

func (number Number) SolveFor(numDirections []Direction, currVal int) int {
	return currVal
}

func (number Number) String() string {
	return fmt.Sprint(number.val)
}

type Expr struct {
	op          byte
	left, right Monkey
}

func (expr Expr) Yell() int {
	if expr.op == '*' {
		return expr.left.Yell() * expr.right.Yell()
	} else if expr.op == '/' {
		return expr.left.Yell() / expr.right.Yell()
	} else if expr.op == '+' {
		return expr.left.Yell() + expr.right.Yell()
	} else if expr.op == '-' {
		return expr.left.Yell() - expr.right.Yell()
	} else {
		panic("Invalid operator")
	}
}

func (expr *Expr) Find(num *Number) []Direction {
	if path := expr.left.Find(num); path != nil {
		return append(path, Left)
	} else if path := expr.right.Find(num); path != nil {
		return append(path, Right)
	} else {
		return nil
	}
}

func (expr Expr) SolveFor(numDirections []Direction, currVal int) int {
	if numDirections == nil || len(numDirections) == 0 {
		panic("Invalid Directions")
	}

	dir := numDirections[len(numDirections) - 1]

	if dir == Left {
		rightVal := expr.right.Yell()
		if expr.op == '=' {
			currVal = rightVal
		} else if expr.op == '*' {
			currVal = currVal / rightVal
		} else if expr.op == '/' {
			currVal = currVal * rightVal
		} else if expr.op == '+' {
			currVal = currVal - rightVal
		} else if expr.op == '-' {
			currVal = currVal + rightVal
		} else {
			panic("Invalid Op")
		}

		return expr.left.SolveFor(numDirections[:len(numDirections)-1], currVal)
	} else if dir == Right {
		leftVal := expr.left.Yell()
		if expr.op == '=' {
			currVal = leftVal
		} else if expr.op == '*' {
			currVal = currVal / leftVal
		} else if expr.op == '/' {
			currVal = leftVal / currVal
		} else if expr.op == '+' {
			currVal = currVal - leftVal
		} else if expr.op == '-' {
			currVal = leftVal - currVal
		} else {
			panic("Invalid Op")
		}

		return expr.right.SolveFor(numDirections[:len(numDirections)-1], currVal)
	} else {
		panic("Invalid Directions")
	}
}

func (expr Expr) String() string {
	return fmt.Sprintf("(%s %c %s)", expr.left, expr.op, expr.right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numRegex := regexp.MustCompile(`(\w+): (\d+)`)
	opRegex := regexp.MustCompile(`(\w+): (\w+) ([*/+-]) (\w+)`)

	monkeys := map[string]Monkey{}

	type ExprMatch struct {
		expr  *Expr
		match []string
	}

	exprs := []ExprMatch{}

	for scanner.Scan() {
		line := scanner.Text()

		if matched := numRegex.FindStringSubmatch(line); matched != nil {
			// fmt.Printf("Num: %v\n", matched[1:])

			number, _ := strconv.Atoi(matched[2])

			monkeys[matched[1]] = &Number{val: number}
		} else if matched := opRegex.FindStringSubmatch(line); matched != nil {
			// fmt.Printf("Op: %v\n", matched[1:])

			expr := Expr{op: matched[3][0]}

			monkeys[matched[1]] = &expr

			exprs = append(exprs, ExprMatch{expr: &expr, match: matched})
		}
	}

	for _, exprMatch := range exprs {
		(*exprMatch.expr).left = monkeys[exprMatch.match[2]]
		(*exprMatch.expr).right = monkeys[exprMatch.match[4]]

		// fmt.Printf("%p %#v\n", exprMatch.expr, exprMatch.expr)
	}

	fmt.Println("Part 1")
	
	rootVal := monkeys["root"].Yell()
	fmt.Printf("root yells %d\n", rootVal)
	
	fmt.Println("\nPart 2")
	
	root := monkeys["root"].(*Expr)
	root.op = '='

	humn := monkeys["humn"].(*Number)
	
	// fmt.Printf("%#v\n", monkeys["humn"])
	// fmt.Printf("%s\n", monkeys["root"])

	directions := root.Find(humn)
	fmt.Printf("Directions to humn: %v\n", directions)
	fmt.Printf("Value of humn: %d\n", root.SolveFor(directions, 0))
}
