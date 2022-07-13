package responses

type UserResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
	Error  any    `json:"error"`
}
