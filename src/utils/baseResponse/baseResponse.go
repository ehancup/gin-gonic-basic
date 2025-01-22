package br


type BaseSuccessResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success" binding:"default:true" example:"true"`
	Data   any      `json:"data,omitempty"`
}

type BaseSuccessResponsePagination struct {
	Message string `json:"message"`
	Success bool   `json:"success" binding:"default:true" example:"true"`
	Data   any      `json:"data,omitempty"`
	Pagination PaginationResponse `json:"pagination"`
}


