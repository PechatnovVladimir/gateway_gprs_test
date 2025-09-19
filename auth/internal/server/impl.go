package server

import (
	"context"
	auth "github.com/PechatnovVladimir/gateway_gprs_test/auth/pkg/api"
	"log"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	log.Println("нас вызвали")
	return &auth.LoginResponse{Message: "Hello " + req.Name}, nil
}
