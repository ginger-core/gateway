package request

import (
	"github.com/ginger-core/gateway"
)

func (r *request) WithConn(conn any) gateway.Request {
	return r
}

func (r *request) GetConn() any {
	return nil
}
