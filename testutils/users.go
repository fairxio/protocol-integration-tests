package testutils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func CreateRandomDID() (string, string) {
	return CreateRandomDIDInDomain("fairx.io")
}

func CreateRandomDIDInDomain(domain string) (string, string) {

	// Create random number up to 6 digits
	userDomain := fmt.Sprintf("test%s@%s", EncodeToString(6), domain)
	userDomainEncoded := base64.RawURLEncoding.EncodeToString([]byte(userDomain))
	return userDomain, fmt.Sprintf("did:fairx:%s", userDomainEncoded)

}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
