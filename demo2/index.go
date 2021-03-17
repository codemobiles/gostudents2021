package main

import "fmt"

type user struct {
	username string
	password string
}

// product {name, price, qty}

func main() {
	fmt.Printf("Helloworld")

	u := user{username: "admin", password: "1234"}
	_ = u
	fmt.Println(u)
	fmt.Println(u.username)
}
