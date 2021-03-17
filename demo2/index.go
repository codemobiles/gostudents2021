package main

import (
	"fmt"
	"strings"
)

type user struct {
	username string
	password string
}

func (_ user) clear() user {
	return user{username: "", password: ""}
}

func (clone user) toUpper() user {
	return user{
		username: strings.ToUpper(clone.username),
		password: strings.ToUpper(clone.password),
	}
}

// product {name, price, qty}

func main() {
	fmt.Printf("Helloworld")

	u := user{username: "admin", password: "1234"}
	_ = u
	fmt.Println(u)
	fmt.Println(u.username)
	// u = u.clear()
	fmt.Println("Username : " + u.username)

	u = u.toUpper()
	fmt.Printf("%v", u)
}
