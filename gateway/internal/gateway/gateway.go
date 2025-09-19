package gateway

import (
	"context"
	authPB "github.com/PechatnovVladimir/gateway_gprs_test/auth/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func RunRest() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := authPB.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
