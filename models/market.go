package models

import "time"

// Market represents a prediction market
// Reference: https://docs.polymarket.com/api-reference/markets/list-markets
type Market struct {
	ID                           string          `json:"id,omitempty"`
	Question                     *string         `json:"question,omitempty"`
	ConditionID                  string          `json:"conditionId,omitempty"`
	Slug                         *string         `json:"slug,omitempty"`
	TwitterCardImage             *string         `json:"twitterCardImage,omitempty"`
	ResolutionSource             *string         `json:"resolutionSource,omitempty"`
	EndDate                      *time.Time      `json:"endDate,omitempty"`
	Category                     *string         `json:"category,omitempty"`
	AmmType                      *string         `json:"ammType,omitempty"`
	Liquidity                    *string         `json:"liquidity,omitempty"`
	SponsorName                  *string         `json:"sponsorName,omitempty"`
	SponsorImage                 *string         `json:"sponsorImage,omitempty"`
	StartDate                    *time.Time      `json:"startDate,omitempty"`
	XAxisValue                   *string         `json:"xAxisValue,omitempty"`
	YAxisValue                   *string         `json:"yAxisValue,omitempty"`
	DenominationToken            *string         `json:"denominationToken,omitempty"`
	Fee                          *string         `json:"fee,omitempty"`
	Image                        *string         `json:"image,omitempty"`
	Icon                         *string         `json:"icon,omitempty"`
	LowerBound                   *string         `json:"lowerBound,omitempty"`
	UpperBound                   *string         `json:"upperBound,omitempty"`
	Description                  *string         `json:"description,omitempty"`
	Outcomes                     *string         `json:"outcomes,omitempty"`
	OutcomePrices                *string         `json:"outcomePrices,omitempty"`
	Volume                       *string         `json:"volume,omitempty"`
	Active                       *bool           `json:"active,omitempty"`
	MarketType                   *string         `json:"marketType,omitempty"`
	FormatType                   *string         `json:"formatType,omitempty"`
	LowerBoundDate               *string         `json:"lowerBoundDate,omitempty"`
	UpperBoundDate               *string         `json:"upperBoundDate,omitempty"`
	Closed                       *bool           `json:"closed,omitempty"`
	MarketMakerAddress           string          `json:"marketMakerAddress,omitempty"`
	CreatedBy                    *int            `json:"createdBy,omitempty"`
	UpdatedBy                    *int            `json:"updatedBy,omitempty"`
	CreatedAt                    *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt                    *time.Time      `json:"updatedAt,omitempty"`
	ClosedTime                   *string         `json:"closedTime,omitempty"`
	WideFormat                   *bool           `json:"wideFormat,omitempty"`
	New                          *bool           `json:"new,omitempty"`
	MailchimpTag                 *string         `json:"mailchimpTag,omitempty"`
	Featured                     *bool           `json:"featured,omitempty"`
	Archived                     *bool           `json:"archived,omitempty"`
	ResolvedBy                   *string         `json:"resolvedBy,omitempty"`
	Restricted                   *bool           `json:"restricted,omitempty"`
	MarketGroup                  *int            `json:"marketGroup,omitempty"`
	GroupItemTitle               *string         `json:"groupItemTitle,omitempty"`
	GroupItemThreshold           *string         `json:"groupItemThreshold,omitempty"`
	QuestionID                   *string         `json:"questionID,omitempty"`
	UmaEndDate                   *string         `json:"umaEndDate,omitempty"`
	EnableOrderBook              *bool           `json:"enableOrderBook,omitempty"`
	OrderPriceMinTickSize        *float64        `json:"orderPriceMinTickSize,omitempty"`
	OrderMinSize                 *float64        `json:"orderMinSize,omitempty"`
	UmaResolutionStatus          *string         `json:"umaResolutionStatus,omitempty"`
	CurationOrder                *int            `json:"curationOrder,omitempty"`
	VolumeNum                    *float64        `json:"volumeNum,omitempty"`
	LiquidityNum                 *float64        `json:"liquidityNum,omitempty"`
	EndDateIso                   *string         `json:"endDateIso,omitempty"`
	StartDateIso                 *string         `json:"startDateIso,omitempty"`
	UmaEndDateIso                *string         `json:"umaEndDateIso,omitempty"`
	HasReviewedDates             *bool           `json:"hasReviewedDates,omitempty"`
	ReadyForCron                 *bool           `json:"readyForCron,omitempty"`
	CommentsEnabled              *bool           `json:"commentsEnabled,omitempty"`
	Volume24hr                   *float64        `json:"volume24hr,omitempty"`
	Volume1wk                    *float64        `json:"volume1wk,omitempty"`
	Volume1mo                    *float64        `json:"volume1mo,omitempty"`
	Volume1yr                    *float64        `json:"volume1yr,omitempty"`
	GameStartTime                *string         `json:"gameStartTime,omitempty"`
	SecondsDelay                 *int            `json:"secondsDelay,omitempty"`
	ClobTokenIds                 *string         `json:"clobTokenIds,omitempty"`
	DisqusThread                 *string         `json:"disqusThread,omitempty"`
	ShortOutcomes                *string         `json:"shortOutcomes,omitempty"`
	TeamAID                      *string         `json:"teamAID,omitempty"`
	TeamBID                      *string         `json:"teamBID,omitempty"`
	UmaBond                      *string         `json:"umaBond,omitempty"`
	UmaReward                    *string         `json:"umaReward,omitempty"`
	FpmmLive                     *bool           `json:"fpmmLive,omitempty"`
	Volume24hrAmm                *float64        `json:"volume24hrAmm,omitempty"`
	Volume1wkAmm                 *float64        `json:"volume1wkAmm,omitempty"`
	Volume1moAmm                 *float64        `json:"volume1moAmm,omitempty"`
	Volume1yrAmm                 *float64        `json:"volume1yrAmm,omitempty"`
	Volume24hrClob               *float64        `json:"volume24hrClob,omitempty"`
	Volume1wkClob                *float64        `json:"volume1wkClob,omitempty"`
	Volume1moClob                *float64        `json:"volume1moClob,omitempty"`
	Volume1yrClob                *float64        `json:"volume1yrClob,omitempty"`
	VolumeAmm                    *float64        `json:"volumeAmm,omitempty"`
	VolumeClob                   *float64        `json:"volumeClob,omitempty"`
	LiquidityAmm                 *float64        `json:"liquidityAmm,omitempty"`
	LiquidityClob                *float64        `json:"liquidityClob,omitempty"`
	MakerBaseFee                 *int            `json:"makerBaseFee,omitempty"`
	TakerBaseFee                 *int            `json:"takerBaseFee,omitempty"`
	CustomLiveness               *int            `json:"customLiveness,omitempty"`
	AcceptingOrders              *bool           `json:"acceptingOrders,omitempty"`
	NotificationsEnabled         *bool           `json:"notificationsEnabled,omitempty"`
	Score                        *int            `json:"score,omitempty"`
	ImageOptimized               *OptimizedImage `json:"imageOptimized,omitempty"`
	IconOptimized                *OptimizedImage `json:"iconOptimized,omitempty"`
	Events                       []Event         `json:"events,omitempty"`
	Categories                   []interface{}   `json:"categories,omitempty"` // Categories structure can vary
	Tags                         []EventTag      `json:"tags,omitempty"`
	Creator                      *string         `json:"creator,omitempty"`
	Ready                        *bool           `json:"ready,omitempty"`
	Funded                       *bool           `json:"funded,omitempty"`
	PastSlugs                    *string         `json:"pastSlugs,omitempty"`
	ReadyTimestamp               *time.Time      `json:"readyTimestamp,omitempty"`
	FundedTimestamp              *time.Time      `json:"fundedTimestamp,omitempty"`
	AcceptingOrdersTimestamp     *time.Time      `json:"acceptingOrdersTimestamp,omitempty"`
	Competitive                  *float64        `json:"competitive,omitempty"`
	RewardsMinSize               *float64        `json:"rewardsMinSize,omitempty"`
	RewardsMaxSpread             *float64        `json:"rewardsMaxSpread,omitempty"`
	Spread                       *float64        `json:"spread,omitempty"`
	AutomaticallyResolved        *bool           `json:"automaticallyResolved,omitempty"`
	OneDayPriceChange            *float64        `json:"oneDayPriceChange,omitempty"`
	OneHourPriceChange           *float64        `json:"oneHourPriceChange,omitempty"`
	OneWeekPriceChange           *float64        `json:"oneWeekPriceChange,omitempty"`
	OneMonthPriceChange          *float64        `json:"oneMonthPriceChange,omitempty"`
	OneYearPriceChange           *float64        `json:"oneYearPriceChange,omitempty"`
	LastTradePrice               *float64        `json:"lastTradePrice,omitempty"`
	BestBid                      *float64        `json:"bestBid,omitempty"`
	BestAsk                      *float64        `json:"bestAsk,omitempty"`
	AutomaticallyActive          *bool           `json:"automaticallyActive,omitempty"`
	ClearBookOnStart             *bool           `json:"clearBookOnStart,omitempty"`
	ChartColor                   *string         `json:"chartColor,omitempty"`
	SeriesColor                  *string         `json:"seriesColor,omitempty"`
	ShowGmpSeries                *bool           `json:"showGmpSeries,omitempty"`
	ShowGmpOutcome               *bool           `json:"showGmpOutcome,omitempty"`
	ManualActivation             *bool           `json:"manualActivation,omitempty"`
	NegRiskOther                 *bool           `json:"negRiskOther,omitempty"`
	GameID                       *string         `json:"gameId,omitempty"`
	GroupItemRange               *string         `json:"groupItemRange,omitempty"`
	SportsMarketType             *string         `json:"sportsMarketType,omitempty"`
	Line                         *float64        `json:"line,omitempty"`
	UmaResolutionStatuses        *string         `json:"umaResolutionStatuses,omitempty"`
	PendingDeployment            *bool           `json:"pendingDeployment,omitempty"`
	Deploying                    *bool           `json:"deploying,omitempty"`
	DeployingTimestamp           *time.Time      `json:"deployingTimestamp,omitempty"`
	ScheduledDeploymentTimestamp *time.Time      `json:"scheduledDeploymentTimestamp,omitempty"`
	RfqEnabled                   *bool           `json:"rfqEnabled,omitempty"`
	EventStartTime               *time.Time      `json:"eventStartTime,omitempty"`
}

// Outcome represents a market outcome option
type Outcome struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Volume    string `json:"volume"`
	Liquidity string `json:"liquidity"`
	OutcomeID string `json:"outcome_id"`
	MarketID  string `json:"market_id"`
}

// MarketListResponse market list response
type MarketListResponse struct {
	Markets []Market `json:"markets"`
	Total   int      `json:"total"`
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
}

// MarketDetailResponse market detail response
type MarketDetailResponse struct {
	Market Market `json:"market"`
}
