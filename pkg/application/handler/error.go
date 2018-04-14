package handler

type (
	// ErrorResponse struct for error request response
	ErrorResponse struct {
		Message string `json:"message"`
		Error   error  `json:"error,omitempty"`
	}
)
