package auth

import (
	"avitotask/banners-service/internals/services"
)

type HttpHandlerImpl struct {
	services.UserService
}

func NewAuthHttpHandler(userService services.UserService) HttpHandlerImpl {
	return HttpHandlerImpl{userService}
}
