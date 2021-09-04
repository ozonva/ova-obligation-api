package server

import (
	"context"
	"database/sql"
	"errors"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-obligation-api/internal/entity"
	"github.com/ozonva/ova-obligation-api/internal/repo"
	"github.com/ozonva/ova-obligation-api/internal/utils"
	api "github.com/ozonva/ova-obligation-api/pkg/ova-obligation-api"
	queue "github.com/ozonva/ova-obligation-api/pkg/ova-obligation-producer"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ObligationServer struct {
	api.UnimplementedObligationRpcServer
	logger     *zerolog.Logger
	repository repo.Repo
	producer   queue.Producer
}

func NewObligationRpcServer(logger *zerolog.Logger, repository repo.Repo, producer queue.Producer) api.ObligationRpcServer {
	return &ObligationServer{
		UnimplementedObligationRpcServer: api.UnimplementedObligationRpcServer{},
		logger:                           logger,
		repository:                       repository,
		producer:                         producer,
	}
}

func (s *ObligationServer) MultiCreateObligation(context context.Context, request *api.MultiCreateObligationRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("MultiCreateEntity request: %v", request)

	tracer := opentracing.GlobalTracer()

	span := tracer.StartSpan("Multicreate request")
	defer span.Finish()

	obligations := []*entity.Obligation{}
	for _, v := range request.Obligations {
		obligations = append(obligations, &entity.Obligation{
			Title:       v.Title,
			Description: v.Description,
		})
	}

	for _, chunkedObligations := range utils.ChunkObligationPointers(obligations, 5) {
		createSpan := tracer.StartSpan("adding chunked obligations", opentracing.ChildOf(span.Context()))
		err := s.repository.AddEntities(chunkedObligations)
		if err != nil {
			s.logger.Error().Err(err).Msg("")
			continue
		}

		for _, v := range chunkedObligations {
			s.producer.Publish(v.ID, queue.Created)
		}

		createSpan.Finish()
	}

	return &emptypb.Empty{}, nil
}

func (s *ObligationServer) UpdateObligation(context context.Context, request *api.UpdateObligationRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msgf("UpdateEntity request: %v", request)

	obligation, err := s.repository.DescribeEntity(uint64(request.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Not found")
		}

		s.logger.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, "Internal")
	}

	obligation.Description = request.Description
	obligation.Title = request.Title

	err = s.repository.UpdateEntity(obligation)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, "Internal")
	}

	s.producer.Publish(uint(request.Id), queue.Updated)

	return &emptypb.Empty{}, nil
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

	s.producer.Publish(obligation.ID, queue.Created)

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

	s.producer.Publish(uint(request.Id), queue.Created)

	return &emptypb.Empty{}, nil
}
