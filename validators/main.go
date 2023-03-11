package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": Translate(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "成功",
		})
	}
}
func main() {
	fmt.Println("美化gin的validator验证器")
	_ = ValidatorInit("zh")
	g := gin.Default()
	g.POST("/login", LoginHandler)
	_ = g.Run(":3600")
}
