package main

import "fmt"

type IUser interface {
    GetName() string
}

type User struct {
    name string
}

func (this *User) GetName() string  {
    return this.name
}

func main() {
    user1 := User{"Daniel"}
    fmt.Println(user1.GetName())
}