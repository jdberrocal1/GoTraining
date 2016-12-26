package main

import (
	"net/http"
	"encoding/json"
    "./entities/user"
)

func startServer()  {

    http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request)  {
        users := user.Users{
            user.User{"Daniel Berrocal", "jjdanielb", "jjdanielb@gmail.com"},
            user.User{"Jose Berrocal", "jdbr2010", "jdbr2010@gmail.com"},
            user.User{"Mario Rodriguez", "maro01", "maro01@gmail.com"},
        }
        json.NewEncoder(w).Encode(users)
    })

    http.ListenAndServe(":8000", nil)
}


func main() {
    startServer()
}