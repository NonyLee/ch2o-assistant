package resp

type ErrorMsg struct {
	code string
	msg  string
}

type CommonResp struct {
	Success bool        `json:"success"`
	Error   *ErrorMsg   `json:"error"`
	Result  interface{} `json:"result"`
}
