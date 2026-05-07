package route

import (
	"subscription/internal/config"
	"subscription/internal/logger"
	"subscription/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	SubscriptionRoute                = "/subscription"
	SubscriptionRouteWithId          = "/subscription/:id"
	SubscriptionGetSumByFiltersRoute = "/subscription/filters"
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
	subscr := r.Group("")
	{
		subscr.Use(logger.RequestLogger("subscription"))
		subscr.POST(SubscriptionRoute, h.AddSubscription)
		subscr.GET(SubscriptionRoute, h.GetAllSubscription)
		subscr.GET(SubscriptionRouteWithId, h.GetSubscriptionById)
		subscr.GET(SubscriptionGetSumByFiltersRoute, h.GetSumSubscriptionByFilters)
	}

	return r
}
