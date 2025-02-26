package helpers

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func TruncateHash(hash string) string {
	if len(hash) < 16 {
		return hash
	}
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
	return ethValue.Text('f', 6)
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

func ConvertSizeToKb(size uint64) float64 {
	return float64(size) / 1024
}

func GetCurrentEthPrice() string {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"
	resp, err := http.Get(url)
	if err != nil {
		return "0"
	}
	defer resp.Body.Close()
	var data map[string]map[string]float64
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "0"
	}
	return strconv.FormatFloat(data["ethereum"]["usd"], 'f', 6, 64)
}

func GetEthPriceInUsd(amount string) string {
	price, err := strconv.ParseFloat(GetCurrentEthPrice(), 64)
	if err != nil {
		return "0"
	}
	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "0"
	}
	return strconv.FormatFloat(price*amt, 'f', 6, 64)
}
