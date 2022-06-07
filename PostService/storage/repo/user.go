package repo

import (
	pb "github.com/venomuz/project4/PostService/genproto"
)

//PostStorageI ...
type PostStorageI interface {
	PostCreate(*pb.Post) (*pb.Post, error)
	PostGetByID(ID string) (*pb.Post, error)
	PostDeleteByID(ID string) (*pb.OkBOOL, error)
	PostGetAllPosts(ID string) (*pb.AllPost, error)
}
