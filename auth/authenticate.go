package auth

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

func Authenticate(did string) string {

	// integrationtest@fairx.io
	authRequest := map[string]string{
		"id":  did,
		"sig": "FAKESIG",
	}
	authRequestBytes, _ := json.Marshal(&authRequest)

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(authRequestBytes).
		Post("http://localhost:8000/v1.0.0/auth")

	if err != nil {
		return ""
	}

	var jsonResponse map[string]string
	err = json.Unmarshal(resp.Body(), &jsonResponse)
	if err != nil {
		return ""
	}

	jwt := jsonResponse["jwt"]
	return jwt

}
