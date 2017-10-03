package database

import (
	"fmt"
	"golearn/rest_example/lib"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	_ "log"
)

const data_base_url string = "localhost"

func ConnectMongoDB() *mgo.Session {
	session, err := mgo.Dial(data_base_url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return session
}

type MongoDB struct{}

func (MongoDB) CloseDB(db interface{}) {
	(db.(*mgo.Session)).Close()
}

func (MongoDB) GetFutures(db interface{}) (results []lib.Futures, err error) {
	collection := (db.(*mgo.Session)).DB("Instrument").C("Futures")
	err = collection.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return
}

func (MongoDB) GetEquity(db interface{}) (results []lib.Equity, err error) {
	collection := (db.(*mgo.Session)).DB("Instrument").C("Futures")
	err = collection.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return
}
