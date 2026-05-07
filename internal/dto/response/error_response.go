package response

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
