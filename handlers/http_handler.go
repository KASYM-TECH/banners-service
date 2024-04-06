package handlers

import "avitotask/banners-service/internals/services"

type HttpHandlerImpl struct {
	services.UserService
}

func NewHttpHandler(user services.UserService) HttpHandlerImpl {
	return HttpHandlerImpl{user}
}
