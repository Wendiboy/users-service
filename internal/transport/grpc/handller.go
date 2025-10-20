package grpc

import (
	"context"
	"log"

	userpb "github.com/Wendiboy/project-protos/proto/user"
	"github.com/Wendiboy/users-service/internal/user"
)

type Handler struct {
	svc user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Printf("LOG CreateUser email=%d", req.Email)

	userRequest := req
	userToCreate := user.User{
		Email:    userRequest.Email,
		Password: "SuperPassword",
	}

	createdUser, err := h.svc.CreateUser(userToCreate)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	// поменять id на string в proto

	response := &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    123456,
			Email: createdUser.Email,
		},
	}
	return response, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Printf("LOG GetUser id=%d", req.Id)

	user, err := h.svc.GetUserById(string(req.Id))
	if err != nil {
		log.Printf("Error get user by id: %v", err)
		return nil, err
	}
	// поменять id на string в proto
	response := &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    123456,
			Email: user.Email,
		},
	}

	log.Printf("Retrieved user: ID= %d, Email=%s", user.Id, user.Email)
	return response, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	log.Printf("UpdateUser request: id=%d, email=%s", req.Id, req.Email)

	userToUpdate := user.User{
		Id:    string(req.Id),
		Email: req.Email,
	}

	updatedUser, err := h.svc.UpdateUser(userToUpdate)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	// поменять id на string в proto
	response := &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    123123,
			Email: updatedUser.Email,
		},
	}

	log.Printf("Updated user %d, Email=%s", updatedUser.Id, updatedUser.Email)
	return response, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	log.Printf("DeleteUser id %d", req.Id)

	err := h.svc.DeleteUser(string(req.Id))
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return nil, err
	}

	response := &userpb.DeleteUserResponse{
		Success: true,
	}

	log.Printf("Deleted user with id: %d", req.Id)
	return response, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {

	users, err := h.svc.GetAllUsers()
	if err != nil {
		log.Printf("Error listing users: %v", err)
		return nil, err
	}

	var pbUsers []*userpb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    123123,
			Email: user.Email,
		})
	}

	response := &userpb.ListUsersResponse{
		Users: pbUsers,
	}

	log.Printf("Retrieved %d users", len(pbUsers))
	return response, nil
}
