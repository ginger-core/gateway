package gateway

type Body interface {
	Bytes() []byte
}

type ResultGetter interface {
	GetResult() any
}
