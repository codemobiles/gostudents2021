package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hey\n")
	tmp0 := "Hey1234\n"
	clear(&tmp0)
	fmt.Printf(tmp0)
}

func power(obj *int){
	
}

func clear(obj *string) {
	*obj = "Clear"
}
