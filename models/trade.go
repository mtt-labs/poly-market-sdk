package models

import "time"

// Trade represents a trade
type Trade struct {
	ID        string    `json:"id"`
	MarketID  string    `json:"market_id"`
	OutcomeID string    `json:"outcome_id"`
	Price     string    `json:"price"`
	Amount    string    `json:"amount"`
	Side      string    `json:"side"`
	Buyer     string    `json:"buyer"`
	Seller    string    `json:"seller"`
	Timestamp time.Time `json:"timestamp"`
	TxHash    string    `json:"tx_hash"`
}

// Orderbook represents an orderbook
type Orderbook struct {
	MarketID  string           `json:"market_id"`
	OutcomeID string           `json:"outcome_id"`
	Bids      []OrderbookEntry `json:"bids"`
	Asks      []OrderbookEntry `json:"asks"`
}

// OrderbookEntry orderbook entry
type OrderbookEntry struct {
	Price  string `json:"price"`
	Amount string `json:"amount"`
}
