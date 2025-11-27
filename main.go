package main

import (
	"fmt"
	"net/http"
	"Steam-API/routeur"
)

func main () {
	r:= routeur.New()

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", r)
}