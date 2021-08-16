package managers

import "cache/core"

const (
	objectNotFound   = core.Error("object not found")
	objectNotStored  = core.Error("object not stored")
	objectNotUpdated = core.Error("object not updated")
)
