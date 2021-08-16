package delivery

import (
	"cache/app/adapters/handlers"
	_ "cache/app/conf"
	"cache/core"
	// It calls the provides managers to map into injector
	_ "cache/objects/managers"
)

func loadObjectsHandler() (*handlers.ObjectsHandler, error) {
	var handler *handlers.ObjectsHandler

	invokeFun := func(h *handlers.ObjectsHandler) {
		handler = h
	}

	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
