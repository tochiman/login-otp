package controller

import (
	"login-otp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateOTP(c *gin.Context) {
	secret := c.Query("secret")
	otp := c.Query("otp")

	res := service.ValidateOTP(secret, otp)
	if res {
		ResultWrite := "true"
		c.JSON(http.StatusOK, gin.H{
			"result":ResultWrite,
		})
	}else if !res {
		ResultWrite := "false"
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":ResultWrite,
		})
	}
}