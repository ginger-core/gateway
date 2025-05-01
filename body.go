package gateway

type Body interface {
	Bytes() []byte
}

type ResultGetter interface {
	GetResult() any
}

type MapResultGetter interface {
	GetMap() map[string]any
}
