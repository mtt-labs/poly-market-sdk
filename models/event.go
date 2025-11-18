package models

import "time"

// Event represents an event in Polymarket
// Reference: https://docs.polymarket.com/api-reference/events/list-events
type Event struct {
	ID                           string          `json:"id,omitempty"`
	Ticker                       *string         `json:"ticker,omitempty"`
	Slug                         *string         `json:"slug,omitempty"`
	Title                        *string         `json:"title,omitempty"`
	Subtitle                     *string         `json:"subtitle,omitempty"`
	Description                  *string         `json:"description,omitempty"`
	ResolutionSource             *string         `json:"resolutionSource,omitempty"`
	StartDate                    *time.Time      `json:"startDate,omitempty"`
	CreationDate                 *time.Time      `json:"creationDate,omitempty"`
	EndDate                      *time.Time      `json:"endDate,omitempty"`
	Image                        *string         `json:"image,omitempty"`
	Icon                         *string         `json:"icon,omitempty"`
	Active                       *bool           `json:"active,omitempty"`
	Closed                       *bool           `json:"closed,omitempty"`
	Archived                     *bool           `json:"archived,omitempty"`
	New                          *bool           `json:"new,omitempty"`
	Featured                     *bool           `json:"featured,omitempty"`
	Restricted                   *bool           `json:"restricted,omitempty"`
	Liquidity                    *float64        `json:"liquidity,omitempty"`
	Volume                       *float64        `json:"volume,omitempty"`
	OpenInterest                 *float64        `json:"openInterest,omitempty"`
	SortBy                       *string         `json:"sortBy,omitempty"`
	Category                     *string         `json:"category,omitempty"`
	Subcategory                  *string         `json:"subcategory,omitempty"`
	IsTemplate                   *bool           `json:"isTemplate,omitempty"`
	TemplateVariables            *string         `json:"templateVariables,omitempty"`
	PublishedAt                  *string         `json:"published_at,omitempty"`
	CreatedBy                    *string         `json:"createdBy,omitempty"`
	UpdatedBy                    *string         `json:"updatedBy,omitempty"`
	CreatedAt                    *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt                    *time.Time      `json:"updatedAt,omitempty"`
	CommentsEnabled              *bool           `json:"commentsEnabled,omitempty"`
	Competitive                  *float64        `json:"competitive,omitempty"`
	Volume24hr                   *float64        `json:"volume24hr,omitempty"`
	Volume1wk                    *float64        `json:"volume1wk,omitempty"`
	Volume1mo                    *float64        `json:"volume1mo,omitempty"`
	Volume1yr                    *float64        `json:"volume1yr,omitempty"`
	FeaturedImage                *string         `json:"featuredImage,omitempty"`
	DisqusThread                 *string         `json:"disqusThread,omitempty"`
	ParentEvent                  *string         `json:"parentEvent,omitempty"`
	EnableOrderBook              *bool           `json:"enableOrderBook,omitempty"`
	LiquidityAmm                 *float64        `json:"liquidityAmm,omitempty"`
	LiquidityClob                *float64        `json:"liquidityClob,omitempty"`
	NegRisk                      *bool           `json:"negRisk,omitempty"`
	NegRiskMarketID              *string         `json:"negRiskMarketID,omitempty"`
	NegRiskFeeBips               *int            `json:"negRiskFeeBips,omitempty"`
	CommentCount                 *int            `json:"commentCount,omitempty"`
	ImageOptimized               *OptimizedImage `json:"imageOptimized,omitempty"`
	IconOptimized                *OptimizedImage `json:"iconOptimized,omitempty"`
	FeaturedImageOptimized       *OptimizedImage `json:"featuredImageOptimized,omitempty"`
	SubEvents                    []string        `json:"subEvents,omitempty"`
	Markets                      []EventMarket   `json:"markets,omitempty"`
	Tags                         []EventTag      `json:"tags,omitempty"`
	CYOM                         *bool           `json:"cyom,omitempty"`
	ClosedTime                   *time.Time      `json:"closedTime,omitempty"`
	ShowAllOutcomes              *bool           `json:"showAllOutcomes,omitempty"`
	ShowMarketImages             *bool           `json:"showMarketImages,omitempty"`
	AutomaticallyResolved        *bool           `json:"automaticallyResolved,omitempty"`
	EnableNegRisk                *bool           `json:"enableNegRisk,omitempty"`
	AutomaticallyActive          *bool           `json:"automaticallyActive,omitempty"`
	EventDate                    *string         `json:"eventDate,omitempty"`
	StartTime                    *time.Time      `json:"startTime,omitempty"`
	EventWeek                    *int            `json:"eventWeek,omitempty"`
	SeriesSlug                   *string         `json:"seriesSlug,omitempty"`
	Score                        *string         `json:"score,omitempty"`
	Elapsed                      *string         `json:"elapsed,omitempty"`
	Period                       *string         `json:"period,omitempty"`
	Live                         *bool           `json:"live,omitempty"`
	Ended                        *bool           `json:"ended,omitempty"`
	FinishedTimestamp            *time.Time      `json:"finishedTimestamp,omitempty"`
	GmpChartMode                 *string         `json:"gmpChartMode,omitempty"`
	EventCreators                []EventCreator  `json:"eventCreators,omitempty"`
	TweetCount                   *int            `json:"tweetCount,omitempty"`
	Chats                        []EventChat     `json:"chats,omitempty"`
	FeaturedOrder                *int            `json:"featuredOrder,omitempty"`
	EstimateValue                *bool           `json:"estimateValue,omitempty"`
	CantEstimate                 *bool           `json:"cantEstimate,omitempty"`
	EstimatedValue               *string         `json:"estimatedValue,omitempty"`
	Templates                    []EventTemplate `json:"templates,omitempty"`
	SpreadsMainLine              *float64        `json:"spreadsMainLine,omitempty"`
	TotalsMainLine               *float64        `json:"totalsMainLine,omitempty"`
	CarouselMap                  *string         `json:"carouselMap,omitempty"`
	PendingDeployment            *bool           `json:"pendingDeployment,omitempty"`
	Deploying                    *bool           `json:"deploying,omitempty"`
	DeployingTimestamp           *time.Time      `json:"deployingTimestamp,omitempty"`
	ScheduledDeploymentTimestamp *time.Time      `json:"scheduledDeploymentTimestamp,omitempty"`
	GameStatus                   *string         `json:"gameStatus,omitempty"`
}

// OptimizedImage represents an optimized image
type OptimizedImage struct {
	ID                        *string  `json:"id,omitempty"`
	ImageURLSource            *string  `json:"imageUrlSource,omitempty"`
	ImageURLOptimized         *string  `json:"imageUrlOptimized,omitempty"`
	ImageSizeKbSource         *float64 `json:"imageSizeKbSource,omitempty"`
	ImageSizeKbOptimized      *float64 `json:"imageSizeKbOptimized,omitempty"`
	ImageOptimizedComplete    *bool    `json:"imageOptimizedComplete,omitempty"`
	ImageOptimizedLastUpdated *string  `json:"imageOptimizedLastUpdated,omitempty"`
	RelID                     *int     `json:"relID,omitempty"`
	Field                     *string  `json:"field,omitempty"`
	Relname                   *string  `json:"relname,omitempty"`
}

// EventMarket represents a market within an event
type EventMarket struct {
	ID                    string     `json:"id,omitempty"`
	Question              string     `json:"question,omitempty"`
	ConditionID           string     `json:"conditionId,omitempty"`
	Slug                  string     `json:"slug,omitempty"`
	TwitterCardImage      string     `json:"twitterCardImage,omitempty"`
	ResolutionSource      string     `json:"resolutionSource,omitempty"`
	EndDate               *time.Time `json:"endDate,omitempty"`
	Category              string     `json:"category,omitempty"`
	AmmType               string     `json:"ammType,omitempty"`
	Liquidity             string     `json:"liquidity,omitempty"`
	SponsorName           string     `json:"sponsorName,omitempty"`
	SponsorImage          string     `json:"sponsorImage,omitempty"`
	StartDate             *time.Time `json:"startDate,omitempty"`
	XAxisValue            string     `json:"xAxisValue,omitempty"`
	YAxisValue            string     `json:"yAxisValue,omitempty"`
	DenominationToken     string     `json:"denominationToken,omitempty"`
	Fee                   string     `json:"fee,omitempty"`
	Image                 string     `json:"image,omitempty"`
	Icon                  string     `json:"icon,omitempty"`
	LowerBound            string     `json:"lowerBound,omitempty"`
	UpperBound            string     `json:"upperBound,omitempty"`
	Description           string     `json:"description,omitempty"`
	Outcomes              string     `json:"outcomes,omitempty"`
	OutcomePrices         string     `json:"outcomePrices,omitempty"`
	Volume                string     `json:"volume,omitempty"`
	Active                *bool      `json:"active,omitempty"`
	MarketType            string     `json:"marketType,omitempty"`
	FormatType            string     `json:"formatType,omitempty"`
	LowerBoundDate        string     `json:"lowerBoundDate,omitempty"`
	UpperBoundDate        string     `json:"upperBoundDate,omitempty"`
	Closed                *bool      `json:"closed,omitempty"`
	MarketMakerAddress    string     `json:"marketMakerAddress,omitempty"`
	CreatedBy             *int       `json:"createdBy,omitempty"`
	UpdatedBy             *int       `json:"updatedBy,omitempty"`
	CreatedAt             *time.Time `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time `json:"updatedAt,omitempty"`
	ClosedTime            string     `json:"closedTime,omitempty"`
	WideFormat            *bool      `json:"wideFormat,omitempty"`
	New                   *bool      `json:"new,omitempty"`
	MailchimpTag          string     `json:"mailchimpTag,omitempty"`
	Featured              *bool      `json:"featured,omitempty"`
	Archived              *bool      `json:"archived,omitempty"`
	ResolvedBy            string     `json:"resolvedBy,omitempty"`
	Restricted            *bool      `json:"restricted,omitempty"`
	MarketGroup           *int       `json:"marketGroup,omitempty"`
	GroupItemTitle        string     `json:"groupItemTitle,omitempty"`
	GroupItemThreshold    string     `json:"groupItemThreshold,omitempty"`
	QuestionID            string     `json:"questionID,omitempty"`
	UmaEndDate            string     `json:"umaEndDate,omitempty"`
	EnableOrderBook       *bool      `json:"enableOrderBook,omitempty"`
	OrderPriceMinTickSize *float64   `json:"orderPriceMinTickSize,omitempty"`
	OrderMinSize          *float64   `json:"orderMinSize,omitempty"`
	UmaResolutionStatus   string     `json:"umaResolutionStatus,omitempty"`
	CurationOrder         *int       `json:"curationOrder,omitempty"`
	VolumeNum             *float64   `json:"volumeNum,omitempty"`
	LiquidityNum          *float64   `json:"liquidityNum,omitempty"`
	EndDateIso            string     `json:"endDateIso,omitempty"`
	StartDateIso          string     `json:"startDateIso,omitempty"`
	UmaEndDateIso         string     `json:"umaEndDateIso,omitempty"`
	HasReviewedDates      *bool      `json:"hasReviewedDates,omitempty"`
	ReadyForCron          *bool      `json:"readyForCron,omitempty"`
	CommentsEnabled       *bool      `json:"commentsEnabled,omitempty"`
	Volume24hr            *float64   `json:"volume24hr,omitempty"`
	Volume1wk             *float64   `json:"volume1wk,omitempty"`
	Volume1mo             *float64   `json:"volume1mo,omitempty"`
	Volume1yr             *float64   `json:"volume1yr,omitempty"`
	GameStartTime         string     `json:"gameStartTime,omitempty"`
	SecondsDelay          *int       `json:"secondsDelay,omitempty"`
	ClobTokenIds          string     `json:"clobTokenIds,omitempty"`
	DisqusThread          string     `json:"disqusThread,omitempty"`
	ShortOutcomes         string     `json:"shortOutcomes,omitempty"`
	TeamAID               string     `json:"teamAID,omitempty"`
	TeamBID               string     `json:"teamBID,omitempty"`
	UmaBond               string     `json:"umaBond,omitempty"`
	UmaReward             string     `json:"umaReward,omitempty"`
	FpmmLive              *bool      `json:"fpmmLive,omitempty"`
	Volume24hrAmm         *float64   `json:"volume24hrAmm,omitempty"`
	Volume1wkAmm          *float64   `json:"volume1wkAmm,omitempty"`
	Volume1moAmm          *float64   `json:"volume1moAmm,omitempty"`
	Volume1yrAmm          *float64   `json:"volume1yrAmm,omitempty"`
}

// EventTag represents a tag associated with an event
type EventTag struct {
	ID          string     `json:"id,omitempty"`
	Label       string     `json:"label,omitempty"`
	Slug        string     `json:"slug,omitempty"`
	ForceShow   *bool      `json:"forceShow,omitempty"`
	PublishedAt *string    `json:"publishedAt,omitempty"`
	CreatedBy   *int       `json:"createdBy,omitempty"`
	UpdatedBy   *int       `json:"updatedBy,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	ForceHide   *bool      `json:"forceHide,omitempty"`
	IsCarousel  *bool      `json:"isCarousel,omitempty"`
}

// EventCreator represents an event creator
type EventCreator struct {
	ID            string     `json:"id,omitempty"`
	CreatorName   string     `json:"creatorName,omitempty"`
	CreatorHandle string     `json:"creatorHandle,omitempty"`
	CreatorURL    string     `json:"creatorUrl,omitempty"`
	CreatorImage  string     `json:"creatorImage,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}

// EventChat represents a chat associated with an event
type EventChat struct {
	ID           string     `json:"id,omitempty"`
	ChannelID    string     `json:"channelId,omitempty"`
	ChannelName  string     `json:"channelName,omitempty"`
	ChannelImage string     `json:"channelImage,omitempty"`
	Live         *bool      `json:"live,omitempty"`
	StartTime    *time.Time `json:"startTime,omitempty"`
	EndTime      *time.Time `json:"endTime,omitempty"`
}

// EventTemplate represents an event template
type EventTemplate struct {
	ID               string `json:"id,omitempty"`
	EventTitle       string `json:"eventTitle,omitempty"`
	EventSlug        string `json:"eventSlug,omitempty"`
	EventImage       string `json:"eventImage,omitempty"`
	MarketTitle      string `json:"marketTitle,omitempty"`
	Description      string `json:"description,omitempty"`
	ResolutionSource string `json:"resolutionSource,omitempty"`
	NegRisk          *bool  `json:"negRisk,omitempty"`
	SortBy           string `json:"sortBy,omitempty"`
	ShowMarketImages *bool  `json:"showMarketImages,omitempty"`
	SeriesSlug       string `json:"seriesSlug,omitempty"`
	Outcomes         string `json:"outcomes,omitempty"`
}

// ListEventsParams parameters for listing events
type ListEventsParams struct {
	Limit           *int    // Required range: x >= 0
	Offset          *int    // Required range: x >= 0
	Order           *string // Comma-separated list of fields to order by
	Ascending       *bool
	ID              []int
	Slug            []string
	TagID           *int
	ExcludeTagID    []int
	RelatedTags     *bool
	Featured        *bool
	CYOM            *bool
	IncludeChat     *bool
	IncludeTemplate *bool
	Recurrence      *string
	Closed          *bool
	StartDateMin    *time.Time
	StartDateMax    *time.Time
	EndDateMin      *time.Time
	EndDateMax      *time.Time
}
