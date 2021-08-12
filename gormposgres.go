package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type transactions struct {
	id                       int64
	trace_number             string
	blockid                  int64
	txhash                   string
	txcount                  int64
	chaincodename            string
	status                   int64
	creator_msp_id           string
	endorser_msp_id          string
	chaincode_id             string
	write_set                string
	block_hash               string
	channel_genesis_hash     string
	validation_code          string
	envelope_signature       string
	payload_extension        string
	creator_id_bytes         string
	creator_nonce            string
	chaincode_proposal_input string
	tx_response              string
	payload_proposal_hash    string
	endorser_id_bytes        string
	endorser_signature       string
	network_name             string
}

var (
	trans = []transactions{
		{id: 1, trace_number: "test",
			blockid: 2, txhash: "test", txcount: 3,
			chaincodename: "test", status: 4,
			creator_msp_id: "test", endorser_msp_id: "test", chaincode_id: "test",
			write_set: "test", block_hash: "test", channel_genesis_hash: "test",
			validation_code: "test", envelope_signature: "test", payload_extension: "test",
			creator_id_bytes: "test", creator_nonce: "test", chaincode_proposal_input: "test",
			tx_response: "test", payload_proposal_hash: "test", endorser_id_bytes: "test", endorser_signature: "test", network_name: "test"},
	}
)

func main() {
	host := "127.0.0.1"
	username := "hppoc"
	password := "password"
	dbname := "fabricexplorer"
	port := "5432"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connect Success !")
	}

	for index := range trans {

		db.Create(&trans[index])

	}

	id := 1
	var Transactions transactions
	db.Raw("SELECT * FROM transactions WHERE id = ?", id).Scan(Transactions)

	excel(&Transactions)
}
func excel(Transactions *transactions) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", Transactions.id)
	f.SetCellValue("Sheet1", "B1", Transactions.trace_number)
	f.SetCellValue("Sheet1", "C1", Transactions.blockid)
	f.SetCellValue("Sheet1", "D1", Transactions.txhash)
	f.SetCellValue("Sheet1", "E1", Transactions.txcount)
	f.SetCellValue("Sheet1", "G1", Transactions.chaincodename)
	f.SetCellValue("Sheet1", "H1", Transactions.status)
	f.SetCellValue("Sheet1", "I1", Transactions.creator_msp_id)
	f.SetCellValue("Sheet1", "J1", Transactions.endorser_msp_id)
	f.SetCellValue("Sheet1", "K1", Transactions.chaincode_id)
	f.SetCellValue("Sheet1", "M1", Transactions.block_hash)
	f.SetCellValue("Sheet1", "N1", Transactions.channel_genesis_hash)
	f.SetCellValue("Sheet1", "O1", Transactions.validation_code)
	f.SetCellValue("Sheet1", "P1", Transactions.envelope_signature)
	f.SetCellValue("Sheet1", "Q1", Transactions.payload_extension)
	f.SetCellValue("Sheet1", "R1", Transactions.creator_id_bytes)
	f.SetCellValue("Sheet1", "S1", Transactions.creator_nonce)
	f.SetCellValue("Sheet1", "T1", Transactions.tx_response)
	f.SetCellValue("Sheet1", "U1", Transactions.payload_proposal_hash)
	f.SetCellValue("Sheet1", "V1", Transactions.endorser_id_bytes)
	f.SetCellValue("Sheet1", "W1", Transactions.network_name)
	f.SetCellValue("Sheet1", "X1", Transactions.endorser_signature)

}
