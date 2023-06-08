package gateway

type CtxKey string

const (
	RequestIdKey CtxKey = "ReqId"
	IPKey        CtxKey = "IP"
	AgentKey     CtxKey = "Agent"
)
