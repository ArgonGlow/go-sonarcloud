package webservices

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ListRequest List web services
type ListRequest struct{}

// ListResponse is the response for ListRequest
type ListResponse struct {
	WebServices []struct {
		Actions []struct {
			Changelog []struct {
				Description string `json:"description,omitempty"`
				Version     string `json:"version,omitempty"`
			} `json:"changelog,omitempty"`
			DeprecatedSince    string `json:"deprecated_since,omitempty"`
			Description        string `json:"description,omitempty"`
			HasResponseExample bool   `json:"has_response_example,omitempty"`
			Internal           bool   `json:"internal,omitempty"`
			Key                string `json:"key,omitempty"`
			Params             []struct {
				Internal           bool     `json:"internal,omitempty"`
				Key                string   `json:"key,omitempty"`
				MaximumValue       float64  `json:"maximum_value,omitempty"`
				Required           bool     `json:"required,omitempty"`
				MaximumLength      float64  `json:"maximum_length,omitempty"`
				MinimumLength      float64  `json:"minimum_length,omitempty"`
				DefaultValue       string   `json:"default_value,omitempty"`
				DeprecatedKey      string   `json:"deprecated_key,omitempty"`
				DeprecatedKeySince string   `json:"deprecated_key_since,omitempty"`
				DeprecatedSince    string   `json:"deprecated_since,omitempty"`
				Description        string   `json:"description,omitempty"`
				ExampleValue       string   `json:"example_value,omitempty"`
				PossibleValues     []string `json:"possible_values,omitempty"`
			} `json:"params,omitempty"`
			Post bool `json:"post,omitempty"`
		} `json:"actions,omitempty"`
		Description string `json:"description,omitempty"`
		Path        string `json:"path,omitempty"`
	} `json:"web_services,omitempty"`
}

// ResponseExampleRequest Display web service response example
type ResponseExampleRequest struct {
	Action     string `form:"action,omitempty"`     // Action of the web service
	Controller string `form:"controller,omitempty"` // Controller of the web service
}

// ResponseExampleResponse is the response for ResponseExampleRequest
type ResponseExampleResponse struct {
	Example string `json:"example,omitempty"`
	Format  string `json:"format,omitempty"`
}
