package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func NewSalt() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func HashPassword(password, salt string) string {
	return hex.EncodeToString(PBKDF2([]byte(password), []byte(salt), 1000, 64))
}

func PBKDF2(pass, salt []byte, iter, keyLen int) []byte {
	hLen := 32
	out := make([]byte, 0, keyLen)
	for block := 1; len(out) < keyLen; block++ {
		u := prf(pass, salt, block)
		t := append([]byte(nil), u...)
		for i := 1; i < iter; i++ {
			u = hmacSum(pass, u)
			for j := range t {
				t[j] ^= u[j]
			}
		}
		out = append(out, t[:min(hLen, keyLen-len(out))]...)
	}
	return out
}

func prf(pass, salt []byte, block int) []byte {
	msg := append([]byte(nil), salt...)
	msg = append(msg, byte(block>>24), byte(block>>16), byte(block>>8), byte(block))
	return hmacSum(pass, msg)
}

func hmacSum(key, msg []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return mac.Sum(nil)
}
