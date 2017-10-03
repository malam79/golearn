package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golearn/rest_example/lib"
)

const data_base string = "db/reference.db"

type SqlDB struct{}

func ConnectSQLDB() *sql.DB {
	db, err := sql.Open("sqlite3", data_base)
	if err != nil {
		fmt.Println("Fail to connect to db")
	}
	return db
}

func (SqlDB) CloseDB(db interface{}) {
	(db.(*sql.DB)).Close()
}

func (SqlDB) RunQuery(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Fail to execute query: %s\n", query)
	}
}

func (SqlDB) GetFutures(db interface{}) (results []lib.Futures, err error) {
	rows, err := (db.(*sql.DB)).Query(GetFutureQuery())
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var identifier, currency, marketName, expiryDate, currentRootSymbol sql.NullString
		var underlier, id sql.NullInt64
		var spreadticksize sql.NullFloat64
		var dtype sql.NullInt64
		var imargin, mmargin, mimargin, mmmargin sql.NullFloat64
		err = rows.Scan(&identifier, &currency, &marketName, &expiryDate, &id, &currentRootSymbol,
			&underlier, &spreadticksize, &dtype, &imargin, &mmargin, &mimargin, &mmmargin)
		if err != nil {
			return nil, err
		}
		results = append(results, lib.Futures{
			lib.Instrument{GetValidString(currency),
				GetValidString(marketName),
				GetValidString(identifier),
			},
			GetValidString(currentRootSymbol),
			GetValidString(expiryDate),
			GetValidInt(underlier),
			GetValidFloat(spreadticksize),
			rune(GetValidInt(dtype)),
			GetValidFloat(imargin),
			GetValidFloat(mmargin),
			GetValidFloat(mimargin),
			GetValidFloat(mmmargin),
		})
	}
	return results, nil
}

func (SqlDB) GetEquity(db interface{}) (results []lib.Equity, err error) {
	rows, err := (db.(*sql.DB)).Query(GetEquityQuery())
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var identifier, currency, marketName, startDate sql.NullString
		var multiplier sql.NullFloat64
		var ticksize sql.NullInt64
		err = rows.Scan(&identifier, &currency, &marketName, &startDate, &multiplier, &ticksize)
		results = append(results, lib.Equity{
			lib.Instrument{GetValidString(currency),
				GetValidString(marketName),
				GetValidString(identifier)},
			GetValidString(startDate),
			GetValidFloat(multiplier),
			GetValidInt(ticksize),
		})
	}
	return results, nil
}
