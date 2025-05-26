package domain

import (
	"context"
	"gofiber-restapi/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
