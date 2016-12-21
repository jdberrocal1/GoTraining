package main

import "fmt"
import "strconv"

func main() {
    var num int
    fmt.Println(num)
    var name = "daniel"
    name,num = "Jose Daniel", 85
    lastName := "Berrocal" // := alternative to initialize vars
    edad := 24
    edadStr := strconv.Itoa(edad) //int to string
    edad,_ = strconv.Atoi(edadStr) //string to int // _ is used to ignore vars

    var ( //multiple var declaration
        var1 = 1
        var2, var3 = 2, "3a"
    )

    fmt.Println(name)
    fmt.Println(name, "Edad: " + edadStr)
    fmt.Println(edad)
    fmt.Println(num)
    fmt.Println(lastName)
    fmt.Println(var1, var2, var3)


}