package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	pb "github.com/venomuz/project4/PostService/genproto"
	l "github.com/venomuz/project4/PostService/pkg/logger"
	"github.com/venomuz/project4/PostService/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *PostService) PostCreate(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	id1 := uuid.NewV4()
	req.Id = id1.String()
	for _, media := range req.Medias {
		id2 := uuid.NewV4()
		media.Id = id2.String()
	}

	post, err := s.storage.Post().PostCreate(req)
	if err != nil {
		s.logger.Error("Error while inserting post info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert post")
	}
	return post, err
}
func (s *PostService) PostGetByID(ctx context.Context, req *pb.GetIdFromUser) (*pb.Post, error) {
	post, err := s.storage.Post().PostGetByID(req.Id)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting post info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert post")
	}

	return post, err
}
func (s *PostService) PostDeleteByID(ctx context.Context, req *pb.GetIdFromUser) (*pb.OkBOOL, error) {
	post, err := s.storage.Post().PostDeleteByID(req.Id)
	if err != nil {
		s.logger.Error("Error while getting post info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert post")
	}
	return post, err
}
func (s *PostService) PostGetAllPosts(ctx context.Context, req *pb.GetIdFromUser) (*pb.AllPost, error) {
	post, err := s.storage.Post().PostGetAllPosts(req.Id)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting post info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert post")
	}

	return post, err
}
