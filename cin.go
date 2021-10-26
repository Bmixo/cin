package cin

import (
	"fmt"
	"github.com/Bmixo/cin/middleware/auth"
	"github.com/Bmixo/cin/middleware/recovery"
	"github.com/Bmixo/cin/middleware/zap"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Plugin struct {
	Name string
	Api  func(s *grpc.Server)
}

func Server(au *auth.Auth) *grpc.Server {
	server := grpc.NewServer(
		//cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.StreamServerInterceptor(au.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(au.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)
	au.Api(server)
	i := 0
	for k, _ := range au.NoVerifyMethodMap {
		i++
		fmt.Println("[", i, "]", "by pass", k, "...")
	}
	reflection.Register(server)
	return server
}
