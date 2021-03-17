package main

import "fmt"

func main() {
	fmt.Print("Helloworld\n")
	var tmp1 string = "Hello" // Explicit Declaration
	var tmp2 int = 0
	tmp3 := 1 // Implicit Declaration
	tmp4 := "Hey"
	tmp4 = "CodeMobiles"

	const tmp5 string = "lek"
	// tmp5 = "home"

	var tmp6 []int = []int{1, 2, 3, 4}
	var tmp7 []string = []string{"Angular", "React", "Golang"}

	fmt.Print(tmp1, "\n")
	fmt.Printf("Tmp2 is %d\n", tmp2)
	fmt.Printf("Tmp3 is %d\n", tmp3)
	fmt.Printf("Tmp4 is %s\n", tmp4)
	fmt.Printf("Tmp5 is %s\n", tmp5)
	fmt.Println("Lek")
	fmt.Printf("ค่ามันคือ %v\n", tmp6)
	fmt.Printf("หลักสูตรทั้งหมด %v\n", tmp7)
}
