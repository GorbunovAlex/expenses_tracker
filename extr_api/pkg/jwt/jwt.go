package jwt

import (
	"alex_gorbunov_exptr_api/internal/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func GenerateToken(header string, payload map[string]string, secret string) (string, error) {
	const fn = "jwt.Generate"

	hash := hmac.New(sha256.New, []byte(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	paylaodstr, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("%s: %w", fn, err)
	}
	payload64 := base64.StdEncoding.EncodeToString(paylaodstr)

	message := header64 + "." + payload64

	unsignedStr := header + string(paylaodstr)

	hash.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	tokenStr := message + "." + signature

	return tokenStr, nil
}

func ValidateToken(token string) (bool, error) {
	cfg := config.MustLoad()

	const fn = "jwt.ValidateToken"

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, fmt.Errorf("%s: invalid token", fn)
	}

	header, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, fmt.Errorf("%s: %w", fn, err)
	}

	payload, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, fmt.Errorf("%s: %w", fn, err)
	}

	unsighedStr := string(header) + string(payload)
	hash := hmac.New(sha256.New, []byte(cfg.HTTPServer.JwtSecret))
	hash.Write([]byte(unsighedStr))

	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	if parts[2] != signature {
		return false, nil
	}

	return true, nil
}

func GetSignedToken() (string, error) {
	const fn = "jwt.getSignedToken"

	cfg := config.MustLoad()

	claimsMap := map[string]string{
		"aud": "frontend.epxtr",
		"iss": "backend.exptr",
		"exp": fmt.Sprintf("%d", time.Now().Add(time.Hour).Unix()),
	}

	secret := cfg.HTTPServer.JwtSecret
	header := "HS256"
	token, err := GenerateToken(header, claimsMap, secret)
	if err != nil {
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	return token, nil
}
