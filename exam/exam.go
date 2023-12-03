package main

import (
	"fmt"
)

func findWinner(n int, cut int) {
	var ppl []int
	for i := 1; i <= n; i++ {
		ppl = append(ppl, i)
	}

	curr := 0
	pplCount := n
	for pplCount > 1 {
		curr = (curr + cut - 1)
		if curr >= pplCount {
			curr %= pplCount
		}
		if curr == 0 {
			ppl = ppl[1:]
		} else if curr == len(ppl) {
			ppl = ppl[:curr-1]
		} else {
			restcopy := ppl[curr+1:]
			ppl = ppl[:curr]
			ppl = append(ppl, restcopy...)
		}

		pplCount--
	}

	fmt.Print(ppl[0])
}

func main() {
	findWinner(8, 3)
	findWinner(11, 5)
}
