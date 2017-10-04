package model

import (
	_ "fmt"
)

var Users = make(map[string]User)

var idGenerator = 0

type UserCreate struct {
	Name     string   `json: "Username"`
	Channels []string `json: "Channels"`
}

type User struct {
	ID       int        `json: "UserID"`
	Name     string     `json: "UserName"`
	Channels []*Channel `json: "Channels"`
}

func CreateUser(name string, subsChannel []string) *User {
	_, ok := Users[name]
	if ok == true {
		return nil
	}
	idGenerator++
	newUser := User{idGenerator, name, make([]*Channel, 10)}
	Users[name] = newUser
	for _, chName := range subsChannel {
		AddUserToChannel(chName, &newUser)
	}
	return &newUser
}

func FindUser(name string) *User {
	val, ok := Users[name]
	if ok != true {
		return &val
	}
	return nil
}

func AppendChannel(uname, chName string) {
	user := FindUser(uname)
	if user != nil {
		AddUserToChannel(chName, user)
	}
}
