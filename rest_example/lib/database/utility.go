package database

import (
	"database/sql"
	"fmt"
	"golearn/rest_example/lib"
	"html/template"
	"net/http"
)

func GetFutureQuery() string {
	return fmt.Sprintf(`SELECT ident.Identifier, im.Currency, md.MarketName, 
                        fc.ExpirationDate, contract.* FROM InstrumentMaster AS im
                        JOIN IdentifierMap AS ident ON im.InstrumentID = ident.InstrumentID
                        JOIN MarketDefinition AS md ON im.MarketID = md.MarketID
                        JOIN FuturesCharacteristics AS fc ON im.InstrumentID = fc.InstrumentID
                        JOIN FuturesContractTypes AS contract ON fc.ContractTypeID = contract.ContractTypeID
                        WHERE ident.IdentifierTypeID = 1 AND im.InstTypeID = 2 GROUP BY im.InstrumentID;`)
}

func GetOptionsQuery() string {
	return fmt.Sprintf(`SELECT ident.Identifier, im.Currency, md.MarketName, 
		                max(instrTSM.StartDate) AS StartDt, instrTSM.Multiplier, instrTSM.TickSizeRegimeID, oc.* 
		                FROM InstrumentMaster AS im JOIN IdentifierMap AS ident ON im.InstrumentID = ident.InstrumentID 
	                    JOIN MarketDefinition AS md ON im.MarketID = md.MarketID 
	                    JOIN OptionsCharacteristics AS oc ON im.InstrumentID = oc.InstrumentID 
	                    JOIN InstrumentTickSizeMultiplier AS instrTSM ON im.InstrumentID = instrTSM.InstrumentID 
	                    WHERE ident.IdentifierTypeID = 1 AND im.InstTypeID = 3 GROUP BY im.InstrumentID;`)
}

func GetEquityQuery() string {
	return fmt.Sprintf(`SELECT ident.Identifier, im.Currency, md.MarketName, 
                        max(instrTSM.StartDate) AS StartDt, instrTSM.Multiplier, instrTSM.TickSizeRegimeID
                        FROM InstrumentMaster AS im JOIN IdentifierMap AS ident ON im.InstrumentID = ident.InstrumentID 
                        JOIN MarketDefinition AS md ON im.MarketID = md.MarketID
                        JOIN InstrumentTickSizeMultiplier AS instrTSM ON im.InstrumentID = instrTSM.InstrumentID
                        WHERE ident.IdentifierTypeID = 1 AND im.InstTypeID = 1 GROUP BY im.InstrumentID;`)
}

func GetValidInt(a sql.NullInt64) int {
	if a.Valid {
		return int(a.Int64)
	}
	return 0
}

func GetValidString(a sql.NullString) string {
	if a.Valid {
		return string(a.String)
	}
	return ""
}

func GetValidFloat(a sql.NullFloat64) float64 {
	if a.Valid {
		return float64(a.Float64)
	}
	return 0
}

func WriteHtmlFutures(W http.ResponseWriter, results []lib.Futures) error {
	futures, err := template.ParseFiles("Future.html")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return futures.ExecuteTemplate(W, "Future.html", results)
}
