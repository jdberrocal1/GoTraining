package main

import (
    "fmt"
    "strings"
	"time"
)

func SlowlyName(name string)  {
    chars := strings.Split(name, "")

    for _, char := range(chars){
        time.Sleep(1000 * time.Millisecond)
        fmt.Println(char)
    }
}

func main() {
    go SlowlyName("Daniel")
    fmt.Println("Synchronous")
        var wait string
        fmt.Scanln(&wait) 
}