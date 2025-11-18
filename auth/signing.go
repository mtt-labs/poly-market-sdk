package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

const (
	// PolygonChainID Polygon mainnet Chain ID
	PolygonChainID = 137
	// ClobAuthDomainName EIP-712 Domain name
	ClobAuthDomainName = "ClobAuthDomain"
	// ClobAuthDomainVersion EIP-712 Domain version
	ClobAuthDomainVersion = "1"
	// ClobAuthMessage authentication message
	ClobAuthMessage = "This message attests that I control the given wallet"
)

// L1AuthHeaders L1 authentication headers
// Reference: https://docs.polymarket.com/developers/CLOB/authentication
type L1AuthHeaders struct {
	Address   string // POLY_ADDRESS: Polygon address
	Signature string // POLY_SIGNATURE: CLOB EIP 712 signature
	Timestamp string // POLY_TIMESTAMP: Current UNIX timestamp
	Nonce     string // POLY_NONCE: Nonce, default 0
}

// L2AuthHeaders L2 authentication headers
// Reference: https://docs.polymarket.com/developers/CLOB/authentication
type L2AuthHeaders struct {
	Address    string // POLY_ADDRESS: Polygon address
	Signature  string // POLY_SIGNATURE: HMAC signature
	Timestamp  string // POLY_TIMESTAMP: Current UNIX timestamp
	APIKey     string // POLY_API_KEY: Polymarket API key
	Passphrase string // POLY_PASSPHRASE: Polymarket API key passphrase
}

// Signer signer interface
type Signer interface {
	// SignL1Auth generates L1 authentication signature
	SignL1Auth(address string, timestamp int64, nonce *big.Int) (*L1AuthHeaders, error)
	// SignL2Auth generates L2 authentication signature (requires request method, path, body)
	SignL2Auth(address, method, path, body string, timestamp int64, apiKey, secret, passphrase string) (*L2AuthHeaders, error)
}

// PrivateKeySigner signer that uses private key for signing
type PrivateKeySigner struct {
	privateKey *ecdsa.PrivateKey
}

// NewPrivateKeySigner creates a private key signer
func NewPrivateKeySigner(privateKeyHex string) (*PrivateKeySigner, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	return &PrivateKeySigner{
		privateKey: privateKey,
	}, nil
}

// SignL1Auth generates L1 authentication signature (EIP-712)
// Reference: https://docs.polymarket.com/developers/CLOB/authentication
func (s *PrivateKeySigner) SignL1Auth(address string, timestamp int64, nonce *big.Int) (*L1AuthHeaders, error) {
	if nonce == nil {
		nonce = big.NewInt(0)
	}

	// Validate address format and convert to checksum format
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("invalid address format: %s", address)
	}
	// Convert to checksum format (EIP-55)
	address = common.HexToAddress(address).Hex()

	// Build EIP-712 Domain
	domain := apitypes.TypedDataDomain{
		Name:    ClobAuthDomainName,
		Version: ClobAuthDomainVersion,
		ChainId: math.NewHexOrDecimal256(PolygonChainID),
	}

	// Build EIP-712 Types
	types := apitypes.Types{
		"EIP712Domain": []apitypes.Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
		},
		"ClobAuth": []apitypes.Type{
			{Name: "address", Type: "address"},
			{Name: "timestamp", Type: "string"},
			{Name: "nonce", Type: "uint256"},
			{Name: "message", Type: "string"},
		},
	}

	// Build EIP-712 Message
	message := apitypes.TypedDataMessage{
		"address":   address,
		"timestamp": strconv.FormatInt(timestamp, 10),
		"nonce":     nonce.String(),
		"message":   ClobAuthMessage,
	}

	// Build TypedData
	typedData := apitypes.TypedData{
		Domain:      domain,
		Types:       types,
		PrimaryType: "ClobAuth",
		Message:     message,
	}

	// Sign using EIP-712
	// Calculate domain separator
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("hash domain: %w", err)
	}

	// Calculate typed data hash
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, fmt.Errorf("hash message: %w", err)
	}

	// EIP-712 final hash: keccak256("\x19\x01" || domainSeparator || typedDataHash)
	rawData := append([]byte("\x19\x01"), domainSeparator...)
	rawData = append(rawData, typedDataHash...)
	hash := crypto.Keccak256(rawData)

	// Sign
	signature, err := crypto.Sign(hash, s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("sign: %w", err)
	}

	// Ensure signature is 65 bytes (r + s + v)
	if len(signature) != 65 {
		return nil, fmt.Errorf("invalid signature length: %d", len(signature))
	}

	// crypto.Sign returns signature format: r (32 bytes) + s (32 bytes) + v (1 byte, 0 or 1)
	// Need to convert v to recovery ID (27 or 28)
	// If v is 0, recovery ID is 27; if v is 1, recovery ID is 28
	if signature[64] < 27 {
		signature[64] += 27
	}

	// Signature needs 0x prefix
	signatureHex := "0x" + hex.EncodeToString(signature)

	return &L1AuthHeaders{
		Address:   address, // Address is already in checksum format
		Signature: signatureHex,
		Timestamp: strconv.FormatInt(timestamp, 10),
		Nonce:     nonce.String(),
	}, nil
}

// SignL1AuthWithDefaults generates L1 authentication signature with default values
func (s *PrivateKeySigner) SignL1AuthWithDefaults(address string) (*L1AuthHeaders, error) {
	timestamp := time.Now().Unix()
	nonce := big.NewInt(0)
	return s.SignL1Auth(address, timestamp, nonce)
}

// decryptSecret decrypts secret using passphrase
// Reference: Polymarket clob-client implementation
func decryptSecret(encryptedSecret, passphrase string) (string, error) {
	// Decode base64 encoded secret
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedSecret)
	if err != nil {
		return "", fmt.Errorf("decode encrypted secret: %w", err)
	}

	// Use SHA256 of passphrase as key
	key := sha256.Sum256([]byte(passphrase))

	// Extract IV (first 16 bytes) and ciphertext
	if len(encryptedBytes) < 16 {
		return "", fmt.Errorf("encrypted secret too short")
	}
	iv := encryptedBytes[:16]
	ciphertext := encryptedBytes[16:]

	// Create AES cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", fmt.Errorf("create cipher: %w", err)
	}

	// Decrypt using CBC mode
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove PKCS7 padding
	padLen := int(plaintext[len(plaintext)-1])
	if padLen > len(plaintext) {
		return "", fmt.Errorf("invalid padding")
	}
	plaintext = plaintext[:len(plaintext)-padLen]

	return string(plaintext), nil
}

// toURLSafeBase64 converts base64 string to URL-safe base64
// Reference: Polymarket clob-client implementation
// Convert + to -, / to _, keep = suffix
func toURLSafeBase64(s string) string {
	result := make([]rune, 0, len(s))
	for _, r := range s {
		switch r {
		case '+':
			result = append(result, '-')
		case '/':
			result = append(result, '_')
		default:
			result = append(result, r)
		}
	}
	return string(result)
}

// toBase64 converts URL-safe base64 to base64
// Reference: Polymarket clob-client implementation
// Convert - to +, _ to /, keep = suffix
func toBase64(s string) string {
	result := make([]rune, 0, len(s))
	for _, r := range s {
		switch r {
		case '-':
			result = append(result, '+')
		case '_':
			result = append(result, '/')
		default:
			result = append(result, r)
		}
	}
	return string(result)
}

// SignL2Auth generates L2 authentication signature (HMAC)
// Reference: https://docs.polymarket.com/developers/CLOB/authentication
// Reference: Polymarket clob-client buildPolyHmacSignature implementation
func (s *PrivateKeySigner) SignL2Auth(address, method, path, body string, timestamp int64, apiKey, secret, passphrase string) (*L2AuthHeaders, error) {
	// Build message string to sign
	// Format: timestamp + method + path + (body if exists)
	message := fmt.Sprintf("%d%s%s", timestamp, method, path)
	if body != "" {
		message += body
	}

	// secret is base64 encoded, need to decode first
	secretBytes, err := base64.StdEncoding.DecodeString(toBase64(secret))
	if err != nil {
		return nil, fmt.Errorf("decode secret from base64: %w", err)
	}

	// Sign using HMAC-SHA256
	mac := hmac.New(sha256.New, secretBytes)
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)

	// Encode signature using base64
	sigBase64 := base64.StdEncoding.EncodeToString(signatureBytes)

	// Convert to URL-safe base64 (+ to -, / to _, keep =)
	signature := toURLSafeBase64(sigBase64)

	return &L2AuthHeaders{
		Address:    address,
		Signature:  signature,
		Timestamp:  strconv.FormatInt(timestamp, 10),
		APIKey:     apiKey,
		Passphrase: passphrase,
	}, nil
}

// SignL2AuthWithDefaults generates L2 authentication signature with current timestamp
func (s *PrivateKeySigner) SignL2AuthWithDefaults(address, method, path, body, apiKey, secret, passphrase string) (*L2AuthHeaders, error) {
	timestamp := time.Now().Unix()
	return s.SignL2Auth(address, method, path, body, timestamp, apiKey, secret, passphrase)
}

// GetAddressFromPrivateKey gets address from private key
func GetAddressFromPrivateKey(privateKeyHex string) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address.Hex(), nil
}
