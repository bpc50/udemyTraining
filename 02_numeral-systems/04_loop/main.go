package main

import "fmt"

func main() {
	for i := 65; i < 123; i=i+1 {	//instead of i++, i=i+1 means the same thing.
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
