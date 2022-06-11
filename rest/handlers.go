package rest

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()

	RegisterHandlers(engine)

	return engine
}

func RegisterHandlers(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := engine.Group("/api/v1")
	RegisterUsersHandlers(v1)
}

func RegisterUsersHandlers(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("", getUsers)
		users.GET("/:id", getUserByID)
		users.POST("", createUser)
		users.PUT("/:id", updateUser)
		users.DELETE("/:id", deleteUser)
	}
}
