package main

import (
	"net/http"
)

func startServer()  {

    http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request)  {
        http.ServeFile(w, r, "index.html")
    })

    publicFolder := http.FileServer(http.Dir("public"))
    http.Handle("/public/", http.StripPrefix("/public", publicFolder))

    http.ListenAndServe(":8000", nil)
}

func main() {
    startServer()
}