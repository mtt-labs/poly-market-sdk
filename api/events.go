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

// EventsAPI provides event-related API methods
// Uses Gamma API endpoint: https://gamma-api.polymarket.com
// Reference: https://docs.polymarket.com/api-reference/events/list-events
type EventsAPI struct {
	gammaClient *client.GammaClient
}

// NewEventsAPI creates a new EventsAPI instance
func NewEventsAPI(c *client.Client) *EventsAPI {
	return &EventsAPI{
		gammaClient: client.NewGammaClient(),
	}
}

// ListEvents lists events with optional filters
// Reference: https://docs.polymarket.com/api-reference/events/list-events
func (e *EventsAPI) ListEvents(params *models.ListEventsParams) ([]models.Event, error) {
	endpoint := "/events"

	// Build query parameters
	queryValues := url.Values{}
	if params != nil {
		if params.Limit != 0 {
			queryValues.Set("limit", strconv.Itoa(params.Limit))
		}
		if params.Offset != 0 {
			queryValues.Set("offset", strconv.Itoa(params.Offset))
		}
		if params.Order != "" {
			queryValues.Set("order", params.Order)
		}
		if params.Ascending {
			queryValues.Set("ascending", strconv.FormatBool(params.Ascending))
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
		if params.TagID != 0 {
			queryValues.Set("tag_id", strconv.Itoa(params.TagID))
		}
		if len(params.ExcludeTagID) > 0 {
			for _, tagID := range params.ExcludeTagID {
				queryValues.Add("exclude_tag_id", strconv.Itoa(tagID))
			}
		}
		if params.RelatedTags {
			queryValues.Set("related_tags", strconv.FormatBool(params.RelatedTags))
		}
		if params.Featured {
			queryValues.Set("featured", strconv.FormatBool(params.Featured))
		}
		if params.CYOM {
			queryValues.Set("cyom", strconv.FormatBool(params.CYOM))
		}
		if params.IncludeChat {
			queryValues.Set("include_chat", strconv.FormatBool(params.IncludeChat))
		}
		if params.IncludeTemplate {
			queryValues.Set("include_template", strconv.FormatBool(params.IncludeTemplate))
		}
		if params.Closed {
			queryValues.Set("closed", strconv.FormatBool(params.Closed))
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
	}

	// Add query parameters to endpoint
	if len(queryValues) > 0 {
		endpoint = endpoint + "?" + queryValues.Encode()
	}

	data, err := e.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("list events: %w", err)
	}

	var events []models.Event
	if err := json.Unmarshal(data, &events); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return events, nil
}

// GetEventByID gets an event by ID
// Reference: https://docs.polymarket.com/api-reference/events/get-event-by-id
func (e *EventsAPI) GetEventByID(eventID string) (*models.Event, error) {
	endpoint := fmt.Sprintf("/events/%s", eventID)

	data, err := e.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get event by id: %w", err)
	}

	var event models.Event
	if err := json.Unmarshal(data, &event); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &event, nil
}

// GetEventBySlug gets an event by slug
// Reference: https://docs.polymarket.com/api-reference/events/get-event-by-slug
func (e *EventsAPI) GetEventBySlug(slug string) (*models.Event, error) {
	endpoint := fmt.Sprintf("/events/slug/%s", slug)

	data, err := e.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get event by slug: %w", err)
	}

	var event models.Event
	if err := json.Unmarshal(data, &event); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &event, nil
}

// GetEventTags gets tags related to an event
// Reference: https://docs.polymarket.com/api-reference/events/get-event-tags
func (e *EventsAPI) GetEventTags(tagSlug string) ([]models.EventTag, error) {
	endpoint := fmt.Sprintf("/events/tags/%s", tagSlug)

	data, err := e.gammaClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get event tags: %w", err)
	}

	var tags []models.EventTag
	if err := json.Unmarshal(data, &tags); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return tags, nil
}
