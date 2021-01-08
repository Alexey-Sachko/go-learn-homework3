package types

//easyjson:json
type UserInfo struct {
	Browsers []string `json:"browsers"`
	Email string `json:"email"`
	Name string `json:"name"`
}