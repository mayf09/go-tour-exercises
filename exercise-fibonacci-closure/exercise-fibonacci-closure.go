package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var i, j int

	return func() int {
		i, j = j, i+j
		if i == 0 {
			j = 1
			return 0
		}
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
