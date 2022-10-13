package dto

type DtoResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewDtoResponse(message string, data any) *DtoResponse {
	return &DtoResponse{
		Message: message,
		Data:    data,
	}
}
