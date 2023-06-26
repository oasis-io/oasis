package main

import (
	"fmt"
	"sort"
)

func main() {
	var testMap = make(map[string]string)

	testMap["B"] = "B"
	testMap["A"] = "A"
	testMap["D"] = "D"
	testMap["C"] = "C"
	testMap["E"] = "E"

	for key, value := range testMap {
		fmt.Println(key, ":", value)
	}

	var testSlice []string
	for k, _ := range testMap {
		testSlice = append(testSlice, k)
	}

	sort.Strings(testSlice)

	for _, key := range testSlice {
		if v, ok := testMap[key]; ok {
			fmt.Println(v)
		}
	}

}
