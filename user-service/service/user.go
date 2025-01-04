package service

import (
	"context"
	"project/user-service/model"
	pb "project/user-service/proto"
	"project/user-service/repository"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type ServiceUser struct {
	repo repository.Repository
	log  *zap.Logger
	pb.UnimplementedUserServiceServer
}

func NewServiceUser(repo repository.Repository, log *zap.Logger) *ServiceUser {
	return &ServiceUser{repo: repo, log: log}
}

func (s *ServiceUser) GetAllUsers(ctx context.Context, req *pb.Empty) (*pb.UsersList, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	var filter bool
	isOnline, ok := md["filter"]
	if ok && isOnline[0] != "" {
		// filter = bool(isOnline[0])
		// fmt.Printf("%+v\n", md)
		// fmt.Println(filter, len(isOnline))
		filter = true
	}
	users, err := s.repo.User.GetAllUsers(filter)
	if err != nil {
		return nil, err
	}
	var usersPb []*pb.User
	for _, user := range users {
		usersPb = append(usersPb, &pb.User{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			IsOnline:  user.IsOnline,
		})

	}
	return &pb.UsersList{Users: usersPb}, nil
}
func (s *ServiceUser) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponseSuccess, error) {
	var user model.User
	user.Email = req.Email
	err := s.repo.User.Insert(&user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponseSuccess{Message: "Create User Succes"}, nil
}
func (s *ServiceUser) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponseSuccess, error) {
	var user model.User
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		user.Email = md["email"][0]
	}
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	err := s.repo.User.UpdateProfile(&user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponseSuccess{Message: "Update Profile Succes"}, nil
}
