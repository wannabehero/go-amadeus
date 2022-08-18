package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"time"
)

// Generates request signature with the following algorithm:
//   sha1(nonce + timestamp + sha1(password))
// nonce - random 16 bytes
// timestamp - utc ISO timestamp
// password - provided
func CreateSignature(password string) (timestamp, nonce, signature string) {
	timestamp = time.Now().UTC().Format(time.RFC3339Nano)

	token := make([]byte, 16)
	rand.Read(token)
	nonce = base64.StdEncoding.EncodeToString(token)

	passwordHash := sha1.New()
	io.WriteString(passwordHash, password)

	signatureHash := sha1.New()
	signatureHash.Write(token)
	io.WriteString(signatureHash, timestamp)
	signatureHash.Write(passwordHash.Sum(nil))

	signature = base64.StdEncoding.EncodeToString(signatureHash.Sum(nil))

	return
}
