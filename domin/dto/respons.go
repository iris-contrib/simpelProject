package dto

//Response dto ....
type Response struct {
	Status       bool        `json:"Status"`
	Data         interface{} `json:"Data"`
	ErrorMessage string      `json:"ErrorMessage"`
}

//NewResponse ...
func NewResponse(status bool, data interface{}, errorMessage string) *Response {
	return &Response{Status: status, Data: data, ErrorMessage: errorMessage}
}
