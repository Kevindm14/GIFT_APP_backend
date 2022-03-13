package qr

import (
	"encoding/base64"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func CodeForGift(baseURL, code string) (string, error) {
	path := fmt.Sprintf("%v/%v", baseURL, code)

	q, err := qrcode.New(path, qrcode.Medium)
	if err != nil {
		panic(err)
	}

	q.DisableBorder = true

	var qr []byte
	qr, err = q.PNG(256)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(qr), nil
}
