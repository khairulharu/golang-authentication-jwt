package dto

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type UserResponseWithToken struct {
	UserResponse
	Token string `json:"token"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
