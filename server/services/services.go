package service

import (
	"context"
	"errors"
	"io"
	"log"

	userpb "github.com/ankitanwar/assignment/proto"
	db "github.com/ankitanwar/assignment/server/database"
	"github.com/ankitanwar/assignment/server/domain"
)

var (
	UserService userpb.UserServicesServer = &userServiceStruct{}
)

type userServiceStruct struct {
}

//CreateUser : To Create The New User
func (service *userServiceStruct) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	userDetails := req.GetDetails()
	err := domain.CheckDetails(userDetails)
	if err != nil {
		return nil, err
	}
	saveAndGetID, err := db.UserDB.SaveUserDetails(userDetails)
	if err != nil {
		return nil, err
	}
	response := &userpb.CreateUserResponse{
		Message: "User Has Been Created Successfully",
		UserID:  saveAndGetID,
	}
	return response, nil

}

//GetUserDetail : To Get The Details Of The particular user whose ID is given
func (service *userServiceStruct) GetUserDetail(ctx context.Context, req *userpb.GetUserDetailsByIDRequest) (*userpb.GetUserDetailsByIDResponse, error) {
	userID := req.GetID()
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}
	details, err := db.UserDB.GetDetailsByID(userID)
	if err != nil {
		return nil, err
	}
	response := &userpb.GetUserDetailsByIDResponse{
		Details: details,
	}
	return response, nil
}

//GetDetailsOfListOfUsers : To Get the details of all the user from the userID list
func (service *userServiceStruct) GetDetailsOfListOfUsers(stream userpb.UserServices_GetDetailsOfListOfUsersServer) error {
	for {
		details, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Println("Unable To Read The Stream ", err)
			return errors.New("error while reading the stream")
		} else {
			userDetials, err := db.UserDB.GetDetailsByID(details.ID)
			if err != nil {
				return err
			}
			response := &userpb.GetUserDetailsByIDResponse{
				Details: userDetials,
			}
			err = stream.Send(response)
			if err != nil {
				return errors.New("unable to send the response ")
			}
		}
	}
}
