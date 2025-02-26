package helpers

import (
	"math/big"
	"strings"
	"time"
)

func TruncateHash(hash string) string {
	return hash[:8] + "..." + hash[len(hash)-8:]
}

func FormatEthValue(value string) string {
	if len(value) < 18 {
		value = strings.Repeat("0", 18-len(value)) + value
	}
	weiValue, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "0"
	}
	ethValue := new(big.Float).SetInt(weiValue)
	ethValue = ethValue.Quo(ethValue, big.NewFloat(1e18))
	return ethValue.Text('f', 18)
}

func FormatWeiValue(value string) string {
	if len(value) < 18 {
		value = strings.Repeat("0", 18-len(value)) + value
	}
	weiValue, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "0"
	}
	return weiValue.String()
}

func FormatTimestamp(timestamp uint64) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}

func IsAccountAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if !strings.HasPrefix(address, "0x") {
		return false
	}
	return true
}
