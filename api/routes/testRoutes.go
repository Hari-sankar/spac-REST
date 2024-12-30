package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterTestRoutes(router *gin.Engine) {

	testRouter := router.Group("/test")
	{
		// Serve the Swagger JSON file
		testRouter.GET("/swagger-json", func(c *gin.Context) {
			c.File("./docs/swagger.json")
		})

		// Customize the Swagger UI to fetch from /swagger-json
		swaggerURL := ginSwagger.URL("http://localhost:3000/test/swagger-json") // Replace with your URL
		testRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))
		// testRouter.GET("/swagger/*any", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"status":  "success",
		// 		"message": "Swagger documentation loaded successfully",
		// 	})
		// })
	}
}
