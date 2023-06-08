package gateway

type Status uint16

const (
	StatusUnknown           Status = 0
	StatusOK                Status = 200
	StatusCreated           Status = 201
	StatusNoContent         Status = 204
	StatusMovedPermanently  Status = 301
	StatusFound             Status = 302
	StatusPermanentRedirect Status = 307
	StatusTemporaryRedirect Status = 308
	StatusBadRequest        Status = 400
	StatusDuplicate         Status = 409
)
