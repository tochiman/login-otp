package controller

import (
	"login-otp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	OTP_URI string = "otpauth://totp/verify_OTP?secret="
)

func CreateQRCode(c *gin.Context){
	data := c.Query("string")

	res := service.CreateQRCode(data)
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.Write(res)
}

func Base32Encode(c *gin.Context) {
	string := c.Query("string")

	res := service.Base32Encode(string)
	if res != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": "200 OK",
			"result": res,
		})
	} else if res == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "500 Internal Server Error",
		})
	}
}

func All(c *gin.Context){
	string := c.Query("string")

	encoded := service.Base32Encode(string)
	res := service.CreateQRCode(OTP_URI + encoded)
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.Write(res)
}