package model

type SauronError struct {
	StatusCode       int                 `json:"-"`
	FriendlyMessage  string              `json:"friendlyMessage"`
	TechnicalMessage string              `json:"technicalMessage"`
	ErrorCategory    string              `json:"category"`
	ErrorDetail      []SauronErrorDetail `json:"detail"`
	Path             string              `json:"path"`
}

type SauronErrorDetail struct {
	CustomCode       string `json:"customCode"`
	ErrorDescription string `json:"description"`
	ErrorComponent   string `json:"component"`
}
