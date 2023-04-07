package controller

import (
	"fmt"
	"login-otp/totp"
)

const otpDigits int = 6

func ValidateOTP(secret string, otp string) bool {
	onTimeotp := totp.New([]byte(secret), otpDigits, 0)
	fmt.Println(onTimeotp)
	beforeOtp := totp.New([]byte(secret), otpDigits, -1)
	fmt.Println(beforeOtp)
	return otp == onTimeotp || otp == beforeOtp
}