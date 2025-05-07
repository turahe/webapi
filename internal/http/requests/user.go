package requests

type AuthLoginRequest struct {
	UserName string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password"`
}

type AuthRegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	UserName        string `json:"username" validate:"required,min=3,max=32"`
	Phone           string `json:"phone" validate:"required,numeric"`
	Password        string `json:"password" validate:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
