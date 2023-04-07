package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
)

func New(secret []byte, counter uint64, digits int) string{
	hs := hmacSha1(secret, counter)
	num := dynamicTruncate(hs)
	codeNum := num % uint32(math.Pow10(digits))
	return format(codeNum, digits)
}

func format(code uint32, digits int) string{
	f := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(f, code)
}

func hmacSha1(secret []byte, counter uint64) []byte{
	mac := hmac.New(sha1.New, secret)
	byteCounter := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCounter, counter)

	mac.Write(byteCounter)
	return mac.Sum(nil)
}

func dynamicTruncate(hs []byte) uint32 {
	offset := hs[len(hs)-1] & 0xf

	return uint32(hs[offset]&0x7f) <<24 |
		uint32(hs[offset+1]&0xff)<<16 |
		uint32(hs[offset+2]&0xff)<<8 |
		uint32(hs[offset+3]&0xff)
}