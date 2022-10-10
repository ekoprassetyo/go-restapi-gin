package main

import (
	"go-restapi-gin/controllers/usercontroller"
	"go-restapi-gin/pkg/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// intial DB
	mysql.ConnectDatabase()

	r.GET("/api/users", usercontroller.Index)
	r.GET("/api/user/:id", usercontroller.Show)
	r.POST("/api/user", usercontroller.Create)
	r.PUT("/api/user/:id", usercontroller.Update)
	r.DELETE("/api/user", usercontroller.Delete)

	r.Run()
}
