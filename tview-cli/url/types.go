package url

type CreateHttpUrlConfig struct {
	Url           string `json:"url"`
	FriendlyName  string `json:"friendly_name"`
	Interval      int    `json:"interval"`
	Retries       int    `json:"retries"`
	Timeout       int    `json:"timeout"`
	UpsideDown    bool   `json:"upside_down"`
	MaxRedirects  int    `json:"max_redirects"`
	Method        string `json:"method"`
	AcceptedCodes string `json:"accepted_codes"`
}
