package main

import (
	"unicode"
)

func add(stack *[]int, number int) error {
	*stack = append(*stack, number)
	return nil
}

func drop(stack *[]int) int {
	tmp := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return tmp
}

func dc(in string) int {
	var stack []int = make([]int, 0)
	var tmp int
	var inNum bool = false

	for _, c := range in {
		//fmt.Println(string(c), stack)
		if unicode.IsDigit(c) && !inNum {
			tmp = int(c) - 48
			inNum = true
		} else if unicode.IsDigit(c) && inNum {
			tmp *= 10
			tmp += int(c) - 48
		} else {
			if inNum {
				add(&stack, tmp)
				inNum = false
			}
			switch c {
			case '+':
				a := drop(&stack)
				b := drop(&stack)
				add(&stack, b+a)
			case '-':
				a := drop(&stack)
				b := drop(&stack)
				add(&stack, b-a)
			case '*':
				a := drop(&stack)
				b := drop(&stack)
				add(&stack, b*a)
			case '/':
				a := drop(&stack)
				b := drop(&stack)
				add(&stack, b/a)
			default:
				continue
			}
		}
	}
	return drop(&stack)
}
