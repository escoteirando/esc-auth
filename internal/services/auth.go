package services

import (
	"context"
	"errors"
	"time"

	"github.com/escoteirando/esc-auth/internal/entities"
	"github.com/escoteirando/esc-auth/internal/services/jwt"
	"gofr.dev/pkg/gofr/container"
)

type AuthService struct {
	db              container.DB
	authExpDuration time.Duration
}

var ErrAuthenticationFail = errors.New("authentication fail")

func NewAuthService(db container.DB) *AuthService {
	return &AuthService{
		db:              db,
		authExpDuration: time.Hour * 24,
	}
}

func (s *AuthService) WithAuthExpDuration(exp time.Duration) *AuthService {
	s.authExpDuration = exp
	return s
}

func (s *AuthService) Authenticate(ctx context.Context, username string, password string) (user entities.UserEntity, err error) {
	row := s.db.QueryRowContext(ctx, "SELECT id,username,password,person_id,role FROM users WHERE username = ?", username)
	if row.Err() != nil {
		err = row.Err()
		return
	}
	if err = row.Scan(&user.Id, &user.UserName, &user.Password, &user.PersonId, &user.Role); err != nil {
		return
	}
	if !CheckPasswordHash(password, user.Password) {
		err = ErrAuthenticationFail
	}

	return
}

func (s *AuthService) GetJWT(user entities.UserEntity) (token string, err error) {
	token, err = jwt.CreateToken(jwt.TokenClaims{
		UserId: user.Id,
		Role:   user.Role.String(),
	}, s.authExpDuration)

	return
}

func (s *AuthService) RefreshJWT(oldJWT string) (token string, err error) {
	claims, err := jwt.ValidateToken(token)
	if err != nil {
		return
	}
	var role entities.RoleType

	return s.GetJWT(entities.UserEntity{
		Id:   claims.UserId,
		Role: role.Parse(claims.Role),
	})
}
