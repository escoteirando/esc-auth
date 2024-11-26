package services

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/escoteirando/esc-auth/internal/entities"
	"github.com/golang-jwt/jwt/v5"
	"gofr.dev/pkg/gofr/container"
	"golang.org/x/exp/rand"
)

type AuthService struct {
	db        container.DB
	jwtSecret []byte
}

var ErrAuthenticationFail = errors.New("authentication fail")

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandByteStr(n int) []byte {
	b := make([]byte, n)
	src := rand.NewSource(uint64(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		b[i] = byte(letterBytes[src.Uint64()%uint64(len(letterBytes))])
	}

	return b
}

func NewAuthService(db container.DB) *AuthService {
	jwtSecret := os.Getenv("SECRET")
	var jwt []byte
	if len(jwtSecret) > 0 {
		jwt = []byte(jwtSecret)
	} else {
		jwt = RandByteStr(48)
	}
	return &AuthService{
		db:        db,
		jwtSecret: jwt,
	}
}

func (s *AuthService) Authenticate(ctx context.Context, username string, password string) (user entities.UserEntity, err error) {
	row := s.db.QueryRowContext(ctx, "SELECT id,username,password,person_id FROM users WHERE username = ?", username)
	if row.Err() != nil {
		err = row.Err()
		return
	}
	if err = row.Scan(&user.Id, &user.UserName, &user.Password, &user.PersonId); err != nil {
		return
	}
	if !CheckPasswordHash(password, user.Password) {
		err = ErrAuthenticationFail
	}

	return
}

func (s *AuthService) GetJWT(user entities.UserEntity) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(user.Id)),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})
	token, err = claims.SignedString(s.jwtSecret)
	return
}
