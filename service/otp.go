package service

import (
	"login-otp/totp"
)

const (
	otpDigits int = 6
)

func ValidateOTP(secret string, otp string) bool {
	onTimeotp := totp.New([]byte(secret), otpDigits, 0)
	beforeOtp := totp.New([]byte(secret), otpDigits, -1)
	return otp == onTimeotp || otp == beforeOtp
}