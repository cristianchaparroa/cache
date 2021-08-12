package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type server struct {
	router *gin.Engine
}

func NewServer() *server {
	r := gin.Default()
	s := &server{router: r}
	s.setup()
	return s
}
func (s *server) setup() {
	s.setupRoutes()
}
func (s *server) setupRoutes() {}

func (s *server) Run() {
	port := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf(":%s", port)
	s.router.Run(address)
}
