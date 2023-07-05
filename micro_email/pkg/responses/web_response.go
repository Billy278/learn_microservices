package responses

const (
	InvalidParam       = "invalid param request"
	InvalidBody        = "invalid body request"
	InvalidPayload     = "invalid payload request"
	InvalidQuery       = "invalid query request"
	InternalServer     = "internal server error"
	SomethingWentWrong = "something went wrong"
	Unauthorized       = "unauthorized request"
)

type SuccessRes struct {
	Code int `json:"code"`
	Data any `json:"data"`
}
type FailRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
