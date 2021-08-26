package server

import (
	"context"

	api "github.com/ozonva/ova-obligation-api/pkg/ova-obligation-api"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ObligationServer struct {
	api.UnimplementedObligationRpcServer
	logger *zerolog.Logger
}

func NewObligationRpcServer(logger *zerolog.Logger) api.ObligationRpcServer {
	return &ObligationServer{
		UnimplementedObligationRpcServer: api.UnimplementedObligationRpcServer{},
		logger:                           logger,
	}
}

func (s *ObligationServer) CreateObligation(context context.Context, request *api.CreateObligationRequest) (*api.Obligation, error) {
	s.logger.Info().Msgf("CreateObligation request: %v", request)
	return s.UnimplementedObligationRpcServer.CreateObligation(context, request)
}

func (s *ObligationServer) DescribeObligation(context context.Context, request *api.DescribeObligationRequest) (*api.Obligation, error) {
	s.logger.Info().Msgf("DescribeObligation request: %v", request)
	return s.UnimplementedObligationRpcServer.DescribeObligation(context, request)
}

func (s *ObligationServer) ListObligations(context context.Context, empty *emptypb.Empty) (*api.ListObligationsResponse, error) {
	s.logger.Info().Msgf("ListObligations request: %v", empty)
	return s.UnimplementedObligationRpcServer.ListObligations(context, empty)
}

func (s *ObligationServer) RemoveObligation(context context.Context, request *api.RemoveObligationRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("RemoveObligation request: %v", request)
	return s.UnimplementedObligationRpcServer.RemoveObligation(context, request)
}
