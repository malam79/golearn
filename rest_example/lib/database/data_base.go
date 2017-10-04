package database

import (
	"golearn/rest_example/lib"
)

type DBManager interface {
	GetFutures(interface{}) ([]lib.Futures, error)
	GetEquity(interface{}) ([]lib.Equity, error)
	CloseDB(interface{})
}

func GetFuture(dbManager DBManager, db interface{}) ([]lib.Futures, error) {
	return dbManager.GetFutures(db)
}

func GetEquities(dbManager DBManager, db interface{}) ([]lib.Equity, error) {
	return dbManager.GetEquity(db)
}

func CloseDataBase(dbManager DBManager, db interface{}) {
	dbManager.CloseDB(db)
}
