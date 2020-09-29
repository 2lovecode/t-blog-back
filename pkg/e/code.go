package e

type RCode int64
const (
	Success RCode = 200000
	Error = 500000

	ErrorInvalidParams = 400000

	ErrorExistModule = 400100

	ErrorInvalidUserOrPass = 400200
	ErrorInvalidToken = 400201

)
