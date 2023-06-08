package gateway

type RouterGroup interface {
	Group(path string) RouterGroup
	RegisterMiddlewares(middlewares ...Handler)

	Create(path string, handlers ...Handler)
	Read(path string, handlers ...Handler)
	Update(path string, handlers ...Handler)
	Delete(path string, handlers ...Handler)
}
