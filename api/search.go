package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mtt-labs/poly-market-sdk/client"
	"github.com/mtt-labs/poly-market-sdk/models"
)

// SearchAPI provides search-related API methods
// Uses Gamma API endpoint: https://gamma-api.polymarket.com
// Reference: https://docs.polymarket.com/api-reference/search/search-markets-events-and-profiles
type SearchAPI struct {
	gammaClient *client.GammaClient
}

// NewSearchAPI creates a new SearchAPI instance
func NewSearchAPI(c *client.Client) *SearchAPI {
	return &SearchAPI{
		gammaClient: client.NewGammaClient(),
	}
}

// Search searches for markets, events, and profiles
// Reference: https://docs.polymarket.com/api-reference/search/search-markets-events-and-profiles
func (s *SearchAPI) Search(params *models.SearchParams) (*models.SearchResponse, error) {
	if params == nil || params.Q == "" {
		return nil, fmt.Errorf("search query (q) is required")
	}

	endpoint := "/public-search"

	// Build query parameters
	queryValues := url.Values{}
	queryValues.Set("q", params.Q)

	if params.Cache != nil {
		queryValues.Set("cache", strconv.FormatBool(*params.Cache))
	}
	if params.EventsStatus != nil {
		queryValues.Set("events_status", *params.EventsStatus)
	}
	if params.LimitPerType != nil {
		queryValues.Set("limit_per_type", strconv.Itoa(*params.LimitPerType))
	}
	if params.Page != nil {
		queryValues.Set("page", strconv.Itoa(*params.Page))
	}
	if len(params.EventsTag) > 0 {
		for _, tag := range params.EventsTag {
			queryValues.Add("events_tag", tag)
		}
	}
	if params.KeepClosedMarkets != nil {
		queryValues.Set("keep_closed_markets", strconv.Itoa(*params.KeepClosedMarkets))
	}
	if params.Sort != nil {
		queryValues.Set("sort", *params.Sort)
	}
	if params.Ascending != nil {
		queryValues.Set("ascending", strconv.FormatBool(*params.Ascending))
	}
	if params.SearchTags != nil {
		queryValues.Set("search_tags", strconv.FormatBool(*params.SearchTags))
	}
	if params.SearchProfiles != nil {
		queryValues.Set("search_profiles", strconv.FormatBool(*params.SearchProfiles))
	}
	if params.Recurrence != nil {
		queryValues.Set("recurrence", *params.Recurrence)
	}
	if len(params.ExcludeTagID) > 0 {
		for _, tagID := range params.ExcludeTagID {
			queryValues.Add("exclude_tag_id", strconv.Itoa(tagID))
		}
	}
	if params.Optimized != nil {
		queryValues.Set("optimized", strconv.FormatBool(*params.Optimized))
	}

	endpoint = endpoint + "?" + queryValues.Encode()

	data, err := s.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("search: %w", err)
	}

	var response models.SearchResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &response, nil
}
