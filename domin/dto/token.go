package dto

//Token dto ...
type Token struct {
	AuthToken string `json:"Auth_Token"`
}

//NewToken ...
func NewToken(authToken string) *Token {
	return &Token{AuthToken: authToken}
}
