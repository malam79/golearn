package rest_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golearn/rest_example/lib/database"
	_ "io"
	_ "io/ioutil"
	"net/http"
)

var DB *sql.DB

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SSI")
}

func Symbols(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	var bytes []byte
	if string(q.Get("type")) == "equity" {
		results, err := database.GetEquity(DB)
		if err != nil {
			fmt.Println(err)
		}
		bytes, err = json.Marshal(results)
	} else if string(q.Get("type")) == "future" {
		results, err := database.GetFutures(DB)
		if err != nil {
			fmt.Println(err)
		}
		bytes, _ = json.Marshal(results)
	} else {
		bytes, _ = json.Marshal("Support for others is not available yet")
	}
	fmt.Fprintf(w, string(bytes))
}
