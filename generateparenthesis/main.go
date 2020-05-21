package main

import "fmt"

func main() {
	var s []string
	fmt.Printf("%v", backtrack(s, "", 0, 0, 3))
}

func backtrack(s []string, cur string, open, closes, max int) []string {
	if len(cur) == max*2 {
		return append(s, cur)
	}
	if open < max {
		s = backtrack(s, cur+"(", open+1, closes, max)
		fmt.Printf("sub open[%d] < max[%d]: %v \n", open, max, s)
	}
	if closes < open {
		s = backtrack(s, cur+")", open, closes+1, max)
		fmt.Printf("sub close[%d] < open[%d]: %v \n", closes, open, s)
	}
	return s
}
