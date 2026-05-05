package route

import (
	"subscription/internal/config"
	"subscription/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	envConf  *config.Config
}

func NewRouter(services *service.Service, envConf *config.Config) *Handler {
	return &Handler{services: services, envConf: envConf}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	return r
}
