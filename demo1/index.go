package main

import "fmt"

func main() {
	fmt.Print("Helloworld\n")
	var tmp1 string = "Hello" // Explicit Declaration
	var tmp2 int = 0
	tmp3 := 1 // Implicit Declaration
	tmp4 := "Hey"
	fmt.Print(tmp1, "\n")
	fmt.Printf("Tmp2 is %d\n", tmp2)
	fmt.Printf("Tmp3 is %d\n", tmp3)
	fmt.Printf("Tmp4 is %s\n", tmp4)
}
