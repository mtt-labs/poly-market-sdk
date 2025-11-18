package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/mtt-labs/poly-market-sdk/client"
)

// AuthAPI provides authentication-related API methods
type AuthAPI struct {
	client *client.Client
}

// NewAuthAPI creates a new AuthAPI instance
func NewAuthAPI(c *client.Client) *AuthAPI {
	return &AuthAPI{client: c}
}

// APICredentials API credentials
type APICredentials struct {
	Key        string `json:"key"`
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
}

// CreateAPICredentialsRequest create API credentials request
type CreateAPICredentialsRequest struct {
	Address   string `json:"address"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Signature string `json:"signature"`
}

// CreateAPICredentialsResponse create API credentials response
type CreateAPICredentialsResponse struct {
	Key        string `json:"apiKey"`
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
}

// CreateOrDeriveAPICredentials creates or derives API credentials
// Reference: https://docs.polymarket.com/developers/CLOB/authentication
// If API key already exists, derive it; otherwise create a new one
func (a *AuthAPI) CreateOrDeriveAPICredentials() (*APICredentials, error) {
	// First try to derive (if exists)
	//creds, err := a.DeriveAPICredentials()
	//if err == nil && creds != nil {
	//	return creds, nil
	//}

	// If derivation fails, create new
	return a.CreateAPICredentials()
}

// CreateAPICredentials creates new API credentials
// This endpoint requires L1 Header
// Reference: https://docs.polymarket.com/developers/CLOB/authentication
func (a *AuthAPI) CreateAPICredentials() (*APICredentials, error) {
	endpoint := "/auth/api-key"

	// Generate L1 authentication signature
	address := a.client.GetAddress()
	timestamp := time.Now().Unix()
	nonce := big.NewInt(0)

	signer := a.client.GetSigner()
	l1Headers, err := signer.SignL1Auth(address, timestamp, nonce)
	if err != nil {
		return nil, fmt.Errorf("sign L1 auth: %w", err)
	}

	// Build request body (although docs say these fields are needed, actually only L1 Header may be required)
	req := &CreateAPICredentialsRequest{
		Address:   address,
		Timestamp: l1Headers.Timestamp,
		Nonce:     l1Headers.Nonce,
		Signature: l1Headers.Signature,
	}

	// Send request with L1 Header
	data, err := a.client.PostWithL1(endpoint, req, l1Headers)
	if err != nil {
		return nil, fmt.Errorf("create API credentials: %w", err)
	}

	var response CreateAPICredentialsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	creds := &APICredentials{
		Key:        response.Key,
		Secret:     response.Secret,
		Passphrase: response.Passphrase,
	}

	// Set to client
	a.client.SetAPICredentials(creds.Key, creds.Secret, creds.Passphrase)

	return creds, nil
}

// DeriveAPICredentials derives API credentials
// This endpoint requires L1 Header
// Reference: https://docs.polymarket.com/developers/CLOB/authentication
func (a *AuthAPI) DeriveAPICredentials() (*APICredentials, error) {
	endpoint := "/auth/derive-api-key"

	// Generate L1 authentication signature
	address := a.client.GetAddress()
	timestamp := time.Now().Unix()
	nonce := big.NewInt(0)

	signer := a.client.GetSigner()
	l1Headers, err := signer.SignL1Auth(address, timestamp, nonce)
	if err != nil {
		return nil, fmt.Errorf("sign L1 auth: %w", err)
	}

	// Send request with L1 Header
	data, err := a.client.GetWithL1(endpoint, l1Headers)
	if err != nil {
		return nil, fmt.Errorf("derive API credentials: %w", err)
	}

	var response CreateAPICredentialsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	creds := &APICredentials{
		Key:        response.Key,
		Secret:     response.Secret,
		Passphrase: response.Passphrase,
	}

	// Set to client
	a.client.SetAPICredentials(creds.Key, creds.Secret, creds.Passphrase)

	return creds, nil
}
