package models

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

type HealthCheckResponse struct {
	Service string `json:"service"`
	Status  bool   `json:"status"`
}

type RegisterUserResponse struct {
	ApiKey string `json:"api_key"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type SimpleTransaction struct {
	Hash        string `json:"hash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Amount      string `json:"amount"`
	GasPrice    string `json:"gas_price"`
	GasLimit    uint64 `json:"gas_limit"`
	GasUsed     uint64 `json:"gas_used"`
	Nonce       uint64 `json:"nonce"`
	BlockNumber uint64 `json:"block_number"`
	Value       string `json:"value"`
}

type SimpleBlock struct {
	Number       uint64              `json:"number"`
	Nonce        string              `json:"nonce"`
	Hash         string              `json:"hash"`
	ParentHash   string              `json:"parent_hash"`
	Size         uint64              `json:"size"`
	Miner        string              `json:"miner"`
	Timestamp    uint64              `json:"timestamp"`
	Transactions []SimpleTransaction `json:"transactions"`
}

type SimpleEvent struct {
	LogIndex        uint     `json:"log_index"`
	TransactionHash string   `json:"transaction_hash"`
	BlockNumber     uint64   `json:"block_number"`
	Address         string   `json:"address"`
	Data            string   `json:"data"`
	Topics          []string `json:"topics"`
}
