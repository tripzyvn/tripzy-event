package events

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// CloudEvent represents a CloudEvents v1.0 event
type CloudEvent struct {
	SpecVersion     string          `json:"specversion"`
	Type            string          `json:"type"`
	Source          string          `json:"source"`
	ID              string          `json:"id"`
	Time            time.Time       `json:"time"`
	DataContentType string          `json:"datacontenttype"`
	Data            json.RawMessage `json:"data"`
}

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

// CloudEventConsumer defines the interface for consuming CloudEvents
type CloudEventConsumer interface {
	// ConsumePostEvent processes a post event
	ConsumePostEvent(event *CloudEvent) (*PostData, error)
}

// CloudEventProducer defines the interface for producing CloudEvents
type CloudEventProducer interface {
	// ProduceAnalysisEvent creates and sends an analysis event
	ProduceAnalysisEvent(analysisData *AnalysisData) error
}

// defaultCloudEventHandler is the default implementation of CloudEventConsumer and CloudEventProducer
type defaultCloudEventHandler struct {
	// In a real implementation, this would include clients for receiving and sending events
}

// NewCloudEventHandler creates a new CloudEventConsumer and CloudEventProducer
func NewCloudEventHandler() (CloudEventConsumer, CloudEventProducer) {
	handler := &defaultCloudEventHandler{}
	return handler, handler
}

// ConsumePostEvent processes a post event
func (h *defaultCloudEventHandler) ConsumePostEvent(event *CloudEvent) (*PostData, error) {
	// Check if the event is a post event
	if event.Type != "tripzy.raw.post.created" {
		return nil, nil
	}

	// Unmarshal the event data to PostData
	var postData PostData
	if err := json.Unmarshal(event.Data, &postData); err != nil {
		return nil, err
	}

	// In a real implementation, this would process the post data
	// For now, we'll just return it
	return &postData, nil
}

// ProduceAnalysisEvent creates and sends an analysis event
func (h *defaultCloudEventHandler) ProduceAnalysisEvent(analysisData *AnalysisData) error {
	// Create the CloudEvent
	event := CloudEvent{
		SpecVersion:     "1.0",
		Type:            "tripzy.analysis.completed",
		Source:          "tripzy-engine",
		ID:              uuid.New().String(),
		Time:            time.Now().UTC(),
		DataContentType: "application/json",
	}

	// Marshal the analysis data to JSON
	data, err := json.Marshal(analysisData)
	if err != nil {
		return err
	}
	event.Data = data

	// In a real implementation, this would send the event to a message broker or event bus
	// For now, we'll just log it
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}
	
	// In a real implementation, this would be replaced with actual sending logic
	// log.Printf("Produced CloudEvent: %s", string(eventJSON))
	
	// For demonstration purposes, we'll just return the JSON
	_ = eventJSON
	
	return nil
}