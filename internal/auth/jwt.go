package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type Claims struct {
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

func Sign(sub, secret string) string {
	head := map[string]string{"alg": "HS256", "typ": "JWT"}
	body := Claims{Sub: sub, Exp: time.Now().Add(24 * time.Hour).Unix()}
	a := encode(head)
	b := encode(body)
	unsigned := a + "." + b
	return unsigned + "." + sig(unsigned, secret)
}

func Verify(token, secret string) (Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 || sig(parts[0]+"."+parts[1], secret) != parts[2] {
		return Claims{}, errors.New("bad token")
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return Claims{}, err
	}
	var claims Claims
	if err := json.Unmarshal(raw, &claims); err != nil || claims.Exp < time.Now().Unix() {
		return Claims{}, errors.New("expired token")
	}
	return claims, nil
}

func encode(v any) string {
	b, _ := json.Marshal(v)
	return base64.RawURLEncoding.EncodeToString(b)
}

func sig(unsigned, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(unsigned))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
