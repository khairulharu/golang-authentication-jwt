package dto

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"string"`
	Name     string `json:"name"`
}
