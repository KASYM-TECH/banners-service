package code

var (
	Success = ResultCode{
		Code:    200,
		Message: "Запрос выполнился без ошибок",
	}
	BadRequest = ResultCode{
		Code:    400,
		Message: "Невалидный запрос",
	}
	UserNotFound = ResultCode{
		Code:    401,
		Message: "Пользователь не авторизован",
	}
	Unauthorized = ResultCode{
		Code:    403,
		Message: "Пользователь не имеет доступа",
	}
	BannerNotFound = ResultCode{
		Code:    404,
		Message: "Баннер для пользователя не найден",
	}
	InternalError = ResultCode{
		Code:    500,
		Message: "Внутренняя ошибка сервера",
	}
)
