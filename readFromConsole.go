package main

import (
    "fmt"
)

func main()  {
    var name string
    fmt.Println("Type your name: ")
    fmt.Scanf("%v",&name)
    fmt.Println("Welcome " + name)
}