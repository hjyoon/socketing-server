package auth

import (
	"encoding/base64"
	"testing"
	"time"
)

func TestPasswordHashAndPBKDF2(t *testing.T) {
	got := HashPassword("password", "salt")
	if got != HashPassword("password", "salt") || len(got) != 128 {
		t.Fatalf("hash must be deterministic hex")
	}
	if len(NewSalt()) != 32 {
		t.Fatalf("salt length mismatch")
	}
}

func TestJWTSignVerify(t *testing.T) {
	token := Sign("u", "s")
	claims, err := Verify(token, "s")
	if err != nil || claims.Sub != "u" {
		t.Fatalf("verify failed")
	}
	if _, err := Verify(token, "bad"); err == nil {
		t.Fatalf("bad secret accepted")
	}
	if _, err := Verify("bad", "s"); err == nil {
		t.Fatalf("bad token accepted")
	}
	unsigned := "e30." + base64.RawURLEncoding.EncodeToString([]byte(`{"sub":"u","exp":1}`))
	if _, err := Verify(unsigned+"."+sig(unsigned, "s"), "s"); err == nil {
		t.Fatalf("expired token accepted")
	}
	unsigned = "e30." + base64.RawURLEncoding.EncodeToString([]byte(`{`))
	if _, err := Verify(unsigned+"."+sig(unsigned, "s"), "s"); err == nil {
		t.Fatalf("bad json accepted")
	}
	if time.Now().Unix() == 0 {
		t.Fatalf("unreachable")
	}
}
