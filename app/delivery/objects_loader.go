package delivery

import (
	"cache/app/adapters/handlers"
	// It calls the provides repositories to map into injector
	_ "cache/app/adapters/repositories"
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
