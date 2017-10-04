package API

import (
	"encoding/json"
	"fmt"
	"golearn/SlackReplica/model"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SSI Slack")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.UserCreate
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Fprintf(w, "Invalid Request")
	}
	response := model.CreateUser(user.Name, user.Channels)
	fmt.Fprintf(w, "User: ", response.Name, "Created Successfully")
}
