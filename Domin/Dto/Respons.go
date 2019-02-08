package dto

type Response struct {
	Status       bool        `json:"Status"`
	Data         interface{} `json:"Data"`
	ErrorMessage string      `json:"ErrorMessage"`
}
