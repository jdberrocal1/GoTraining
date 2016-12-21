package main

import "fmt"

//how declare a struct
type User struct {
        age int
        name, lastname, username string
    }

//this is the way how to add methods to a struct
// you can add methods to structs that are declare in the same file
func (this *User) info () string{
    return this.name + " ("+ this.username +") " + this.lastname
}

func main() {
    user := User {age:24,name:"Daniel", lastname:"Berrocal", username:"jdberrocal1"}
    user2 := new(User) // this return a pointer
    user2.name="test"
    fmt.Println(user.info())
    fmt.Println(*user2)// * is to get value from pointer
    
}