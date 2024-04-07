package banner

import "avitotask/banners-service/internals/services"

type HttpHandlerImpl struct {
	services.BannerService
}

func NewBannerHttpHandler(bannerService services.BannerService) HttpHandlerImpl {
	return HttpHandlerImpl{bannerService}
}
