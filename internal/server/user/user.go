package user

import "context"

type UserServer struct {
	UnimplementedUserServiceServer
}

func (s UserServer) GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {

	return &GetUsersResponse{
		Err: "Test string",
	}, nil
}
