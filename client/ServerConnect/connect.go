package connect

import (
	"fmt"

	userpb "github.com/ankitanwar/assignment/proto"
	"google.golang.org/grpc"
)

var (
	Client userpb.UserServicesClient
	CC     *grpc.ClientConn
)

//ConnectServer : To Connect To the gRPC server
func ConnectServer() {
	opts := grpc.WithInsecure()
	var err error
	CC, err = grpc.Dial("localhost:8082", opts)
	if err != nil {
		fmt.Println("Error while connection to the server", err.Error())
		panic(err)
	}
	Client = userpb.NewUserServicesClient(CC)
	fmt.Println("Connection to Server is successfull")
}
