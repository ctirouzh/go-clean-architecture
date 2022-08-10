package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	//TODO: prepare a response
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	// TODO: Add jwt generated token and any other data according to business logic
}
