package code

var (
	Success = ResultCode{
		Code:    200,
		Message: "Запрос выполнился без ошибок",
	}
	Created = ResultCode{
		Code:    201,
		Message: "Идентификатор созданного баннера",
	}
	BadRequest = ResultCode{
		Code:    400,
		Message: "Некоректные данные",
	}
	Unauthorized = ResultCode{
		Code:    401,
		Message: "Пользователь не авторизован",
	}
	Forbidden = ResultCode{
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
