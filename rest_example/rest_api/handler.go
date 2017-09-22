package rest_api

import (
	"database/sql"
	"fmt"
	"golearn/rest_example/lib/database"
	_ "io"
	_ "io/ioutil"
	"net/http"
)

var DB *sql.DB

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome To Mujtaba REST Service!")
}

func Futures(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Writing symbols to browser")
	results, err := database.GetFutures(DB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Writing symbols to browser")
	database.WriteHtmlFutures(w, results)
}
