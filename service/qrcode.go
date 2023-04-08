package service

import (
	"github.com/skip2/go-qrcode"
	"encoding/base32"
)

const (
	// Size of the QRCode... ex)256 => 256*256
	QRCODE_SIZE int=256
)

func CreateQRCode(content string) []byte {
	qrCode, err := qrcode.Encode(content, qrcode.Medium, QRCODE_SIZE)
	if err != nil {
		return []byte(err.Error())
	}

	return qrCode
}

func Base32Encode(str string) string {
	byteSlice := []byte(str)
	return base32.StdEncoding.EncodeToString(byteSlice)
}