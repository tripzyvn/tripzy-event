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

// LocationSuggestionData represents the data field of a location suggestion event
type LocationSuggestionData struct {
	GuessedName     string    `json:"guessed_name"`
	SuggestedBy     string    `json:"suggested_by"`
	OriginalCaption string    `json:"original_caption"`
	SourcePostURL   string    `json:"source_post_url"`
	Hashtags        []string  `json:"hashtags"`
	EstimatedRegion string    `json:"estimated_region"`
	GeoHint         *GeoHint  `json:"geo_hint,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	Confidence      float64   `json:"confidence"`
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

// TrendingScoreData represents the data field of a trending score event
type TrendingScoreData struct {
	LocationID      string    `json:"location_id"`
	TrendWindow     string    `json:"trend_window"`
	MentionCount    int       `json:"mention_count"`
	EngagementScore int       `json:"engagement_score"`
	TrendingScore   float64   `json:"trending_score"`
	TopHashtags     []string  `json:"top_hashtags"`
	SentimentScore  float64   `json:"sentiment_score"`
	RankInRegion    int       `json:"rank_in_region"`
	Confidence      float64   `json:"confidence"`
	LastMentionedAt time.Time `json:"last_mentioned_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
