package main

import "fmt"

func main() {
	//	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	// %d = decimal   %b = binary   %x = hexadecimal
	// \n = next line command
	// the three 42's are progressively applied to each, so, the first 42 is assigned to the %d, the second to %b, etc

	//fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	//fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
