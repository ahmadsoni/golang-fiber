package services

import (
	"context"
	"errors"
	"gofiber-restapi/domain"
	"gofiber-restapi/dto"
	"gofiber-restapi/internal/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &authService{
		conf:           cnf,
		userRepository: userRepository,
	}
}

func (a authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		return dto.AuthResponse{}, err
	}
	if user.Id == "" {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}
	checkPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if checkPass != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}
	claim := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))

	if err != nil {
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}
