package jwt

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/escoteirando/esc-auth/internal/metadata"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/rand"
)

const EnvJWTSecret = "JWT_SECRET"

type TokenClaims struct {
	UserId int
	Role   string
}

var getJWTSecret = sync.OnceValue(func() []byte {
	secret := getJWTSecretFromEnv()
	if len(secret) == 0 {
		secret = getJWTSecretRandom()
	}
	return secret
})

func getJWTSecretFromEnv() []byte {
	secret := os.Getenv(EnvJWTSecret)
	return []byte(secret)
}

func getJWTSecretRandom() []byte {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const n = 48
	jwtSecret := make([]byte, n)
	src := rand.NewSource(uint64(time.Now().UnixNano())) //nolint:gosec
	for i := 0; i < n; i++ {
		jwtSecret[i] = letterBytes[src.Uint64()%uint64(len(letterBytes))]
	}
	fmt.Printf("MISSING ENV %s - JWT SECRET RANDOMICALLY GENERATED: %s\n", EnvJWTSecret, jwtSecret)
	return jwtSecret
}

func CreateToken(claims TokenClaims, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(claims.UserId),                               // Subject: user id
		"iss": fmt.Sprintf("%s v%s", metadata.AppName, metadata.Version), // Issuer
		"aud": claims.Role,                                               // Audience (user role)
		"exp": time.Now().Add(exp).Unix(),                                // Expiration time
		"iat": time.Now().Unix(),
	})
	return token.SignedString(getJWTSecret())
}

func ValidateToken(tokenString string) (claims *TokenClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = fmt.Errorf("invalid token")
		return
	}

	// user id
	uid, _ := token.Claims.GetSubject()
	var userId int
	if userId, err = strconv.Atoi(uid); err != nil {
		err = fmt.Errorf("invalid subject: %s", uid)
		return
	}

	// role
	aud, _ := token.Claims.GetAudience()
	if len(aud) == 0 {
		err = fmt.Errorf("invalid audience")
		return
	}

	//
	iss, _ := token.Claims.GetIssuer()
	if iss != fmt.Sprintf("%s v%s", metadata.AppName, metadata.Version) {
		err = fmt.Errorf("invalid issuer %s", iss)
	}
	if err == nil {
		claims = &TokenClaims{
			UserId: userId,
			Role:   aud[0],
		}
	}
	return
}
