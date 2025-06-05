package events

// CloudEventConsumer defines the interface for consuming CloudEvents
type CloudEventConsumer interface {
	// ConsumePostEvent processes a post event
	ConsumePostEvent(event *CloudEvent) (*PostData, error)
}

// CloudEventProducer defines the interface for producing CloudEvents
type CloudEventProducer interface {
	// ProduceAnalysisEvent creates and sends an analysis event
	ProduceAnalysisEvent(analysisData *AnalysisData) error
	// ProducePostEvent creates and sends a post event
	ProducePostEvent(postData *PostData) error
}