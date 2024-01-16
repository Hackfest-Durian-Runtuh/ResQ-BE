package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func SendSMS(token, to, msg string) error {
	// Rubah message menjadi url encoded
	msg = url.QueryEscape(msg)
	url := fmt.Sprintf("https://websms.co.id/api/smsgateway-otp?token=%s&to=%s&msg=%s", token, to, msg)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	fmt.Printf("Response: %s\n", body)
	return nil
}

func GenerateOTP() string {
	return EncodeToString(6)
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, _ := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return ""
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
