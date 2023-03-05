package model

type RubikError struct {
	StatusCode       int                `json:"-"`
	FriendlyMessage  string             `json:"friendlyMessage"`
	TechnicalMessage string             `json:"technicalMessage"`
	ErrorCategory    string             `json:"category"`
	ErrorDetail      []RubikErrorDetail `json:"detail"`
	Path             string             `json:"path"`
}

type RubikErrorDetail struct {
	CustomCode       string `json:"customCode"`
	ErrorDescription string `json:"description"`
	ErrorComponent   string `json:"component"`
}
