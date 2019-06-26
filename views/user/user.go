package user

import (
	"github.com/gin-gonic/gin"
	"net/http")

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