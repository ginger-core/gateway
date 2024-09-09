package gateway

type RouterGroup interface {
	Group(path string) RouterGroup
	RegisterMiddlewares(middlewares ...Handler)

	On(method Method, handlers ...Handler)
	OnPath(method Method, path string, handlers ...Handler)
}

type Method string

const (
	Unknown Method = ""
	Create  Method = "CREATE"
	Read    Method = "READ"
	Update  Method = "UPDATE"
	Delete  Method = "DELETE"
)
