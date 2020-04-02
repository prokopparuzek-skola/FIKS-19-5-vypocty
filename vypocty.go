package main

import "fmt"

const (
	ZPRAVA = true
	ZLEVA  = false
)

type Operator_t struct {
	normal   rune
	priority int
}

func main() {
	var N int

	fmt.Scan(&N)
	for ; N > 0; N-- {
		var P, A int
		var operatory map[rune]Operator_t = make(map[rune]Operator_t)
		var asociativity []bool = make([]bool, A)
		var in string

		fmt.Scan(&P, &A)
		for ; P > 0; P-- {
			var s, t rune
			var K int
			var tmp Operator_t

			fmt.Scan(&s, &t, &K)
			tmp.normal = t
			tmp.priority = K
			operatory[s] = tmp
		}
		for i := 0; i < A; i++ {
			var tmp rune

			fmt.Scan(&tmp)
			switch tmp {
			case 'P':
				asociativity[i] = ZPRAVA
			case 'L':
				asociativity[i] = ZLEVA
			}
		}
		fmt.Scanln(&in)
	}
}
