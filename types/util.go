package types

type ApiResponse struct {
	Result      int64       `json:"result"`
	Description string      `json:"description"`
	ErrCode     interface{} `json:"errCode"`
}

func NewApiResponse(description string, result int64, errCode any) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
		ErrCode:     errCode,
	}
}
