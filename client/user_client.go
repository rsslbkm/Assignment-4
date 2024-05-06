package main

import (
	"context"
	pb "path/to/your/proto/package"

	"google.golang.org/grpc"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: pb.NewUserServiceClient(conn),
	}
}

func (c *UserClient) AddUser(user *pb.User) (*pb.User, error) {
	// Call AddUser RPC
	return c.client.AddUser(context.Background(), user)
}

func (c *UserClient) GetUser(id int32) (*pb.User, error) {
	// Call GetUser RPC
	return c.client.GetUser(context.Background(), &pb.GetUserRequest{Id: id})
}

func (c *UserClient) ListUsers() ([]*pb.User, error) {
	// Call ListUsers RPC
	stream, err := c.client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}
	var users []*pb.User
	for {
		user, err := stream.Recv()
		if err != nil {
			break
		}
		users = append(users, user)
	}
	return users, nil
}
