package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jz222/loggy/middlewares"
)

func createSubrouter(prefix string, router *gin.Engine) *gin.RouterGroup {
	return router.Group(prefix)
}

// InitRoutes attaches all routes to the router.
func InitRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors)

	authRoutes(createSubrouter("/auth", router))
	eventRoutes(createSubrouter("/event", router))
	loggingRoutes(createSubrouter("/logging", router))
	serviceRoutes(createSubrouter("/service", router))
	userRoutes(createSubrouter("/user", router))
	organizationRoutes(createSubrouter("/organization", router))
}
