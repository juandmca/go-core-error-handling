package model

type RubikError struct {
	Error            error              `json:"error"`
	StatusCode       int                `json:"statusCode"`
	ErrorDescription string             `json:"description"`
	ErrorCategory    string             `json:"category"`
	ErrorDetail      []RubikErrorDetail `json:"detail"`
}

type RubikErrorDetail struct {
	ErrorCode        string `json:"code"`
	ErrorDescription string `json:"description"`
	ErrorComponent   string `json:"component"`
}
