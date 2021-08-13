package delivery

func setupObjectsRoutes(s *server) {
	handler, err := loadObjectsHandler()
	if err != nil {
		panic(err)
	}

	objectGroup := s.router.Group("/object")
	objectGroup.GET(":key", handler.GetByKey)
	objectGroup.POST("", handler.Save)
	objectGroup.PUT(":key", handler.Save)
	objectGroup.DELETE(":key", handler.DeleteByKey)
}
