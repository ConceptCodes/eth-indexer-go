package models

type HomeData struct {
	TxCount      int64
	BlockCount   int64
	Transactions []SimpleTransaction
	AvgBlockTime uint64
}
