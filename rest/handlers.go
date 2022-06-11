package rest

import (
	"net/http"

	docs "github.com/alehechka/swagger-api/docs"
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
	RegisterSwaggerHandlers(engine)

	v1 := engine.Group("/api/v1")
	RegisterUsersHandlers(v1)
}

func RegisterSwaggerHandlers(engine *gin.Engine) {
	engine.GET("/", func(ctx *gin.Context) { ctx.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html") })
	engine.GET("/swagger/*any", func(ctx *gin.Context) {
		docs.SwaggerInfo.BasePath = ctx.Request.Host
	}, ginSwagger.WrapHandler(swaggerFiles.Handler))
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
