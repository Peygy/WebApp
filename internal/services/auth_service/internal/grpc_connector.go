package internal

import (
	"context"
	"time"

	"github.com/peygy/nektoyou/internal/pkg/grpc"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	pb "github.com/peygy/nektoyou/internal/pkg/protos/graph_auth"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/services"
)

func InitAuthGrpcServer(server *grpc.GrpcServer, tokenManager services.ITokenManager, logger logger.ILogger) {
	grpcServer := &grpcServer{tokenManager: tokenManager, logger: logger}
	pb.RegisterSignInServiceServer(server.Engine, grpcServer)
}

type grpcServer struct {
	pb.UnimplementedSignInServiceServer
	tokenManager services.ITokenManager
	logger       logger.ILogger
}

func (s *grpcServer) GeneratePairOfTokens(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponce, error) {
	at, err := s.tokenManager.NewAccessToken(in.Username, time.Minute*5)
	if err != nil {
		s.logger.Error("error during creation of access token: " + err.Error())
		return nil, err
	}

	rt, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		s.logger.Error("error during creation of refresh token: " + err.Error())
		return nil, err
	}

	return &pb.SignInResponce{AccessToken: at, RefreshToken: rt}, nil
}
