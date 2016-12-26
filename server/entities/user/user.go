package user

import (
	"fmt"
)

func init()  {
    fmt.Println("Initializing package")
}

type User struct {
    Name string `json:"name"`
    Username string `json:"username"`
    Email string `json:"email"`
}

type Users []User