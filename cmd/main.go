package main

import (
	"github.com/gin-gonic/gin"
	"fubangyun.com/basearch/gin/go_gin_gorm/views/user"
)



func main() {
	router := gin.Default()
	router.POST("/longin",user.Login)
	router.POST("/register",user.Register)
	router.Run(":8000")
}


