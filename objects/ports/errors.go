package ports

import "cache/core"

const (
	ObjectNotFound   = core.Error("object not found")
	ObjectNotStored  = core.Error("object not stored")
	ObjectNotUpdated = core.Error("object not updated")
)
