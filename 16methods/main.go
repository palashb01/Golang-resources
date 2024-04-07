package main

import "fmt"

func main() {
	emp1 := User{"hitesh", "pbaderia0", true, 22}
	emp1.GetStatus()
	emp1.SetEmail()
}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}

// This is method of the User struct in golang
func (u User) GetStatus() {
	fmt.Println(u.status)
}

func (u User) SetEmail(){
	u.Email = "pbaderia0"
	fmt.Println("Email is set to: ", u.Email)
}