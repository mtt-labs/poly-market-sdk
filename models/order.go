package models

import "time"

// OrderType order type enumeration
type OrderType string

const (
	// OrderTypeFOK Fill-Or-Kill: must be filled immediately or cancelled
	OrderTypeFOK OrderType = "FOK"
	// OrderTypeFAK Fill-And-Kill: fill available portion immediately, cancel the rest
	OrderTypeFAK OrderType = "FAK"
	// OrderTypeGTC Good-Til-Cancelled: valid until cancelled
	OrderTypeGTC OrderType = "GTC"
	// OrderTypeGTD Good-Til-Date: valid until specified date
	OrderTypeGTD OrderType = "GTD"
)

// OrderSide order side enumeration (buy or sell)
type OrderSide int

const (
	// OrderSideBuy buy order
	OrderSideBuy OrderSide = 0
	// OrderSideSell sell order
	OrderSideSell OrderSide = 1
)

// SignatureType signature type enumeration
type SignatureType int

const (
	// SignatureTypeEmailMagic EIP712 signatures signed by EOAs that own Polymarket Proxy wallets
	SignatureTypeEmailMagic SignatureType = 1
	// SignatureTypeBrowserWallet EIP712 signatures signed by EOAs that own Polymarket Gnosis safes
	SignatureTypeBrowserWallet SignatureType = 2
	// SignatureTypeEOA ECDSA EIP712 signatures signed by EOAs
	SignatureTypeEOA SignatureType = 0
)

// SignedOrder signed order object (according to Polymarket CLOB API documentation)
// Reference: https://docs.polymarket.com/developers/CLOB/orders/create-order
type SignedOrder struct {
	Salt          int64  `json:"salt"`          // Random salt value for creating unique orders
	Maker         string `json:"maker"`         // Maker address (funder)
	Signer        string `json:"signer"`        // Signer address
	Taker         string `json:"taker"`         // Taker address (operator)
	TokenID       string `json:"tokenId"`       // ERC1155 token ID (conditional token)
	MakerAmount   string `json:"makerAmount"`   // Maximum amount maker is willing to pay
	TakerAmount   string `json:"takerAmount"`   // Minimum amount taker will pay to maker
	Expiration    string `json:"expiration"`    // Unix expiration timestamp
	Nonce         string `json:"nonce"`         // Maker's exchange nonce
	FeeRateBps    string `json:"feeRateBps"`    // Fee rate in basis points, required by operator
	Side          string `json:"side"`          // Buy or sell enumeration index
	SignatureType int    `json:"signatureType"` // Signature type enumeration index
	Signature     string `json:"signature"`     // Hex-encoded signature
}

// CreateOrderRequest create order request (according to Polymarket CLOB API documentation)
// Reference: https://docs.polymarket.com/developers/CLOB/orders/create-order
type CreateOrderRequest struct {
	Order     *SignedOrder `json:"order"`     // Signed order object
	Owner     string       `json:"owner"`     // API key of the order owner
	OrderType OrderType    `json:"orderType"` // Order type ("FOK", "GTC", "GTD", "FAK")
}

// CreateOrderResponse create order response (according to Polymarket CLOB API documentation)
type CreateOrderResponse struct {
	Success     bool     `json:"success"`          // Whether successful (false indicates server-side error)
	ErrorMsg    string   `json:"errorMsg"`         // Error message (if success = false or client error)
	OrderID     string   `json:"orderId"`          // Order ID
	OrderHashes []string `json:"orderHashes"`      // Settlement transaction hashes if order is fillable and triggers matching
	Status      string   `json:"status,omitempty"` // Order status: "matched", "live", "delayed", "unmatched"
}

// OrderStatus order status
type OrderStatus string

const (
	OrderStatusMatched   OrderStatus = "matched"   // Order placed and matched with existing orders
	OrderStatusLive      OrderStatus = "live"      // Order placed and on the order book
	OrderStatusDelayed   OrderStatus = "delayed"   // Order is fillable but affected by matching delay
	OrderStatusUnmatched OrderStatus = "unmatched" // Order is fillable but delay failed, placement succeeded
)

// Order represents an order (for querying order details)
type Order struct {
	ID              string      `json:"id"`
	Status          OrderStatus `json:"status"`
	Owner           string      `json:"owner"`
	MakerAddress    string      `json:"maker_address"`
	Market          string      `json:"market"`
	AssetID         string      `json:"asset_id"`
	Side            string      `json:"side"`
	OriginalSize    string      `json:"original_size"`
	SizeMatched     string      `json:"size_matched"`
	Price           string      `json:"price"`
	Outcome         string      `json:"outcome"`
	Expiration      string      `json:"expiration"`
	OrderType       string      `json:"order_type"`
	AssociateTrades []Trade     `json:"associate_trades"`
	CreatedAt       int64       `json:"created_at"`

	// Legacy fields for backward compatibility
	MarketID     string    `json:"market_id,omitempty"`
	OutcomeID    string    `json:"outcome_id,omitempty"`
	TokenID      string    `json:"token_id,omitempty"`
	Size         string    `json:"size,omitempty"`
	Amount       string    `json:"amount,omitempty"`
	FilledAmount string    `json:"filled_amount,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
}

// OrderListResponse order list response
type OrderListResponse struct {
	Orders []Order `json:"orders"`
	Total  int     `json:"total,omitempty"`
	Page   int     `json:"page,omitempty"`
	Limit  int     `json:"limit,omitempty"`
}

// GetActiveOrdersResponse get active orders response
// Reference: https://docs.polymarket.com/developers/CLOB/orders/get-active-order
type GetActiveOrdersResponse struct {
	Data       []Order `json:"data"`        // Array of orders
	NextCursor string  `json:"next_cursor"` // Next page cursor (for pagination)
	Limit      int     `json:"limit"`       // Limit count
	Count      int     `json:"count"`       // Current returned count
}

// CancelOrderResponse response for cancel order operations
// Reference: https://docs.polymarket.com/developers/CLOB/orders/cancel-orders
type CancelOrderResponse struct {
	Canceled    []string          `json:"canceled"`     // List of canceled orders
	NotCanceled map[string]string `json:"not_canceled"` // Map of order ID -> reason for orders that couldn't be canceled
}

// CancelMarketOrdersParams parameters for canceling orders from a market
// Reference: https://docs.polymarket.com/developers/CLOB/orders/cancel-orders
type CancelMarketOrdersParams struct {
	Market  string // Market condition ID (optional)
	AssetID string // Asset/token ID (optional)
}

// OrderScoringResponse response for single order scoring check
// Reference: https://docs.polymarket.com/developers/CLOB/orders/check-scoring
type OrderScoringResponse struct {
	Scoring bool `json:"scoring"` // Indicates if the order is scoring or not
}

// OrdersScoringResponse response for multiple orders scoring check
// It's a dictionary that maps order ID to scoring status
// Reference: https://docs.polymarket.com/developers/CLOB/orders/check-scoring
type OrdersScoringResponse map[string]bool

// CreateAndPostOrderParams parameters for creating and posting an order
// Reference: https://github.com/Polymarket/clob-client
type CreateAndPostOrderParams struct {
	TokenID string  // ERC1155 token ID (conditional token)
	Price   float64 // Order price
	Side    int     // Order side: 0=BUY, 1=SELL
	Size    float64 // Order size
}

// CreateAndPostOrderConfig configuration for creating and posting an order
// Reference: https://github.com/Polymarket/clob-client
type CreateAndPostOrderConfig struct {
	TickSize string // Price precision (e.g., "0.001"), if empty will be fetched from API automatically
	NegRisk  *bool  // Whether to use negative risk contract, if nil will be fetched from API automatically
}
