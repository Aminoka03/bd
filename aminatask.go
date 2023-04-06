// Write a program that takes a list of strings as input and returns the longest string in the list.
package main

import (
	"fmt"
	"unicode/utf8"
)

func findLongest(list []string) string {
	longest := ""
	for _, str := range list {
		if len(str) > utf8.RuneCountInString(longest) {
			longest = str
		}
	}
	return longest
}

func main() {
	var n int
	fmt.Scan(&n)
	var stringsList []string
	for i := 0; i < n; i++ {
		var q string
		fmt.Scan(&q)
		stringsList = append(stringsList, q)
	}
	longest := findLongest(stringsList)
	fmt.Print(longest)
}
