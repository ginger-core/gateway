package gateway

type Body interface {
	Bytes() []byte
}

type ResultGetter interface {
	GetDeliveryResult() any
}

type MapResultGetter interface {
	GetMap() map[string]any
}
