package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mtt-labs/poly-market-sdk/auth"
)

// SignatureType signature type
// Reference: https://docs.polymarket.com/quickstart/orders/first-order
type SignatureType int

const (
	// SignatureTypeEOA ECDSA EIP712 signatures signed by EOAs
	SignatureTypeEOA SignatureType = 0
	// SignatureTypeEmailMagic EIP712 signatures signed by EOAs that own Polymarket Proxy wallets
	SignatureTypeEmailMagic SignatureType = 1
	// SignatureTypeBrowserWallet EIP712 signatures signed by EOAs that own Polymarket Gnosis safes
	SignatureTypeBrowserWallet SignatureType = 2
)

// Client is the main client of Polymarket SDK
type Client struct {
	baseURL       string
	httpClient    *http.Client
	apiKey        string        // API key (obtained via create_or_derive_api_creds)
	apiSecret     string        // API secret
	apiPassphrase string        // API passphrase
	privateKey    string        // Private key (for signing)
	chainID       int           // Chain ID, default 137 (Polygon)
	signatureType SignatureType // Signature type
	funder        string        // Proxy address (based on login method)
	signer        auth.Signer   // Signer
	address       string        // Address derived from private key
}

// Config is the client configuration
// Reference: https://docs.polymarket.com/quickstart/orders/first-order
type Config struct {
	BaseURL       string        // API base URL, default "https://clob.polymarket.com"
	PrivateKey    string        // Private key (required)
	ChainID       int           // Chain ID, default 137 (Polygon)
	SignatureType SignatureType // Signature type (0=EOA, 1=Email/Magic, 2=Browser Wallet)
	Funder        string        // Proxy address (based on login method, optional)
	APIKey        string        // API key (optional, can be obtained via create_or_derive_api_creds)
	APISecret     string        // API secret (optional)
	APIPassphrase string        // API passphrase (optional)
	Timeout       time.Duration
	HTTPClient    *http.Client
}

// NewClient creates a new Polymarket client
// Reference: https://docs.polymarket.com/quickstart/orders/first-order
func NewClient(config *Config) (*Client, error) {
	if config == nil {
		return nil, fmt.Errorf("config is required")
	}

	if config.PrivateKey == "" {
		return nil, fmt.Errorf("private key is required")
	}

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://clob.polymarket.com" // Polymarket default API address
	}

	chainID := config.ChainID
	if chainID == 0 {
		chainID = 137 // Polygon mainnet
	}

	timeout := config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: timeout,
		}
	}

	// Create signer
	signer, err := auth.NewPrivateKeySigner(config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("create signer: %w", err)
	}

	// Get address from private key
	address, err := auth.GetAddressFromPrivateKey(config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("get address from private key: %w", err)
	}

	return &Client{
		baseURL:       baseURL,
		httpClient:    httpClient,
		privateKey:    config.PrivateKey,
		chainID:       chainID,
		signatureType: config.SignatureType,
		funder:        config.Funder,
		apiKey:        config.APIKey,
		apiSecret:     config.APISecret,
		apiPassphrase: config.APIPassphrase,
		signer:        signer,
		address:       address,
	}, nil
}

// doRequest executes HTTP request
func (c *Client) doRequest(method, endpoint string, body interface{}, l1Headers *auth.L1AuthHeaders, l2Headers *auth.L2AuthHeaders) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := c.baseURL + endpoint
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Add L1 Header (if needed)
	// Reference: https://docs.polymarket.com/developers/CLOB/authentication
	if l1Headers != nil {
		req.Header.Set("POLY_ADDRESS", l1Headers.Address)
		req.Header.Set("POLY_SIGNATURE", l1Headers.Signature)
		req.Header.Set("POLY_TIMESTAMP", l1Headers.Timestamp)
		req.Header.Set("POLY_NONCE", l1Headers.Nonce)
	}

	// Add L2 Header (if needed)
	// Reference: https://docs.polymarket.com/developers/CLOB/authentication
	if l2Headers != nil {
		req.Header.Set("POLY_ADDRESS", l2Headers.Address)
		req.Header.Set("POLY_SIGNATURE", l2Headers.Signature)
		req.Header.Set("POLY_TIMESTAMP", l2Headers.Timestamp)
		req.Header.Set("POLY_API_KEY", l2Headers.APIKey)
		req.Header.Set("POLY_PASSPHRASE", l2Headers.Passphrase)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// Get executes GET request
func (c *Client) Get(endpoint string) ([]byte, error) {
	return c.doRequest(http.MethodGet, endpoint, nil, nil, nil)
}

// GetWithL1 executes GET request (with L1 Header)
func (c *Client) GetWithL1(endpoint string, l1Headers *auth.L1AuthHeaders) ([]byte, error) {
	return c.doRequest(http.MethodGet, endpoint, nil, l1Headers, nil)
}

// GetWithL2 executes GET request (with L2 Header)
func (c *Client) GetWithL2(endpoint string, l2Headers *auth.L2AuthHeaders) ([]byte, error) {
	return c.doRequest(http.MethodGet, endpoint, nil, nil, l2Headers)
}

// Post executes POST request
func (c *Client) Post(endpoint string, body interface{}) ([]byte, error) {
	return c.doRequest(http.MethodPost, endpoint, body, nil, nil)
}

// PostWithL1 executes POST request (with L1 Header)
func (c *Client) PostWithL1(endpoint string, body interface{}, l1Headers *auth.L1AuthHeaders) ([]byte, error) {
	return c.doRequest(http.MethodPost, endpoint, body, l1Headers, nil)
}

// PostWithL2 executes POST request (with L2 Header)
func (c *Client) PostWithL2(endpoint string, body interface{}, l2Headers *auth.L2AuthHeaders) ([]byte, error) {
	return c.doRequest(http.MethodPost, endpoint, body, nil, l2Headers)
}

// DeleteWithL2 executes DELETE request (with L2 Header)
func (c *Client) DeleteWithL2(endpoint string, body interface{}, l2Headers *auth.L2AuthHeaders) ([]byte, error) {
	return c.doRequest(http.MethodDelete, endpoint, body, nil, l2Headers)
}

// GetAPIKey gets API key from client config
func (c *Client) GetAPIKey() string {
	return c.apiKey
}

// GetAddress gets client address
func (c *Client) GetAddress() string {
	return c.address
}

// GetSigner gets signer
func (c *Client) GetSigner() auth.Signer {
	return c.signer
}

// SetAPICredentials sets API credentials
// Use after calling create_or_derive_api_creds
func (c *Client) SetAPICredentials(key, secret, passphrase string) {
	c.apiKey = key
	c.apiSecret = secret
	c.apiPassphrase = passphrase
}

// GetAPISecret gets API secret
func (c *Client) GetAPISecret() string {
	return c.apiSecret
}

// GetAPIPassphrase gets API passphrase
func (c *Client) GetAPIPassphrase() string {
	return c.apiPassphrase
}

// GetPrivateKey gets private key
func (c *Client) GetPrivateKey() string {
	return c.privateKey
}

// GetChainID gets Chain ID
func (c *Client) GetChainID() int {
	return c.chainID
}

// GetSignatureType gets signature type
func (c *Client) GetSignatureType() SignatureType {
	return c.signatureType
}

// GetFunder gets funder address
func (c *Client) GetFunder() string {
	return c.funder
}
