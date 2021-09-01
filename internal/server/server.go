package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ozonva/ova-obligation-api/internal/entity"
	"github.com/ozonva/ova-obligation-api/internal/repo"
	api "github.com/ozonva/ova-obligation-api/pkg/ova-obligation-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ObligationServer struct {
	api.UnimplementedObligationRpcServer
	logger     *zerolog.Logger
	repository repo.Repo
}

func NewObligationRpcServer(logger *zerolog.Logger, repository repo.Repo) api.ObligationRpcServer {
	return &ObligationServer{
		UnimplementedObligationRpcServer: api.UnimplementedObligationRpcServer{},
		logger:                           logger,
		repository:                       repository,
	}
}

func (s *ObligationServer) CreateObligation(context context.Context, request *api.CreateObligationRequest) (*api.CreateObligationResponce, error) {
	s.logger.Info().Msgf("CreateObligation request: %v", request)

	obligation := entity.Obligation{
		Title:       request.Title,
		Description: request.Description,
	}

	err := s.repository.AddEntity(&obligation)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, "Internal")
	}

	return &api.CreateObligationResponce{
		Id: uint32(obligation.ID),
	}, err
}

func (s *ObligationServer) DescribeObligation(context context.Context, request *api.DescribeObligationRequest) (*api.DescribeObligationResponse, error) {
	s.logger.Info().Msgf("DescribeObligation request: %v", request)

	result, err := s.repository.DescribeEntity(uint64(request.Id))
	if err != nil {
		s.logger.Error().Err(err).Msg("")

		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Not found")
		}

		return nil, status.Error(codes.Internal, "Internal")
	}

	return &api.DescribeObligationResponse{
		Id:          uint32(result.ID),
		Title:       result.Title,
		Description: result.Description,
	}, nil
}

func (s *ObligationServer) ListObligations(context context.Context, request *api.ListObligationsRequest) (*api.ListObligationsResponse, error) {
	s.logger.Info().Msgf("ListObligations request: %v", request)

	result, err := s.repository.ListEntities(uint64(request.Limit), uint64(request.Offset))
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, "Internal")
	}

	list := make([]*api.Obligation, 0)
	for _, v := range result {
		list = append(list, &api.Obligation{
			Id:          uint32(v.ID),
			Title:       v.Title,
			Description: v.Description,
		})
	}

	return &api.ListObligationsResponse{
		Items: list,
	}, nil
}

func (s *ObligationServer) RemoveObligation(context context.Context, request *api.RemoveObligationRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("RemoveObligation request: %v", request)

	err := s.repository.RemoveEntity(uint64(request.Id))
	if err != nil {
		s.logger.Error().Err(err).Msg("")

		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Not found")
		}

		return nil, status.Error(codes.Internal, "Internal")
	}

	return &emptypb.Empty{}, nil
}
