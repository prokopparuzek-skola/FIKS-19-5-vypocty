package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const (
	ZPRAVA = true
	ZLEVA  = false
)

type Operator_t struct {
	normal   rune
	priority int
}

func rozklad(in string, operatory map[rune]Operator_t, asociativity []bool) (out string) {
	var queue []string = make([]string, 0)
	var stack []rune = make([]rune, 0)
	var inNum bool = false
	var tmp []rune = make([]rune, 0)

	for _, c := range in {
		if unicode.IsNumber(c) && !inNum {
			tmp = append(tmp, c)
			inNum = true
		} else if unicode.IsDigit(c) && inNum {
			tmp = append(tmp, c)
		} else {
			if inNum {
				queue = append(queue, string(tmp))
				tmp = make([]rune, 0)
				inNum = false
			}
			switch c {
			case '(':
				stack = append(stack, '(')
			case ')':
				//fmt.Println(stack)
				for stack[len(stack)-1] != '(' {
					queue = append(queue, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = stack[:len(stack)-1]
			case ' ':
				continue
			default:
				//fmt.Println(c, stack)
				for o, data := range operatory {
					if c == o {
						if len(stack) > 0 {
							for ((asociativity[data.priority-1] == ZLEVA && data.priority >= operatory[stack[len(stack)-1]].priority) || (asociativity[data.priority-1] == ZPRAVA && data.priority > operatory[stack[len(stack)-1]].priority)) && stack[len(stack)-1] != '(' {
								queue = append(queue, string(stack[len(stack)-1]))
								stack = stack[:len(stack)-1]
								if len(stack) == 0 {
									break
								}
							}
						}
						stack = append(stack, c)
					}
				}
			}
		}
		//fmt.Println(stack, queue)
	}
	for i := len(stack) - 1; i >= 0; i-- {
		queue = append(queue, string(stack[i]))
	}
	for _, s := range queue {
		if unicode.IsNumber(rune(s[0])) {
			out += s + " "
		} else {
			out += string(operatory[rune(s[0])].normal)
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int

	//fmt.Scan(&N)
	fmt.Fscanf(reader, "%d\n", &N)
	//fmt.Println(N)
	for ; N > 0; N-- {
		var P, A int
		var operatory map[rune]Operator_t = make(map[rune]Operator_t)
		var asociativity []bool
		var in string
		var tmpS string

		//fmt.Scan(&P, &A)
		fmt.Fscanf(reader, "%d %d\n", &P, &A)
		//fmt.Println(P, A)
		for ; P > 0; P-- {
			var s, t string
			var K int
			var tmp Operator_t

			//fmt.Scan(&s, &t, &K)
			fmt.Fscanf(reader, "%s %s %d\n", &s, &t, &K)
			//fmt.Println(s, t, K)
			tmp.normal = rune(t[0])
			tmp.priority = K
			operatory[rune(s[0])] = tmp
		}
		asociativity = make([]bool, A)
		for i := 0; i < A; i++ {
			//fmt.Scan(&tmpS)
			fmt.Fscanf(reader, "%s ", &tmpS)
			switch rune(tmpS[0]) {
			case 'R':
				asociativity[i] = ZPRAVA
			case 'L':
				asociativity[i] = ZLEVA
			}
		}
		//fmt.Println(operatory)
		//fmt.Println(asociativity)
		in, _ = reader.ReadString('\n')
		//fmt.Println(in)
		rpn := rozklad(in, operatory, asociativity)
		//fmt.Println(rpn)
		fmt.Println(dc(rpn))
	}
}
