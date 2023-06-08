package gateway

type Processor interface {
	Process(handler Handler, request Request, respond bool) (ok bool)
}

func (c *controller) Process(handler Handler, request Request, respond bool) (ok bool) {
	result, err := handler.Handle(request)
	if err != nil {
		c.RespondError(request, err)
		return
	}
	if respond && !request.HasResponded() {
		c.Respond(request, StatusUnknown, result)
	}
	return true
}
