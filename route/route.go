package route

import (
	"ninth-learn/handler"
	"ninth-learn/service"

	_ "ninth-learn/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           News Management API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	api := r.Group("/books")
	{
		api.GET("", server.GetBooks)
		api.GET(":id", server.GetBookById)
		api.POST("", server.CreateBook)
		api.PUT(":id", server.UpdateBook)
		api.DELETE(":id", server.DeleteBook)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
