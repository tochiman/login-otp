package main

import (
    "bufio"
    "fmt"
    "os"

	"login-otp/controller"
)

func main() {
    fmt.Print("input? ")
    // Scannerを使って一行読み
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	
	res := controller.ValidateOTP("tochiman", scanner.Text())

	fmt.Println(res)
}