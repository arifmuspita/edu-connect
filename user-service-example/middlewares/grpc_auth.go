package middlewares

import (
	"context"
	"strings"

	"github.com/zuyatna/edu-connect/user-service-example/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
)

var publicEndpoints = map[string]bool{
	"/user.UserService/RegisterUser": true,
	"/user.UserService/LoginUser":    true,
}

func SelectiveAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if publicEndpoints[info.FullMethod] {
		return handler(ctx, req)
	}

	return AuthGRPCInterceptor(ctx, req, info, handler)
}

func AuthGRPCInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
	}

	tokenParts := strings.Split(authHeader[0], " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token format")
	}

	tokenString := tokenParts[1]
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	userID := (*claims)["user_id"].(string)
	newCtx := context.WithValue(ctx, UserIDKey, userID)

	return handler(newCtx, req)
}
