package goztl

type Digicert struct {
	APIBaseURL       string `json:"api_base_url"`
	APIToken         string `json:"api_token"`
	ProfileGUID      string `json:"profile_guid"`
	BusinessUnitGUID string `json:"business_unit_guid"`
	SeatType         string `json:"seat_type"`
	SeatIDMapping    string `json:"seat_id_mapping"`
	DefaultSeatEmail string `json:"default_seat_email"`
}

type IDent struct {
	URL            string `json:"url"`
	BearerToken    string `json:"bearer_token"`
	RequestTimeout int    `json:"request_timeout"`
	MaxRetries     int    `json:"max_retries"`
}

type MicrosoftCA struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StaticChallenge struct {
	Challenge string `json:"challenge"`
}
