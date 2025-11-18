package models

// Account represents user account information
type Account struct {
	Address     string `json:"address"`
	Balance     string `json:"balance"`
	TotalVolume string `json:"total_volume"`
	TotalTrades int    `json:"total_trades"`
}

// Position represents user position in a market
type Position struct {
	MarketID   string `json:"market_id"`
	OutcomeID  string `json:"outcome_id"`
	Outcome    string `json:"outcome"`
	Amount     string `json:"amount"`
	AvgPrice   string `json:"avg_price"`
	TotalValue string `json:"total_value"`
}

// PositionListResponse position list response
type PositionListResponse struct {
	Positions []Position `json:"positions"`
	Total     int        `json:"total"`
}
