package baseresponse


type BaseSuccessResponse[T any] struct {
	Message string `json:"message"`
	Success bool   `json:"success" binding:"default:true" example:"true"`
	Data    T      `json:"data,omitempty"`
}