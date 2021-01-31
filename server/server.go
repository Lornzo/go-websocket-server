package server

import (
	"crypto/sha1"
	"encoding/base64"
)

func GetSecWebSocketAccept(key string) string {
	result := ""
	if len(key) > 0 {
		var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")
		h := sha1.New()
		h.Write([]byte(key))
		h.Write(keyGUID)
		result = base64.StdEncoding.EncodeToString(h.Sum(nil))
	}
	return result
}
