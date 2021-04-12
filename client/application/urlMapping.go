package application

import "github.com/ankitanwar/assignment/client/controllers"

func urlMapping() {
	router.POST("/newuser", controllers.UserController.CreateNewUser)
	router.POST("/details/list", controllers.UserController.GetDetailsOfListOfUsers)
	router.GET("/detail/:userID", controllers.UserController.GetUserDetailByID)
	router.GET("/ping", controllers.Ping)
}
