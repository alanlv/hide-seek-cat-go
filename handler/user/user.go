package user

type CreateRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateResponse struct {
	UserName string `json:"user_name"`
}
