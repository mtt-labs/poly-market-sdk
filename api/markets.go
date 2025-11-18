package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/mtt-labs/poly-market-sdk/client"
	"github.com/mtt-labs/poly-market-sdk/models"
)

// MarketsAPI provides market-related API methods
// Uses Gamma API endpoint: https://gamma-api.polymarket.com
// Reference: https://docs.polymarket.com/developers/gamma-markets-api/overview
type MarketsAPI struct {
	gammaClient *client.GammaClient
}

// NewMarketsAPI creates a new MarketsAPI instance
func NewMarketsAPI(c *client.Client) *MarketsAPI {
	return &MarketsAPI{
		gammaClient: client.NewGammaClient(),
	}
}

// ListMarketsParams parameters for listing markets
// Reference: https://docs.polymarket.com/api-reference/markets/list-markets
type ListMarketsParams struct {
	Limit               *int    // Required range: x >= 0
	Offset              *int    // Required range: x >= 0
	Order               *string // Comma-separated list of fields to order by
	Ascending           *bool
	ID                  []int
	Slug                []string
	ClobTokenIDs        []string
	ConditionIDs        []string
	MarketMakerAddress  []string
	LiquidityNumMin     *float64
	LiquidityNumMax     *float64
	VolumeNumMin        *float64
	VolumeNumMax        *float64
	StartDateMin        *time.Time
	StartDateMax        *time.Time
	EndDateMin          *time.Time
	EndDateMax          *time.Time
	TagID               *int
	RelatedTags         *bool
	CYOM                *bool
	UmaResolutionStatus *string
	GameID              *string
	SportsMarketTypes   []string
	RewardsMinSize      *float64
	QuestionIDs         []string
	IncludeTag          *bool
	Closed              *bool
}

// GetMarkets gets a list of markets
// Returns an array of markets directly
// Reference: https://docs.polymarket.com/api-reference/markets/list-markets
func (m *MarketsAPI) GetMarkets(params *ListMarketsParams) ([]models.Market, error) {
	endpoint := "/markets"

	// Build query parameters
	queryValues := url.Values{}
	if params != nil {
		if params.Limit != nil {
			queryValues.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Offset != nil {
			queryValues.Set("offset", strconv.Itoa(*params.Offset))
		}
		if params.Order != nil {
			queryValues.Set("order", *params.Order)
		}
		if params.Ascending != nil {
			queryValues.Set("ascending", strconv.FormatBool(*params.Ascending))
		}
		if len(params.ID) > 0 {
			for _, id := range params.ID {
				queryValues.Add("id", strconv.Itoa(id))
			}
		}
		if len(params.Slug) > 0 {
			for _, slug := range params.Slug {
				queryValues.Add("slug", slug)
			}
		}
		if len(params.ClobTokenIDs) > 0 {
			for _, tokenID := range params.ClobTokenIDs {
				queryValues.Add("clob_token_ids", tokenID)
			}
		}
		if len(params.ConditionIDs) > 0 {
			for _, conditionID := range params.ConditionIDs {
				queryValues.Add("condition_ids", conditionID)
			}
		}
		if len(params.MarketMakerAddress) > 0 {
			for _, addr := range params.MarketMakerAddress {
				queryValues.Add("market_maker_address", addr)
			}
		}
		if params.LiquidityNumMin != nil {
			queryValues.Set("liquidity_num_min", strconv.FormatFloat(*params.LiquidityNumMin, 'f', -1, 64))
		}
		if params.LiquidityNumMax != nil {
			queryValues.Set("liquidity_num_max", strconv.FormatFloat(*params.LiquidityNumMax, 'f', -1, 64))
		}
		if params.VolumeNumMin != nil {
			queryValues.Set("volume_num_min", strconv.FormatFloat(*params.VolumeNumMin, 'f', -1, 64))
		}
		if params.VolumeNumMax != nil {
			queryValues.Set("volume_num_max", strconv.FormatFloat(*params.VolumeNumMax, 'f', -1, 64))
		}
		if params.StartDateMin != nil {
			queryValues.Set("start_date_min", params.StartDateMin.Format(time.RFC3339))
		}
		if params.StartDateMax != nil {
			queryValues.Set("start_date_max", params.StartDateMax.Format(time.RFC3339))
		}
		if params.EndDateMin != nil {
			queryValues.Set("end_date_min", params.EndDateMin.Format(time.RFC3339))
		}
		if params.EndDateMax != nil {
			queryValues.Set("end_date_max", params.EndDateMax.Format(time.RFC3339))
		}
		if params.TagID != nil {
			queryValues.Set("tag_id", strconv.Itoa(*params.TagID))
		}
		if params.RelatedTags != nil {
			queryValues.Set("related_tags", strconv.FormatBool(*params.RelatedTags))
		}
		if params.CYOM != nil {
			queryValues.Set("cyom", strconv.FormatBool(*params.CYOM))
		}
		if params.UmaResolutionStatus != nil {
			queryValues.Set("uma_resolution_status", *params.UmaResolutionStatus)
		}
		if params.GameID != nil {
			queryValues.Set("game_id", *params.GameID)
		}
		if len(params.SportsMarketTypes) > 0 {
			for _, marketType := range params.SportsMarketTypes {
				queryValues.Add("sports_market_types", marketType)
			}
		}
		if params.RewardsMinSize != nil {
			queryValues.Set("rewards_min_size", strconv.FormatFloat(*params.RewardsMinSize, 'f', -1, 64))
		}
		if len(params.QuestionIDs) > 0 {
			for _, questionID := range params.QuestionIDs {
				queryValues.Add("question_ids", questionID)
			}
		}
		if params.IncludeTag != nil {
			queryValues.Set("include_tag", strconv.FormatBool(*params.IncludeTag))
		}
		if params.Closed != nil {
			queryValues.Set("closed", strconv.FormatBool(*params.Closed))
		}
	}

	// Add query parameters to endpoint
	if len(queryValues) > 0 {
		endpoint = endpoint + "?" + queryValues.Encode()
	}

	data, err := m.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get markets: %w", err)
	}

	// Response is directly an array of markets
	var markets []models.Market
	if err := json.Unmarshal(data, &markets); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return markets, nil
}

// GetMarketByID gets market details by ID
// Reference: https://docs.polymarket.com/api-reference/markets/get-market-by-id
func (m *MarketsAPI) GetMarketByID(marketID string) (*models.Market, error) {
	endpoint := fmt.Sprintf("/markets/%s", marketID)

	data, err := m.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get market: %w", err)
	}

	// Response is directly a Market object, not wrapped
	var market models.Market
	if err := json.Unmarshal(data, &market); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &market, nil
}

// GetMarketBySlug gets market details by slug
// Reference: https://docs.polymarket.com/api-reference/markets/get-market-by-slug
func (m *MarketsAPI) GetMarketBySlug(slug string) (*models.Market, error) {
	endpoint := fmt.Sprintf("/markets/slug/%s", slug)

	data, err := m.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get market by slug: %w", err)
	}

	// Response is directly a Market object, not wrapped
	var market models.Market
	if err := json.Unmarshal(data, &market); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &market, nil
}

// GetMarketTags gets tags for a market by ID
// Reference: https://docs.polymarket.com/api-reference/markets/get-market-tags-by-id
func (m *MarketsAPI) GetMarketTags(marketID string) ([]models.EventTag, error) {
	endpoint := fmt.Sprintf("/markets/%s/tags", marketID)

	data, err := m.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get market tags: %w", err)
	}

	var tags []models.EventTag
	if err := json.Unmarshal(data, &tags); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return tags, nil
}

// GetMarketTrades gets market trade history
// Note: This might be in CLOB API or Data-API, not Gamma API
// Keeping for backward compatibility but may need to be moved to a different API
func (m *MarketsAPI) GetMarketTrades(marketID string, page, limit int) ([]models.Trade, error) {
	// This endpoint might not be in Gamma API
	// For now, return an error indicating this needs to be implemented differently
	return nil, fmt.Errorf("GetMarketTrades is not available in Gamma API, please use CLOB or Data-API")
}

// GetMarketOrderbook gets market orderbook
// Note: Orderbook is typically in CLOB API, not Gamma API
// Keeping for backward compatibility but may need to be moved to a different API
func (m *MarketsAPI) GetMarketOrderbook(marketID, outcomeID string) (*models.Orderbook, error) {
	// This endpoint is in CLOB API, not Gamma API
	// For now, return an error indicating this needs to be implemented differently
	return nil, fmt.Errorf("GetMarketOrderbook is not available in Gamma API, please use CLOB API")
}
