package user

import (
	"fubangyun.com/basearch/gin/go_gin_gorm/models"
	"github.com/gin-gonic/gin"

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

	err,user := models.Mydb.GetOnePassport(login_json.User)
	if err != nil{
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{"status":err})
		return
	}
	if len(user) == 0{
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{"status" : "The account does not exist"})
		return
	}
	if user[0].PassWord == login_json.PassWord {
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{"status" : "login successfully"})
		return
	}
}

func Register(c *gin.Context) {
	var login_json LoginJson
	if err := c.ShouldBind(&login_json); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Register information is not complete"})
		return
	}
	username := login_json.User
	password := login_json.PassWord
	err, user := models.Mydb.GetOnePassport(username)
	if len(user) != 0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account already exists"})
	} else {
		err = models.Mydb.AddOnePassport(username, password)
		if err == nil{
			c.JSON(http.StatusOK, gin.H{"status":"registered successfully"})
		}else {
			c.JSON(http.StatusOK, gin.H{"status": err})
		}

	}
}