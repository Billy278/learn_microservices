package middleware

import (
	"context"
	"errors"
	"micro_balance/modules/proto"
	"micro_balance/pkg/crypto"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func WithMiddlewareUnarryInceptor() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(serverInterceptor)
}

func serverInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	if err := AuthKeyGRPC(ctx); err != nil {
		return &proto.Response{
			Code:  http.StatusUnauthorized,
			Error: err.Error(),
		}, err
	}
	// Calls the handler
	h, err := handler(ctx, req)

	return h, err
}

func AuthKeyGRPC(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("NOT AUHTHORIZED")
	}
	key := md["key"]
	if len(key) == 0 {
		return errors.New("NOT AUHTHORIZED")
	}
	if key[0] != crypto.SharedKey {
		return errors.New("NOT AUHTHORIZED")
	}
	return nil
}
