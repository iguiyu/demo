package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iguiyu/demo/api/middlewares"
	"github.com/itsjamie/gin-cors"
)

func MakeHandlersWithRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET,PUT,POST,DELETE",
		RequestHeaders:  "Origin,Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          1 * time.Minute,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router := engine.Group("/")
	router.Use(middlewares.OAuthUser())

	// router.GET("/", ...)

	return engine
}
