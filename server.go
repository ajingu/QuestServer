package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"string"`
	IsPaid bool `json:"isPaid"`
}

func handleRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello!")

	return
}

func handleGetAll(w http.ResponseWriter, r *http.Request){
	users, err := retrieveAll()
	if err != nil{
		fmt.Fprintln(w, err.Error())
	}

    output, err := json.MarshalIndent(users, "", "\t")
    if err != nil{
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.Write(output)

	return
}

func main(){
	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/users/all", handleGetAll)

    server.ListenAndServe()
}