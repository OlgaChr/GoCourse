package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// SleepTime периодичность запроса
const SleepTime = 1 * time.Second

// BinanceTrade сделка binance
type BinanceTrade struct {
	ID           int
	IsBestMatch  bool
	IsBuyerMaker bool
	Price        string
	Qty          string
	QuoteQty     string
	Time         int
}

// PoloniexTrade сделка poloniex
type PoloniexTrade struct {
	Amount        string
	Date          string
	GlobalTradeID int
	OrderNumber   int
	Rate          string
	Total         string
	TradeID       int
	Type          string
}

// TradesStore хранилище
type TradesStore struct {
	BinanceTrades  []BinanceTrade
	PoloniexTrades []PoloniexTrade
}

// TradeClient клиент
type TradeClient struct {
	client *http.Client
}

func newTradeClient() TradeClient {
	return TradeClient{
		client: &http.Client{},
	}
}

func (tc *TradeClient) getTrades(url string, params map[string]string) (body []byte, err error) {
	// создаем get запрос
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return make([]byte, 0), err
	}
	// определяем парметры запроса
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	// выполняем запрос
	resp, err := tc.client.Do(req)

	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return make([]byte, 0), err
	}

	return body, nil
}

func main() {
	trades := TradesStore{}
	tradeClient := newTradeClient()
	// получаем торги с binance
	go func() {
		for {
			result, err := tradeClient.getTrades("https://api.binance.com/api/v3/trades", map[string]string{"symbol": "ETHBTC"})
			if err != nil {
				log.Println(err)
			}
			json.Unmarshal(result, &trades.BinanceTrades)
			time.Sleep(SleepTime)
			fmt.Println("binance", &trades.BinanceTrades[0])
		}
	}()
	// получаем торги с poloniex
	for {
		result, err := tradeClient.getTrades("https://poloniex.com/public?command=returnTradeHistory", map[string]string{"currencyPair": "BTC_ETH"})
		if err != nil {
			log.Println(err)
		}
		json.Unmarshal(result, &trades.PoloniexTrades)
		time.Sleep(SleepTime)
		fmt.Println("poloniex", &trades.PoloniexTrades[0])
	}
}
