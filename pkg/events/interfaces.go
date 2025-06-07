package events

// PostEventConsumer defines the interface for consuming post events
type PostEventConsumer interface {
	// ConsumePostEvent processes a post event
	ConsumePostEvent(event *CloudEvent) (*PostData, error)
}

// LocationSuggestionEventConsumer defines the interface for consuming location suggestion events
type LocationSuggestionEventConsumer interface {
	// ConsumeLocationSuggestionEvent processes a location suggestion event
	ConsumeLocationSuggestionEvent(event *CloudEvent) (*LocationSuggestionData, error)
}

// TrendingScoreEventConsumer defines the interface for consuming trending score events
type TrendingScoreEventConsumer interface {
	// ConsumeTrendingScoreEvent processes a trending score event
	ConsumeTrendingScoreEvent(event *CloudEvent) (*TrendingScoreData, error)
}

// AnalysisEventProducer defines the interface for producing analysis events
type AnalysisEventProducer interface {
	// ProduceAnalysisEvent creates and sends an analysis event
	ProduceAnalysisEvent(analysisData *AnalysisData) error
}

// PostEventProducer defines the interface for producing post events
type PostEventProducer interface {
	// ProducePostEvent creates and sends a post event
	ProducePostEvent(postData *PostData) error
}

// LocationSuggestionEventProducer defines the interface for producing location suggestion events
type LocationSuggestionEventProducer interface {
	// ProduceLocationSuggestionEvent creates and sends a location suggestion event
	ProduceLocationSuggestionEvent(suggestionData *LocationSuggestionData) error
}

// TrendingScoreEventProducer defines the interface for producing trending score events
type TrendingScoreEventProducer interface {
	// ProduceTrendingScoreEvent creates and sends a trending score event
	ProduceTrendingScoreEvent(trendingScoreData *TrendingScoreData) error
}

// CloudEventConsumer defines the interface for consuming CloudEvents
// It embeds all consumer interfaces for backward compatibility
type CloudEventConsumer interface {
	PostEventConsumer
	LocationSuggestionEventConsumer
	TrendingScoreEventConsumer
}

// CloudEventProducer defines the interface for producing CloudEvents
// It embeds all producer interfaces for backward compatibility
type CloudEventProducer interface {
	AnalysisEventProducer
	PostEventProducer
	LocationSuggestionEventProducer
	TrendingScoreEventProducer
}
