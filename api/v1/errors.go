package v1

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrUsernameOrPassword  = newError(402, "Username or Password error")
	ErrPasswordNotMatch    = newError(403, "Password not match")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrEmailAlreadyUse = newError(1001, "The email is already in use.")
	ErrCannotRefresh   = newError(1002, "Can not refresh account.")
	Err2FARequired     = newError(1003, "2FA is required.")
	Err2FACode         = newError(1004, "2FA code error.")
)
