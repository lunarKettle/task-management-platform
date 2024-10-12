package grpc_handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"

	"project_management_service/internal/models"
	"project_management_service/internal/repository"
	pb "project_management_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.ProjectManagementServiceServer
	projectRepository repository.ProjectRepository
}

func NewGRPCServer() *GRPCServer {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := &GRPCServer{}
	pb.RegisterProjectManagementServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	db := &repository.Database{}
	db.OpenConnetion()
	server.projectRepository = repository.NewProjectRepository(db)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return server
}

func (s *GRPCServer) GetProject(ctx context.Context, request *pb.ProjectRequest) (*pb.ProjectResponse, error) {
	project, err := s.projectRepository.GetById(request.ProjectId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "project with id %d not found", request.ProjectId)
		}
		return nil, status.Errorf(codes.Internal, "failed to get project: %v", err)
	}

	return &pb.ProjectResponse{
		ProjectId:          project.Id,
		ProjectName:        project.Name,
		ProjectDescription: project.Description,
		StartDate:          timestamppb.New(project.StartDate),
		PlannedEndDate:     timestamppb.New(project.PlannedEndDate),
		ActualEndDate:      timestamppb.New(project.ActualEndDate),
		Status:             project.Status,
		Priority:           project.Priority,
		TeamId:             project.TeamId,
		Budget:             project.Budget,
	}, nil
}

func (s *GRPCServer) CreateProject(ctx context.Context, r *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	newProject := models.Project{
		Name:           r.ProjectName,
		Description:    r.ProjectDescription,
		StartDate:      r.StartDate.AsTime(),
		PlannedEndDate: r.PlannedEndDate.AsTime(),
		ActualEndDate:  r.ActualEndDate.AsTime(),
		Status:         r.Status,
		Priority:       r.Priority,
		TeamId:         r.TeamId,
		Budget:         r.Budget,
	}
	id, err := s.projectRepository.Create(newProject)
	if err != nil {
		return nil, fmt.Errorf("error adding record to database: %w", err)
	}
	return &pb.CreateProjectResponse{ProjectId: id}, nil
}
