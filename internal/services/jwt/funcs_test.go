package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndValidateToken(t *testing.T) {
	tokenString, err := CreateToken(TokenClaims{UserId: 1234, Role: "test_admin"}, time.Hour)
	t.Logf("Token: %s", tokenString)
	if !assert.NoError(t, err) {
		return
	}

	claim, err := ValidateToken(tokenString)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, 1234, claim.UserId)
	assert.Equal(t, "test_admin", claim.Role)
}
func Test_getJWTSecretFromEnv(t *testing.T) {
	t.Setenv(EnvJWTSecret, "")
	secret := getJWTSecretFromEnv()
	assert.Empty(t, secret)

	t.Setenv(EnvJWTSecret, "test")
	secret = getJWTSecretFromEnv()
	assert.Equal(t, []byte{0x74, 0x65, 0x73, 0x74}, secret)
}

func Test_getJWTSecretRandom(t *testing.T) {
	secret := getJWTSecretRandom()
	assert.Len(t, secret, 48)
}
