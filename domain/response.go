package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
