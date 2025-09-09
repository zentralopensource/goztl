package goztl

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
