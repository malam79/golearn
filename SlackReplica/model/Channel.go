package model

import (
	_ "fmt"
)

var Channels = make(map[string]Channel)

type Channel struct {
	Name  string  `json: "Channel"`
	Users []*User `json: "Users"`
}

func CreateChannel(name string) *Channel {
	_, ok := Channels[name]
	if ok == true {
		return nil
	}
	ch := Channel{name, make([]*User, 10)}
	Channels[name] = ch
	return &ch
}

func FindChannel(name string) *Channel {
	val, ok := Channels[name]
	if ok == true {
		return nil
	}
	return &val
}

func AddUserToChannel(name string, user *User) {
	ch := FindChannel(name)
	if ch != nil {
		ch.Users = append(ch.Users, user)
	}
}
