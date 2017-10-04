package main

import (
	"fmt"
	"golearn/rest_example/lib/database"
	"golearn/rest_example/rest_api"
	"net/http"
)

func main() {
	rest_api.DB = database.ConnectMongoDB()
	var dbManager database.MongoDB
	defer database.CloseDataBase(dbManager, rest_api.DB)
	router := rest_api.NewRouter()
	fmt.Println(http.ListenAndServe(":8080", router))
}
