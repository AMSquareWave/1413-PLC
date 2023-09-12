package main

import "fmt"

func main() {
	var counter int = 50
	var i int

	for i = counter; i > 0; i-- {
		fmt.Println(i)
		if i%2 == 0 {
			i++
			i++
		} else {
			i--
			i--
		}
	}
}
