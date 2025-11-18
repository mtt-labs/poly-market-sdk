# Polymarket Go SDK

A Go SDK for interacting with the Polymarket prediction market platform.

## Features

- 📊 **Market Data** - Get market lists, details, search markets
- 📈 **Order Management** - Create, query, cancel orders
- 💼 **Account Information** - Query account balance and positions
- 🔍 **Trade History** - Get market trade records
- 📖 **Orderbook** - View market orderbook

## Installation

```bash
go get github.com/yourusername/poly-market-sdk
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "poly-market-sdk/api"
    "poly-market-sdk/client"
    "poly-market-sdk/polymarket"
)

func main() {
    // Create SDK with configuration
    config := &client.Config{
        PrivateKey: "your-private-key",
        ChainID:    137,
    }
    sdk, err := polymarket.New(config)
    if err != nil {
        log.Fatal(err)
    }
    
    // Get market list
    limit := 10
    offset := 0
    markets, err := sdk.Markets.GetMarkets(&api.ListMarketsParams{
        Limit:  &limit,
        Offset: &offset,
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d markets\n", len(markets))
}
```

### Custom Configuration

```go
import (
    "time"
    "poly-market-sdk/client"
    "poly-market-sdk/polymarket"
)

// Create custom configuration
config := &client.Config{
    BaseURL: "https://clob.polymarket.com",
    APIKey:  "your-api-key", // If authentication is required
    Timeout: 30 * time.Second,
}

sdk := polymarket.New(config)
```

## API Usage Examples

### Market API

```go
// Get market list
limit := 10
offset := 0
markets, err := sdk.Markets.GetMarkets(&api.ListMarketsParams{
    Limit:  &limit,
    Offset: &offset,
})

// Get market details by ID
market, err := sdk.Markets.GetMarketByID("market-id")

// Get market by slug
market, err := sdk.Markets.GetMarketBySlug("market-slug")

// Get market tags
tags, err := sdk.Markets.GetMarketTags("market-id")

// Search markets, events, and profiles (use Search API)
searchResults, err := sdk.Search.Search(&models.SearchParams{
    Q: "bitcoin",
})

// Get market trade history
trades, err := sdk.Markets.GetMarketTrades("market-id", 1, 20)

// Get orderbook
orderbook, err := sdk.Markets.GetMarketOrderbook("market-id", "outcome-id")
```

### Order API

```go
// Create order
orderReq := &models.CreateOrderRequest{
    MarketID:  "market-id",
    OutcomeID: "outcome-id",
    Side:      models.OrderSideBid,
    Price:     "0.50",
    Amount:    "100",
}

order, err := sdk.Orders.CreateOrder(orderReq)

// Get order details
order, err := sdk.Orders.GetOrder("order-id")

// Get active orders list
ordersResp, err := sdk.Orders.GetActiveOrders(&api.GetActiveOrdersParams{
    Market: "0x...", // Optional: specify market condition ID
})
// ordersResp.Data contains array of orders
// ordersResp.NextCursor is for pagination
// ordersResp.Count is the current returned count

// Cancel single order
err := sdk.Orders.CancelOrder("order-id")

// Cancel multiple orders in batch
err := sdk.Orders.CancelOrders([]string{"order-id-1", "order-id-2"})
```

### Account API

```go
// Get account information
account, err := sdk.Account.GetAccount("0x...")

// Get account positions
positions, err := sdk.Account.GetPositions("0x...")
```

## Project Structure

```
poly-market-sdk/
├── client/          # HTTP client
├── api/             # API interface wrappers
│   ├── markets.go   # Market-related API
│   ├── orders.go    # Order-related API
│   └── account.go   # Account-related API
├── models/          # Data models
│   ├── market.go    # Market model
│   ├── order.go     # Order model
│   ├── account.go   # Account model
│   └── trade.go     # Trade model
├── polymarket.go    # SDK main entry point
└── main.go          # Example code
```

## Development

```bash
# Run examples
go run main.go

# Run tests
go test ./...

# Build
go build
```

## Notes

- Current SDK is designed based on Polymarket's public API
- Some features may require API key authentication
- API endpoints may need to be adjusted according to actual Polymarket API documentation
- It is recommended to test thoroughly before using in production environment

## License

MIT License

## Contributing

Issues and Pull Requests are welcome!
