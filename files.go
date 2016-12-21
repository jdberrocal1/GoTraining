package main

import (
	"fmt"
    "io/ioutil"
	"os"
	"bufio"
)

func openFile(filepath string) string {
    fileData, err := ioutil.ReadFile(filepath)
    if err != nil {
        panic(err) //return stack of error
    } else {
        return string(fileData)
    }
}

func readFile(filepath string) string {
    file, err := os.Open(filepath)
    defer file.Close() //defer: eject the code at the end of the function at return moment
    if err != nil {
        panic(err)
    } else {
        scanner := bufio.NewScanner(file)
        textFromFile := ""
        for scanner.Scan() {
            line := scanner.Text()
            textFromFile += "\n"+line
        }
        return textFromFile
    }
    
}

func main()  {
    fileData := openFile("./test.txt")
    fmt.Println(fileData)
    fileText := readFile("./test.txt")
    fmt.Println(fileText)
}