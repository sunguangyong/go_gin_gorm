package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"fubangyun.com/basearch/gin/go_gin_gorm/models"

	"net/http"
)

type LoginJson struct{
	User string `from:"user" json:"user" xml:"user" binding:"required"`
	PassWord string `from:"password" json:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var login_json LoginJson
	if err := c.ShouldBind(&login_json); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
		return
	}
	if login_json.User != "manna" || login_json.PassWord != "123456"{
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{"status":"unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status":"you are logged in"})


}

func Register(c *gin.Context) {
	var login_json LoginJson
	if err := c.ShouldBind(&login_json); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Register information is not complete"})
		return
	}
	username := login_json.User
	password := login_json.PassWord
	fmt.Println(username,password)
	err := models.AddOnePassport(username, password)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{"status":"you are logged in"})

}