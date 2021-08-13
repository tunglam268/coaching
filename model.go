package model

type Transactions struct {
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
	trans = []Transactions{
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
