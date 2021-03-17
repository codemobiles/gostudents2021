package main

import (
	"fmt"
	"strings"
)

type user struct {
	username string
	password string
}

func (self *user) empty() {
	self.username = ""
	self.password = ""
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

	u1 := user{username: "admin", password: "1234"}
	u1.empty()
	fmt.Printf("\nu1 is %v\n", u1)

}
