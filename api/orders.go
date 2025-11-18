package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/mtt-labs/poly-market-sdk/auth"
	"github.com/mtt-labs/poly-market-sdk/client"
	"github.com/mtt-labs/poly-market-sdk/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/builder"
	ordermodel "github.com/polymarket/go-order-utils/pkg/model"
)

const COLLATERAL_TOKEN_DECIMALS = 6
const CONDITIONAL_TOKEN_DECIMALS = 6

// OrdersAPI provides order-related API methods
type OrdersAPI struct {
	client    *client.Client
	tickSizes map[string]string // Cache for tickSize, key is tokenID
	feeRates  map[string]int    // Cache for feeRateBps, key is tokenID
	negRisks  map[string]bool   // Cache for negRisk, key is tokenID
	mu        sync.RWMutex      // RWMutex to protect caches
}

// NewOrdersAPI creates a new OrdersAPI instance
func NewOrdersAPI(c *client.Client) *OrdersAPI {
	return &OrdersAPI{
		client:    c,
		tickSizes: make(map[string]string),
		feeRates:  make(map[string]int),
		negRisks:  make(map[string]bool),
	}
}

// generateL2Headers generates L2 authentication headers (helper function)
func (o *OrdersAPI) generateL2Headers(method, path, body string) (*auth.L2AuthHeaders, error) {
	apiKey := o.client.GetAPIKey()
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	secret := o.client.GetAPISecret()
	passphrase := o.client.GetAPIPassphrase()
	if secret == "" || passphrase == "" {
		return nil, fmt.Errorf("API secret and passphrase are required")
	}

	address := o.client.GetAddress()
	timestamp := time.Now().Unix()

	signer := o.client.GetSigner()
	l2Headers, err := signer.SignL2Auth(address, method, path, body, timestamp, apiKey, secret, passphrase)
	if err != nil {
		return nil, fmt.Errorf("sign L2 auth: %w", err)
	}

	return l2Headers, nil
}

// CreateOrder creates and submits an order (according to Polymarket CLOB API documentation)
// Reference: https://docs.polymarket.com/developers/CLOB/orders/create-order
// This endpoint requires L2 Header (API key)
//
// Parameters:
//   - signedOrder: Signed order object (needs to be signed in advance)
//   - orderType: Order type (FOK, GTC, GTD, FAK)
//   - apiKey: API key of the order owner (if empty, will use the API key from client config)
func (o *OrdersAPI) CreateOrder(signedOrder *models.SignedOrder, orderType models.OrderType, apiKey string) (*models.CreateOrderResponse, error) {
	endpoint := "/order"

	// If apiKey is not provided, try to get it from client config
	if apiKey == "" {
		apiKey = o.client.GetAPIKey()
		if apiKey == "" {
			return nil, fmt.Errorf("API key is required for creating orders")
		}
	}

	// Check if API secret and passphrase are available
	secret := o.client.GetAPISecret()
	passphrase := o.client.GetAPIPassphrase()
	if secret == "" || passphrase == "" {
		return nil, fmt.Errorf("API secret and passphrase are required for creating orders")
	}

	req := &models.CreateOrderRequest{
		Order:     signedOrder,
		Owner:     apiKey,
		OrderType: orderType,
	}

	// Serialize request body for signing (use compact format to ensure consistency with sent body)
	reqBodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}
	reqBodyStr := string(reqBodyBytes)

	// Generate L2 headers (use helper function to ensure consistency)
	l2Headers, err := o.generateL2Headers("POST", endpoint, reqBodyStr)
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send request with L2 headers
	data, err := o.client.PostWithL2(endpoint, req, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}

	var response models.CreateOrderResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	// Check if there is an error message
	if response.ErrorMsg != "" {
		return &response, fmt.Errorf("order placement error: %s", response.ErrorMsg)
	}

	// If success is false, it indicates a server-side error
	if !response.Success {
		return &response, fmt.Errorf("server error: %s", response.ErrorMsg)
	}

	return &response, nil
}

// GetOrder gets order details
// This endpoint requires L2 headers
func (o *OrdersAPI) GetOrder(orderID string) (*models.Order, error) {
	endpoint := fmt.Sprintf("/orders/%s", orderID)

	// Generate L2 headers
	l2Headers, err := o.generateL2Headers("GET", endpoint, "")
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send request with L2 headers
	data, err := o.client.GetWithL2(endpoint, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("get order: %w", err)
	}

	var order models.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &order, nil
}

// GetActiveOrdersParams parameters for getting active orders
type GetActiveOrdersParams struct {
	ID      string // Order ID (optional)
	Market  string // Market condition ID (optional)
	AssetID string // Asset/token ID (optional)
}

// GetActiveOrders gets a list of active orders
// This endpoint requires L2 headers
// Reference: https://docs.polymarket.com/developers/CLOB/orders/get-active-order
func (o *OrdersAPI) GetActiveOrders(params *GetActiveOrdersParams) (*models.GetActiveOrdersResponse, error) {
	endpoint := "/data/orders"

	// Build query parameters
	if params != nil {
		queryValues := url.Values{}
		if params.ID != "" {
			queryValues.Set("id", params.ID)
		}
		if params.Market != "" {
			queryValues.Set("market", params.Market)
		}
		if params.AssetID != "" {
			queryValues.Set("asset_id", params.AssetID)
		}

		// If there are query parameters, add them to endpoint
		if len(queryValues) > 0 {
			endpoint = endpoint + "?" + queryValues.Encode()
		}
	}

	// Generate L2 headers
	l2Headers, err := o.generateL2Headers("GET", endpoint, "")
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send request with L2 headers
	data, err := o.client.GetWithL2(endpoint, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("get active orders: %w", err)
	}

	// Response format: { "data": [], "next_cursor": "...", "limit": 500, "count": 0 }
	var response models.GetActiveOrdersResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &response, nil
}

// CancelOrder cancels a single order
// This endpoint requires L2 headers
// Reference: https://docs.polymarket.com/developers/CLOB/orders/cancel-orders
func (o *OrdersAPI) CancelOrder(orderID string) (*models.CancelOrderResponse, error) {
	endpoint := "/order"

	// Build request body with orderID
	reqBody := map[string]interface{}{
		"orderID": orderID,
	}

	// Serialize request body for signing
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// Generate L2 headers (DELETE request with body)
	l2Headers, err := o.generateL2Headers("DELETE", endpoint, string(reqBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send DELETE request with L2 headers
	data, err := o.client.DeleteWithL2(endpoint, reqBody, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("cancel order: %w", err)
	}

	var response models.CancelOrderResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &response, nil
}

// CancelOrders cancels multiple orders in batch
// This endpoint requires L2 headers
// Reference: https://docs.polymarket.com/developers/CLOB/orders/cancel-orders
func (o *OrdersAPI) CancelOrders(orderIDs []string) (*models.CancelOrderResponse, error) {
	endpoint := "/orders"

	// Build request body (array of order IDs)
	reqBody := orderIDs

	// Serialize request body for signing
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// Generate L2 headers (DELETE request with body)
	l2Headers, err := o.generateL2Headers("DELETE", endpoint, string(reqBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send DELETE request with L2 headers
	data, err := o.client.DeleteWithL2(endpoint, reqBody, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("cancel orders: %w", err)
	}

	var response models.CancelOrderResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &response, nil
}

// CancelAllOrders cancels all open orders posted by a user
// This endpoint requires L2 headers
// Reference: https://docs.polymarket.com/developers/CLOB/orders/cancel-orders
func (o *OrdersAPI) CancelAllOrders() (*models.CancelOrderResponse, error) {
	endpoint := "/cancel-all"

	// Generate L2 headers (DELETE request with empty body)
	l2Headers, err := o.generateL2Headers("DELETE", endpoint, "")
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send DELETE request with L2 headers
	data, err := o.client.DeleteWithL2(endpoint, nil, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("cancel all orders: %w", err)
	}

	var response models.CancelOrderResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &response, nil
}

// CancelMarketOrders cancels orders from a market
// This endpoint requires L2 headers
// Reference: https://docs.polymarket.com/developers/CLOB/orders/cancel-orders
func (o *OrdersAPI) CancelMarketOrders(params *models.CancelMarketOrdersParams) (*models.CancelOrderResponse, error) {
	endpoint := "/cancel-market-orders"

	// Build request body
	reqBody := make(map[string]interface{})
	if params != nil {
		if params.Market != "" {
			reqBody["market"] = params.Market
		}
		if params.AssetID != "" {
			reqBody["asset_id"] = params.AssetID
		}
	}

	// Serialize request body for signing
	var reqBodyBytes []byte
	var err error
	if len(reqBody) > 0 {
		reqBodyBytes, err = json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("marshal request: %w", err)
		}
	}

	// Generate L2 headers (DELETE request with body)
	l2Headers, err := o.generateL2Headers("DELETE", endpoint, string(reqBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send DELETE request with L2 headers
	var body interface{}
	if len(reqBody) > 0 {
		body = reqBody
	}
	data, err := o.client.DeleteWithL2(endpoint, body, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("cancel market orders: %w", err)
	}

	var response models.CancelOrderResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &response, nil
}

// CheckOrderScoring checks if a single order is eligible or scoring for Rewards purposes
// This endpoint requires L2 headers
// Reference: https://docs.polymarket.com/developers/CLOB/orders/check-scoring
func (o *OrdersAPI) CheckOrderScoring(orderID string) (*models.OrderScoringResponse, error) {
	endpoint := "/order-scoring"

	// Build query parameters
	queryValues := url.Values{}
	queryValues.Set("order_id", orderID)
	endpoint = endpoint + "?" + queryValues.Encode()

	// Generate L2 headers
	l2Headers, err := o.generateL2Headers("GET", endpoint, "")
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send request with L2 headers
	data, err := o.client.GetWithL2(endpoint, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("check order scoring: %w", err)
	}

	var response models.OrderScoringResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &response, nil
}

// CheckOrdersScoring checks if multiple orders are eligible or scoring for Rewards purposes
// This endpoint requires L2 headers
// Reference: https://docs.polymarket.com/developers/CLOB/orders/check-scoring
func (o *OrdersAPI) CheckOrdersScoring(orderIDs []string) (models.OrdersScoringResponse, error) {
	endpoint := "/orders-scoring"

	// Build request body
	reqBody := map[string]interface{}{
		"orderIds": orderIDs,
	}

	// Serialize request body for signing
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// Generate L2 headers
	l2Headers, err := o.generateL2Headers("POST", endpoint, string(reqBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("generate L2 headers: %w", err)
	}

	// Send request with L2 headers
	data, err := o.client.PostWithL2(endpoint, reqBody, l2Headers)
	if err != nil {
		return nil, fmt.Errorf("check orders scoring: %w", err)
	}

	var response models.OrdersScoringResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return response, nil
}

// CreateAndPostOrder creates and posts an order (similar to createAndPostOrder in TypeScript client)
// Reference: https://github.com/Polymarket/clob-client
// This method uses go-order-utils to build and sign orders
func (o *OrdersAPI) CreateAndPostOrder(
	params *models.CreateAndPostOrderParams,
	config *models.CreateAndPostOrderConfig,
	orderType models.OrderType,
) (*models.CreateOrderResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params is required")
	}
	if config == nil {
		return nil, fmt.Errorf("config is required")
	}

	// Get private key
	privateKeyHex := o.client.GetPrivateKey()
	if privateKeyHex == "" {
		return nil, fmt.Errorf("private key is required")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	// Get address and configuration
	address := o.client.GetAddress()
	chainID := o.client.GetChainID()
	funder := o.client.GetFunder()
	if funder == "" {
		funder = address // If no funder, use address
	}

	signatureTypeInt := int(o.client.GetSignatureType())
	if signatureTypeInt == 0 {
		signatureTypeInt = int(client.SignatureTypeEOA)
	}

	// Get tickSize (if not provided, fetch from API)
	tickSizeStr := config.TickSize
	if tickSizeStr == "" {
		var err error
		tickSizeStr, err = o.GetTickSize(params.TokenID)
		if err != nil {
			return nil, fmt.Errorf("get tick size: %w", err)
		}
	}

	// Parse tickSize
	tickSize, err := strconv.ParseFloat(tickSizeStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid tickSize: %w", err)
	}

	// Calculate price (round to tickSize)
	roundedPrice := roundToTickSize(params.Price, tickSize)

	var makerAmount, takerAmount *big.Int
	if params.Side == 0 { // BUY
		// makerAmount = price * size (USDC, 6 decimals)
		// Use big.Float for precise calculation
		makerAmountFloat := big.NewFloat(roundedPrice * params.Size)
		makerAmountFloat.Mul(makerAmountFloat, big.NewFloat(1e6)) // USDC has 6 decimals
		makerAmount, _ = makerAmountFloat.Int(nil)
		// takerAmount = size (tokens, 6 decimals)
		takerAmountFloat := big.NewFloat(params.Size)
		takerAmountFloat.Mul(takerAmountFloat, big.NewFloat(1e6)) // Tokens have 6 decimals
		takerAmount, _ = takerAmountFloat.Int(nil)
	} else { // SELL
		// makerAmount = size (tokens, 6 decimals)
		makerAmountFloat := big.NewFloat(params.Size)
		makerAmountFloat.Mul(makerAmountFloat, big.NewFloat(1e6)) // Tokens have 6 decimals
		makerAmount, _ = makerAmountFloat.Int(nil)
		// takerAmount = price * size (USDC, 6 decimals)
		takerAmountFloat := big.NewFloat(roundedPrice * params.Size)
		takerAmountFloat.Mul(takerAmountFloat, big.NewFloat(1e6)) // USDC has 6 decimals
		takerAmount, _ = takerAmountFloat.Int(nil)
	}

	// Get nonce (default 0, should actually be fetched from API)
	nonce := big.NewInt(0)

	// Get feeRateBps from API
	feeRateBpsInt, err := o.GetFeeRateBps(params.TokenID)
	if err != nil {
		return nil, fmt.Errorf("get fee rate bps: %w", err)
	}
	feeRateBps := big.NewInt(int64(feeRateBpsInt))

	// Set expiration time
	// Only GTD (Good-Til-Date) orders need expiration, others should be 0
	var expiration int64
	if orderType == models.OrderTypeGTD {
		expiration = time.Now().Add(30 * 24 * time.Hour).Unix()
	} else {
		expiration = 0
	}

	// Get negRisk (if not provided, fetch from API)
	negRisk := config.NegRisk
	if negRisk == nil {
		var err error
		negRiskValue, err := o.GetNegRisk(params.TokenID)
		if err != nil {
			return nil, fmt.Errorf("get neg risk: %w", err)
		}
		negRisk = &negRiskValue
	}

	// Determine which contract to use
	var contract ordermodel.VerifyingContract
	if *negRisk {
		contract = ordermodel.NegRiskCTFExchange
	} else {
		contract = ordermodel.CTFExchange
	}

	// Build OrderData
	orderData := &ordermodel.OrderData{
		Maker:         funder,
		Taker:         common.HexToAddress("0x0000000000000000000000000000000000000000").Hex(), // Public order
		TokenId:       params.TokenID,
		MakerAmount:   makerAmount.String(),
		TakerAmount:   takerAmount.String(),
		FeeRateBps:    feeRateBps.String(),
		Nonce:         nonce.String(),
		Expiration:    strconv.FormatInt(expiration, 10),
		Side:          params.Side,
		SignatureType: signatureTypeInt,
		Signer:        address, // Signer address
	}

	// Use go-order-utils to build and sign order
	chainIDBig := big.NewInt(int64(chainID))
	orderBuilder := builder.NewExchangeOrderBuilderImpl(chainIDBig, nil)
	signedOrder, err := orderBuilder.BuildSignedOrder(privateKey, orderData, contract)
	if err != nil {
		return nil, fmt.Errorf("build signed order: %w", err)
	}

	// Convert to our SignedOrder format
	ourSignedOrder := &models.SignedOrder{
		Salt:          signedOrder.Salt.Int64(),
		Maker:         signedOrder.Maker.Hex(),
		Signer:        signedOrder.Signer.Hex(),
		Taker:         signedOrder.Taker.Hex(),
		TokenID:       signedOrder.TokenId.String(),
		MakerAmount:   signedOrder.MakerAmount.String(),
		TakerAmount:   signedOrder.TakerAmount.String(),
		Expiration:    signedOrder.Expiration.String(),
		Nonce:         signedOrder.Nonce.String(),
		FeeRateBps:    signedOrder.FeeRateBps.String(),
		Side:          strconv.FormatInt(int64(params.Side), 10),
		SignatureType: signatureTypeInt,
		Signature:     "0x" + hex.EncodeToString(signedOrder.Signature),
	}
	fmt.Println(ourSignedOrder)
	// Call CreateOrder to submit order
	return o.CreateOrder(ourSignedOrder, orderType, "")
}

// roundToTickSize rounds price to the specified tickSize
func roundToTickSize(price, tickSize float64) float64 {
	return float64(int64(price/tickSize+0.5)) * tickSize
}

// GetTickSizeResponse response for getting tickSize
type GetTickSizeResponse struct {
	MinimumTickSize float64 `json:"minimum_tick_size"`
}

// GetTickSize gets tickSize for the specified tokenID (with cache)
// Reference: https://github.com/Polymarket/clob-client
func (o *OrdersAPI) GetTickSize(tokenID string) (string, error) {
	// First check cache
	o.mu.RLock()
	if tickSize, exists := o.tickSizes[tokenID]; exists {
		o.mu.RUnlock()
		return tickSize, nil
	}
	o.mu.RUnlock()

	// Cache miss, fetch from API
	endpoint := "/tick-size"
	queryValues := url.Values{}
	queryValues.Set("token_id", tokenID)
	endpoint = endpoint + "?" + queryValues.Encode()

	data, err := o.client.Get(endpoint)
	if err != nil {
		return "", fmt.Errorf("get tick size: %w", err)
	}

	var response GetTickSizeResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return "", fmt.Errorf("unmarshal response: %w", err)
	}

	// Convert float64 to string for storage and return
	tickSizeStr := strconv.FormatFloat(response.MinimumTickSize, 'f', -1, 64)

	// Store result in cache
	o.mu.Lock()
	o.tickSizes[tokenID] = tickSizeStr
	o.mu.Unlock()

	return tickSizeStr, nil
}

// GetFeeRateBpsResponse response for getting feeRateBps
type GetFeeRateBpsResponse struct {
	BaseFee int `json:"base_fee"`
}

// GetFeeRateBps gets feeRateBps for the specified tokenID (with cache)
// Reference: https://github.com/Polymarket/clob-client
func (o *OrdersAPI) GetFeeRateBps(tokenID string) (int, error) {
	// First check cache
	o.mu.RLock()
	if feeRate, exists := o.feeRates[tokenID]; exists {
		o.mu.RUnlock()
		return feeRate, nil
	}
	o.mu.RUnlock()

	// Cache miss, fetch from API
	endpoint := "/fee-rate"
	queryValues := url.Values{}
	queryValues.Set("token_id", tokenID)
	endpoint = endpoint + "?" + queryValues.Encode()

	data, err := o.client.Get(endpoint)
	if err != nil {
		return 0, fmt.Errorf("get fee rate bps: %w", err)
	}

	var response GetFeeRateBpsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}

	// Store result in cache
	o.mu.Lock()
	o.feeRates[tokenID] = response.BaseFee
	o.mu.Unlock()

	return response.BaseFee, nil
}

// GetNegRiskResponse response for getting negRisk
type GetNegRiskResponse struct {
	NegRisk bool `json:"neg_risk"`
}

// GetNegRisk gets negRisk for the specified tokenID (with cache)
// Reference: https://github.com/Polymarket/clob-client
func (o *OrdersAPI) GetNegRisk(tokenID string) (bool, error) {
	// First check cache
	o.mu.RLock()
	if negRisk, exists := o.negRisks[tokenID]; exists {
		o.mu.RUnlock()
		return negRisk, nil
	}
	o.mu.RUnlock()

	// Cache miss, fetch from API
	endpoint := "/neg-risk"
	queryValues := url.Values{}
	queryValues.Set("token_id", tokenID)
	endpoint = endpoint + "?" + queryValues.Encode()

	data, err := o.client.Get(endpoint)
	if err != nil {
		return false, fmt.Errorf("get neg risk: %w", err)
	}

	var response GetNegRiskResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return false, fmt.Errorf("unmarshal response: %w", err)
	}

	// Store result in cache
	o.mu.Lock()
	o.negRisks[tokenID] = response.NegRisk
	o.mu.Unlock()

	return response.NegRisk, nil
}
