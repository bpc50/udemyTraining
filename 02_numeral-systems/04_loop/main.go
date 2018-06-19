package main

import "fmt"

func main() {
	for i := 5; i < 50000; i=i+5 {	//instead of i++, i=i+1 means the same thing.
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
