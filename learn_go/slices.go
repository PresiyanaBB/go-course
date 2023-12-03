package main

import (
	"strings"
)

// func main() {
// 	a1 := [...]int{1, 2}
// 	a2 := a1
// 	a2[0] = 3
// 	fmt.Printf("%v , %v , %t \n", a1, a2, a1 == a2)
// 	//
// 	sourceMap := make(map[any]any)
// 	sourceMap["one"] = 1
// 	sourceMap["two"] = 2
// 	sourceMap["three"] = 3
// 	sourceMap["four"] = 4
// 	sourceMap["five"] = 5
// 	destMap := CopyMap(sourceMap)
// 	fmt.Printf("%v %p %p\n", destMap, &sourceMap, &destMap)
// 	//
// 	smth := make(map[string]int, 10)
// 	smth["1"] = 1
// 	//

// }

func CopyMap(m map[any]any) map[any]any {
	result := make(map[any]any, len(m))
	for key, val := range m {
		result[key] = val
	}
	return result
}

func wordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		counts[word]++
	}
	return counts
}
