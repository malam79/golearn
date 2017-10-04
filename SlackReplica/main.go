package main

import (
	"fmt"
	"golearn/SlackReplica/API"
	"net/http"
)

func main() {
	router := API.NewRouter()
	fmt.Println(http.ListenAndServe(":8080", router))
}
