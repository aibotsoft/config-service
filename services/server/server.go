package server

import (
	"context"
	"github.com/aibotsoft/config-service/services/handler"
	pb "github.com/aibotsoft/gen/confpb"
	"github.com/aibotsoft/micro/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
	cfg     *config.Config
	log     *zap.SugaredLogger
	gs      *grpc.Server
	handler *handler.Handler
	pb.UnimplementedConfServer
}

func NewServer(cfg *config.Config, log *zap.SugaredLogger, handler *handler.Handler) *Server {
	return &Server{cfg: cfg, log: log, handler: handler, gs: grpc.NewServer()}
}

func (s *Server) Serve() error {
	addr := net.JoinHostPort("", s.cfg.Service.GrpcPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "net.Listen error")
	}
	pb.RegisterConfServer(s.gs, s)
	s.log.Infow("gRPC server listens", "service", s.cfg.Service.Name, "addr", addr)
	return s.gs.Serve(lis)
}
func (s *Server) Close() {
	s.log.Debug("begin gRPC server gracefulStop")
	s.gs.GracefulStop()
	s.handler.Close()
	s.log.Debug("end gRPC server gracefulStop")
}
func (s *Server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	//s.log.Info("ping request")
	return &pb.PingResponse{}, nil
}
func (s *Server) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	got, err := s.handler.GetConfig(ctx, req.GetServiceName())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "config not fount for service: %q", req.GetServiceName())
	}
	return &pb.GetConfigResponse{ServiceConfig: got}, nil
}
func (s *Server) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := s.handler.GetAccount(ctx, req.GetServiceName())
	if err != nil {
		s.log.Infow("account not fount for service", "name", req.GetServiceName())
		return nil, status.Errorf(codes.NotFound, "account not fount for service: %q", req.GetServiceName())
	}
	return &pb.GetAccountResponse{Account: account}, nil
}
func (s *Server) GetCurrency(ctx context.Context, req *pb.GetCurrencyRequest) (*pb.GetCurrencyResponse, error) {
	currency, err := s.handler.GetCurrency(ctx)
	if err != nil {
		s.log.Infow("get currency error")
		return nil, status.Errorf(codes.Internal, "get_currency_error")
	}
	return &pb.GetCurrencyResponse{CurrencyList: currency}, nil
}
func (s *Server) GetServices(ctx context.Context, req *pb.GetServicesRequest) (*pb.GetServicesResponse, error) {
	services, err := s.handler.GetServices(ctx)
	if err != nil {
		s.log.Infow("get services error")
		return nil, status.Errorf(codes.Internal, "get services error")
	}
	return &pb.GetServicesResponse{Services: services}, nil
}
func (s *Server) GetNetStatus(ctx context.Context, req *pb.GetNetStatusRequest) (*pb.GetNetStatusResponse, error) {
	return &pb.GetNetStatusResponse{Status: s.handler.NetStatus}, nil
}
