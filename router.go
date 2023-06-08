package gateway

type Key string

const (
	RequestKey         Key = "request"
	CancelKey          Key = "cancel"
	RequestIdHeaderKey Key = "X-Request-ID"
)
