package goztl

// HTTP

type HTTPHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Endpoint metadata

// EndpointField is what an OPTIONS request tells about one field of an endpoint.
type EndpointField struct {
	Type     string `json:"type"`
	Required bool   `json:"required"`
	ReadOnly bool   `json:"read_only"`
}

// EndpointOptions is the answer of an OPTIONS request on an endpoint.
//
// Zentral describes a method only if the token has the permission to use it, so an absent method
// is not the same as a method without fields.
type EndpointOptions struct {
	Name    string                              `json:"name"`
	Actions map[string]map[string]EndpointField `json:"actions"`
}

// SupportsField tells if an endpoint accepts a field for a method.
//
// known is false when the answer does not describe the method. supported has no meaning then, and
// the caller must not conclude that the field is absent.
func (eo EndpointOptions) SupportsField(method, name string) (supported, known bool) {
	fields, ok := eo.Actions[method]
	if !ok {
		return false, false
	}
	_, ok = fields[name]
	return ok, true
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
