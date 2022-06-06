package repo

import (
	pb "github.com/venomuz/project4/UserService/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	GetByID(ID string) (*pb.User, error)
	DeleteByID(ID string) (*pb.GetIdFromUserID, error)
	GetAllUserFromDb(empty *pb.Empty) (*pb.AllUser, error)
}
