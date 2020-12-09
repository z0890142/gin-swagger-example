package main

import (
	_ "swagger_test/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title swagger test
// @version 1.0
// @description swagger test example
// @schemes http,https
// @host localhost:8088
// @contact.name Skyler
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @BasePath /api/v1

var a []account

func main() {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g := r.Group("/api/v1")
	{

		g.POST("/Account", InsertAccount)
		g.GET("/Account", GetAccount)
		g.DELETE("/Account/:index", DeleteAccount)
	}

	r.Run(":8087")
}

type account struct {
	Account  string
	Password string
}
