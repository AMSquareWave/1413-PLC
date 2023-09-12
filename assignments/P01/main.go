package main

import (
	"fmt"

	"example.com/go-demo-1/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println("The best mascot is " + mascot.BestMascot() + ".")
	fmt.Println("The second best mascot is " + mascot.SecondBestMascot() + ".")
	fmt.Println(quote.Go())
}
