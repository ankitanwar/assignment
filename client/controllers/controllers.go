package controllers

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"sync"

	connect "github.com/ankitanwar/assignment/client/ServerConnect"
	userpb "github.com/ankitanwar/assignment/proto"
	"github.com/gin-gonic/gin"
)

var (
	UserController userControllerInterface = &userControllerStruct{}
	wg             sync.WaitGroup
)

type userControllerInterface interface {
	CreateNewUser(c *gin.Context)
	GetUserDetailByID(c *gin.Context)
	GetDetailsOfListOfUsers(c *gin.Context)
}

type userControllerStruct struct {
}

type userIDList struct {
	UserIDs []int64 `json:"userID"`
}

//Ping : To Check wether server is up and running or not
func Ping(c *gin.Context) {
	c.JSON(http.StatusAccepted, "pong")
}

//CreateNewUser : Controller To Create The New USer
func (controller *userControllerStruct) CreateNewUser(c *gin.Context) {
	details := &userpb.User{}
	if err := c.ShouldBindJSON(details); err != nil {
		c.JSON(http.StatusBadRequest, "error while binding the data")
		return
	}
	req := &userpb.CreateUserRequest{Details: details}
	response, err := connect.Client.CreateUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, response)
}

//GetUserDetailByID : Controller To Fetch the particular user detail by the given ID
func (controller *userControllerStruct) GetUserDetailByID(c *gin.Context) {
	id := c.Param("userID")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to fetch the userID")
		return
	}
	req := &userpb.GetUserDetailsByIDRequest{ID: int64(userID)}
	response, err := connect.Client.GetUserDetail(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)

}

//GetDetailsOfListOfUsers : Controller To Fetch all the details of all the user from the userID list
func (controller *userControllerStruct) GetDetailsOfListOfUsers(c *gin.Context) {
	details := &userIDList{}
	if err := c.ShouldBindJSON(details); err != nil {
		c.JSON(http.StatusBadRequest, "unable to bind the json")
		return
	}
	stream, err := connect.Client.GetDetailsOfListOfUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, "unable to fetch the user details")
		return
	}
	wg.Add(1)
	go func() {
		defer stream.CloseSend()
		for _, id := range details.UserIDs {
			request := &userpb.GetUserDetailsByIDRequest{
				ID: id,
			}
			err = stream.Send(request)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "unable to send the response to the clinet")
			}
		}
	}()

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				return
			} else if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				wg.Done()
				return
			} else {
				c.JSON(http.StatusAccepted, response)
			}
		}
	}()
	wg.Wait()
}
