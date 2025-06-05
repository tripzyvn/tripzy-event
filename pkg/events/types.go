package events

import (
	"time"

	"github.com/cloudevents/sdk-go/v2/event"
)

// CloudEvent is an alias for the CloudEvents SDK event type
type CloudEvent = event.Event

// PostData represents the data field of a post event
type PostData struct {
	PostID              string    `json:"post_id"`
	Source              string    `json:"source"`
	URL                 string    `json:"url"`
	Caption             string    `json:"caption"`
	Hashtags            []string  `json:"hashtags"`
	Likes               int       `json:"likes"`
	Comments            int       `json:"comments"`
	Views               int       `json:"views"`
	CreatedAt           time.Time `json:"created_at"`
	GeoHint             *GeoHint  `json:"geo_hint,omitempty"`
	GuessedLocationName string    `json:"guessed_location_name,omitempty"`
}

// GeoHint represents geographical coordinates
type GeoHint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// AnalysisData represents the data field of an analysis event
type AnalysisData struct {
	TaskID         string                 `json:"task_id"`
	DataSource     string                 `json:"data_source"`
	DataReference  string                 `json:"data_reference"`
	ProcessingType string                 `json:"processing_type"`
	Results        map[string]interface{} `json:"results"`
	CreatedAt      time.Time              `json:"created_at"`
}