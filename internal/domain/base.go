package domain

// ResponseResult stores the status of a query
type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
