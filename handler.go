package api

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	DBPostGres "github.com/tunglam268/coaching/blob/main/model.go/db"
	CoachingModel "github.com/tunglam268/coaching/blob/main/model.go/model"
)

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	db := DBPostGres.DBConnection{}
	db = *DBPostGres.NewDBConnection()
	db.OpenConnection()
	rows, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		log.Fatal(err)
	}

	var transactions []CoachingModel.Transaction

	for rows.Next() {
		var trans CoachingModel.Transaction
		rows.Scan(&trans.Id, &trans.TraceNumber, &trans.Block_id, &trans.Txhash, &trans.Txcount)
		transactions = append(transactions, trans)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	db := DBPostGres.DBConnection{}
	db = *DBPostGres.NewDBConnection()
	db.OpenConnection()
	var t CoachingModel.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `INSERT INTO transactions (id , tracenumber , block_id, txhash , txcount) VALUES ($1,$2,$3,$4,$5)`
	_, err = db.Exec(sqlStatement, t.Id, t.TraceNumber, t.Block_id, t.Txhash, t.Txcount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
