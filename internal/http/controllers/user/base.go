package user

import "webapi/internal/app/user"

type UserHTTPHandler struct {
	app user.UserApp
}

func NewUserHTTPHandler(app user.UserApp) *UserHTTPHandler {
	return &UserHTTPHandler{app: app}
}
