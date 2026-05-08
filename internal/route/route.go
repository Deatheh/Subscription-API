package route

import (
	"subscription/internal/config"
	"subscription/internal/logger"
	"subscription/internal/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	SubscriptionRoute                = "/subscription"
	SubscriptionRouteWithId          = "/subscription/:id"
	SubscriptionGetSumByFiltersRoute = "/subscription/filters"
	SwaggerRoute                     = "/swagger/*any"
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
	r.GET(SwaggerRoute, ginSwagger.WrapHandler(swaggerfiles.Handler))
	subscr := r.Group("")
	{
		subscr.Use(logger.RequestLogger("subscription"))
		subscr.POST(SubscriptionRoute, h.HandleAddSubscription)
		subscr.GET(SubscriptionRoute, h.HandleGetAllSubscription)
		subscr.GET(SubscriptionRouteWithId, h.HandleGetSubscriptionById)
		subscr.PUT(SubscriptionRouteWithId, h.HandleUpdateSubscription)
		subscr.DELETE(SubscriptionRouteWithId, h.HandleDeleteSubscription)
		subscr.GET(SubscriptionGetSumByFiltersRoute, h.HandleGetSumSubscriptionByFilters)

	}

	return r
}
