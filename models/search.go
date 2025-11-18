package models

import "time"

// SearchResponse represents the search response
// Reference: https://docs.polymarket.com/api-reference/search/search-markets-events-and-profiles
type SearchResponse struct {
	Events     []Event          `json:"events,omitempty"`
	Tags       []SearchTag      `json:"tags,omitempty"`
	Profiles   []Profile        `json:"profiles,omitempty"`
	Pagination SearchPagination `json:"pagination,omitempty"`
}

// SearchTag represents a tag in search results
type SearchTag struct {
	ID         string `json:"id,omitempty"`
	Label      string `json:"label,omitempty"`
	Slug       string `json:"slug,omitempty"`
	EventCount *int   `json:"event_count,omitempty"`
}

// Profile represents a user profile in search results
type Profile struct {
	ID                    string          `json:"id,omitempty"`
	Name                  string          `json:"name,omitempty"`
	User                  *int            `json:"user,omitempty"`
	Referral              string          `json:"referral,omitempty"`
	CreatedBy             *int            `json:"createdBy,omitempty"`
	UpdatedBy             *int            `json:"updatedBy,omitempty"`
	CreatedAt             *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time      `json:"updatedAt,omitempty"`
	UtmSource             string          `json:"utmSource,omitempty"`
	UtmMedium             string          `json:"utmMedium,omitempty"`
	UtmCampaign           string          `json:"utmCampaign,omitempty"`
	UtmContent            string          `json:"utmContent,omitempty"`
	UtmTerm               string          `json:"utmTerm,omitempty"`
	WalletActivated       *bool           `json:"walletActivated,omitempty"`
	Pseudonym             string          `json:"pseudonym,omitempty"`
	DisplayUsernamePublic *bool           `json:"displayUsernamePublic,omitempty"`
	ProfileImage          string          `json:"profileImage,omitempty"`
	Bio                   string          `json:"bio,omitempty"`
	ProxyWallet           string          `json:"proxyWallet,omitempty"`
	ProfileImageOptimized *OptimizedImage `json:"profileImageOptimized,omitempty"`
	IsCloseOnly           *bool           `json:"isCloseOnly,omitempty"`
	IsCertReq             *bool           `json:"isCertReq,omitempty"`
	CertReqDate           *time.Time      `json:"certReqDate,omitempty"`
}

// SearchPagination represents pagination information in search results
type SearchPagination struct {
	HasMore      *bool `json:"hasMore,omitempty"`
	TotalResults *int  `json:"totalResults,omitempty"`
}

// SearchParams parameters for search
// Reference: https://docs.polymarket.com/api-reference/search/search-markets-events-and-profiles
type SearchParams struct {
	Q                 string // required - search query
	Cache             *bool
	EventsStatus      *string
	LimitPerType      *int
	Page              *int
	EventsTag         []string
	KeepClosedMarkets *int
	Sort              *string
	Ascending         *bool
	SearchTags        *bool
	SearchProfiles    *bool
	Recurrence        *string
	ExcludeTagID      []int
	Optimized         *bool
}
