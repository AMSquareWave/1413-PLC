package main

import "fmt"

func main() {
	var counter int = 50

	for counter > 0 {
		fmt.Println(counter)
		if counter%2 == 0 {
			counter += 1
		} else {
			counter -= 3
		}
	}
}
