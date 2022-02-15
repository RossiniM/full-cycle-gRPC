package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func user(w http.ResponseWriter, req *http.Request) {
	var user User

	json.NewDecoder(req.Body).Decode(&user)

	fmt.Println(user)
}

type User struct {
	Id    int8
	Name  string
	Email string
}

func main() {
	http.HandleFunc("/user", user)
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello\n")
	})
	http.ListenAndServe(":8090", nil)

}
