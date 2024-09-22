package constant

const (
	// Account Custom Error Message
	ErrorAccountNotFound = "account not found"

	// Account Transaction Custom Error Message
	ErrorTransactionNotFound = "transaction not found"

	// User Custom Error Message
	ErrorUserAlreadyExists = "user already exists"
	ErrorUserNotFound      = "user not found"

	// HTTP Error Message
	ErrorInternalServerError = "internal server error"
)

var mapErrorToHTTPStatusCode = map[string]int{
	ErrorAccountNotFound: 404,

	ErrorTransactionNotFound: 404,

	ErrorUserAlreadyExists: 409,
	ErrorUserNotFound:      404,

	ErrorInternalServerError: 500,
}

func GetHTTPStatusCodeByError(err string) int {
	if statusCode, ok := mapErrorToHTTPStatusCode[err]; ok {
		return statusCode
	}

	return 500
}
