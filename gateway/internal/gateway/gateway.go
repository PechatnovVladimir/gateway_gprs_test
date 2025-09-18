package gateway

import (
	"context"
	authPB "github.com/PechatnovVladimir/gateway_gprs_test/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := authPB.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "auth_service:50051", opts)
	if err != nil {
		return err
	}

	mainMux := http.NewServeMux()
	mainMux.Handle("/v1/", mux)

	port := "8080"
	log.Printf("🚀 Gateway listening on :%s", port)
	return http.ListenAndServe(":"+port, mainMux)

}
