package polymarket

import (
	"fmt"

	"github.com/mtt-labs/poly-market-sdk/api"
	"github.com/mtt-labs/poly-market-sdk/client"
)

// Polymarket is the main entry point of the SDK
type Polymarket struct {
	Client  *client.Client
	Markets *api.MarketsAPI
	Orders  *api.OrdersAPI
	Auth    *api.AuthAPI
	Events  *api.EventsAPI
	Search  *api.SearchAPI
}

// New creates a new Polymarket SDK instance
// Reference: https://docs.polymarket.com/quickstart/orders/first-order
func New(config *client.Config) (*Polymarket, error) {
	c, err := client.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Polymarket{
		Client:  c,
		Markets: api.NewMarketsAPI(c),
		Orders:  api.NewOrdersAPI(c),
		Auth:    api.NewAuthAPI(c),
		Events:  api.NewEventsAPI(c),
		Search:  api.NewSearchAPI(c),
	}, nil
}

// NewWithDefaults creates SDK instance with default config (only for APIs that don't require authentication)
// Note: Most operations require a private key, please use New() with a Config
func NewWithDefaults() (*Polymarket, error) {
	// Return error because private key is required
	return nil, fmt.Errorf("private key is required, please use New() with a Config")
}
