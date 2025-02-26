package models

type HomeData struct {
	TxCount      int64
	BlockCount   int64
	Transactions []SimpleTransaction
	AvgBlockTime uint64
}

type BlockData struct {
	Block      SimpleBlock
	Txs        []SimpleTransaction
	PageNumber int64
	TotalPages int64
	TxCount    int64
}

type AccountData struct {
	Address    string
	Txs        []SimpleTransaction
	TxCount    int64
	PageSize   int64
	PageNumber int64
	TotalPages int64
}
