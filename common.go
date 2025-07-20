package goztl

// HTTP

type HTTPHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Event Filters

type EventFilter struct {
	Tags       []string `json:"tags,omitempty"`
	EventType  []string `json:"event_type,omitempty"`
	RoutingKey []string `json:"routing_key,omitempty"`
}

type EventFilterSet struct {
	ExcludedEventFilters []EventFilter `json:"excluded_event_filters,omitempty"`
	IncludedEventFilters []EventFilter `json:"included_event_filters,omitempty"`
}
