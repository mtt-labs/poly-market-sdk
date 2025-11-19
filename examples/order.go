package main

import (
	"fmt"
	"log"

	"github.com/mtt-labs/poly-market-sdk/api"
	"github.com/mtt-labs/poly-market-sdk/client"
	"github.com/mtt-labs/poly-market-sdk/models"
	"github.com/mtt-labs/poly-market-sdk/polymarket"
)

func main() {
	// Get private key from environment variable (please keep private key secure in actual use)
	privateKey := "b418828d5e0ac24d2b9ef418aad29b74ecfef9cb2b56d4ccfe79109b5742f7f5" //#This is your Private Key. Export from https://reveal.magic.link/polymarket or from your Web3 Extension
	if privateKey == "" {
		log.Fatal("Please set POLYMARKET_PRIVATE_KEY environment variable")
	}

	// Create client configuration
	// Reference: https://docs.polymarket.com/quickstart/orders/first-order
	config := &client.Config{
		PrivateKey:    privateKey,
		ChainID:       137,                                          // Polygon mainnet
		SignatureType: client.SignatureTypeEmailMagic,               // 0=EOA, 1=Email/Magic, 2=Browser Wallet
		Funder:        "0x47D348AD8bEB6F1259F5326859e9F185C28B0348", // This is the address listed below your profile picture when using the Polymarket site.
		// Funder: "0x...", // If using Proxy, set this address
	}

	// Create SDK instance
	sdk, err := polymarket.New(config)
	if err != nil {
		log.Fatalf("Failed to create SDK: %v", err)
	}

	limit := 10
	offset := 0
	markets, err := sdk.Markets.GetMarkets(&api.ListMarketsParams{
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil {
		log.Printf("Failed to get markets: %v", err)
		return
	}
	fmt.Printf("Found %d markets\n", len(markets))

	// Create or derive API credentials (required for first-time use)
	creds, err := sdk.Auth.DeriveAPICredentials()
	if err != nil {
		log.Printf("Failed to get API credentials: %v", err)
		log.Println("Note: Some operations may require creating API credentials first")
	} else {
		fmt.Printf("API Key: %s\n", creds.Key)
	}

	sdk.Client.SetAPICredentials(creds.Key, creds.Secret, creds.Passphrase)

	//orderResp, err := sdk.Orders.CreateAndPostOrder(&models.CreateAndPostOrderParams{
	//	TokenID: "82202994941777288823087700378947997402758931782856358751785693028970403111094", // ERC1155 token ID
	//	Price:   0.5,                                                                             // Order price
	//	Side:    0,                                                                               // 0=BUY, 1=SELL
	//	Size:    5,
	//}, &models.CreateAndPostOrderConfig{}, models.OrderTypeGTC)
	//if err != nil {
	//	log.Printf("Failed to create order: %v", err)
	//} else {
	//	fmt.Printf("Order created successfully: %v", orderResp)
	//}

	// fmt.Println("SDK initialized successfully!")
	searchResp, err := sdk.Search.Search(&models.SearchParams{
		Q: "What price will Bitcoin hit in 2025",
	})
	if err != nil {
		log.Printf("Failed to search: %v", err)
		return
	}
	fmt.Println(searchResp)

	event, err := sdk.Events.GetEventByID("16096")
	if err != nil {
		return
	}
	fmt.Println(event)
}
