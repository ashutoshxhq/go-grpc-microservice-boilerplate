package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/egnitelabs/engine/internal/config"
	"github.com/egnitelabs/engine/internal/logger"
	"github.com/egnitelabs/engine/internal/middleware"
	"github.com/egnitelabs/engine/internal/server/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer  *grpc.Server
	httpHandler http.Handler
	config      *config.Config
	logger      *zap.Logger
}

func New(conf *config.Config) *Server {
	opts := []grpc.ServerOption{}
	opts = middleware.AddGRPCLogger(logger.Log, opts)

	return &Server{
		grpcServer: grpc.NewServer(opts...),
		config:     conf,
		logger:     logger.Log,
	}
}

func (s *Server) StartGRPC() {
	lis, err := net.Listen("tcp", s.config.GRPCServerAddress())
	if err != nil {
		logger.Log.Sugar().Errorf("failed to listen: %v", err)
	}
	profileServer := user.UserServer{}
	user.RegisterUserServiceServer(s.grpcServer, &profileServer)
	logger.Log.Info("gRPC server starting")

	if err := s.grpcServer.Serve(lis); err != nil {
		logger.Log.Sugar().Errorf("failed to serve grpc: %s", err)
	}
}

func (s *Server) StartHTTP() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register grpc-gateway
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, s.config.GRPCServerAddress(), opts)
	if err != nil {
		log.Fatal(err)
	}

	s.httpHandler = cors.Default().Handler(middleware.AddRequestID(
		middleware.AddHTTPLogger(logger.Log, mux)))

	logger.Log.Info("http server starting")
	if err = http.ListenAndServe(s.config.HTTPServerAddress(), s.httpHandler); err != nil {
		logger.Log.Sugar().Errorf("failed to serve http: %v", err)
	}
}
