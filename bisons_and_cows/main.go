package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//parse -> string(a) , uint(a), []rune(a)

func play() {
	cows := 0
	bisons := 0
	reader := bufio.NewReader(os.Stdin)
	t := rand.Intn(9999) + 1001
	if t > 9999 {
		t = 1234
	}
	target := strconv.Itoa(t)
	for {
		guess, _ := reader.ReadString('\n')
		guess = guess[:4]
		if strings.Compare(guess, target) == 0 {
			fmt.Printf("good job you found %v", target)
			return
		}
		for i := 0; i < 4; i++ {
			if guess[i] == target[i] {
				bisons++
			} else {
				for j := 0; j < 4; j++ {
					if guess[i] == target[j] && i != j {
						cows++
					}
				}
			}
		}
		if bisons == 4 {
			fmt.Printf("good job you found %v", target)
			return
		}
		fmt.Printf("cows: %v , bisons: %v \n", cows, bisons)
		bisons = 0
		cows = 0
	}
}

func main() {
	play()
}
