package main

import (
	"login-otp/controller"

	"github.com/gin-gonic/gin"
)

func main() {

    router := gin.Default()

    apipath := router.Group("/api")
	v1path := apipath.Group("/v1")
    {
        v1path.GET("/base32", controller.Base32Encode)
        v1path.GET("/auth", controller.ValidateOTP)
    }
    qrpath := v1path.Group("/qrcode")
    {
        qrpath.GET("/create", controller.CreateQRCode)
        qrpath.GET("all", controller.All)
    }

    router.Run(":8080")
}