package rest_api

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"golearn/rest_example/lib/database"
	"gopkg.in/mgo.v2"
	_ "io"
	_ "io/ioutil"
	"net/http"
)

var DB *mgo.Session

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SSI")
}

func Symbols(w http.ResponseWriter, req *http.Request) {
	var dbManager database.MongoDB
	q := req.URL.Query()
	var bytes []byte
	if string(q.Get("type")) == "equity" {
		results, err := database.GetEquities(dbManager, DB)
		if err != nil {
			fmt.Println(err)
		}
		bytes, err = json.Marshal(results)
	} else if string(q.Get("type")) == "future" {
		results, err := database.GetFuture(dbManager, DB)
		if err != nil {
			fmt.Println(err)
		}
		bytes, _ = json.Marshal(results)
	} else {
		bytes, _ = json.Marshal("Support for others is not available yet")
	}
	fmt.Fprintf(w, string(bytes))
}
