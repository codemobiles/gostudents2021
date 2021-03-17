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

	// Call functions
	fn1()
	fnNumber2(2)
	fn3(1, 3)

	// Collection datatype
	var dictOfWords map[string]string = map[string]string{"dog": "หมา", "cat": "แมว"}
	dictOfRate := map[string]float32{"THB": 30.12, "JP": 14.2}
	fmt.Printf("%v\n", dictOfWords)
	fmt.Printf("Dog is %s\n", dictOfWords["dog"])
	fmt.Printf("THB is %.2f\n", dictOfRate["THB"])


	// Condition
	if tmp1 == "Hellod" {
		fmt.Printf("tmp1 is Hello")
	}else{
		fmt.Printf("tmp1 is NOT Hello")
	}
}

func fn1() {
	fmt.Println("Fn1")
}

// camel - fnNumber
// Pacal - FnNumber
// snake - fn_number
func fnNumber2(count int) {
	fmt.Printf("Count: %d\n", count)
}

func fn3(a int, b int) int {
	return a + b
}
