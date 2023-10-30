package admin

type AdminRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
