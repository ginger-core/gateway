package gateway

type Authorization interface {
	Set(key string, value any)
	Get(key string) any

	GetApplicantId() any
	GetApplicant() any
}
