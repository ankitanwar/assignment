package db

import (
	"errors"
	"log"

	userpb "github.com/ankitanwar/assignment/proto"
)

const (
	insertUser         = "INSERT INTO users(fname,city,phone,height,is_married)VALUES(?,?,?,?,?) "
	getUserDetailsByID = "SELECT fname,city,phone,height,is_married FROM users WHERE id=?;"
	createTable        = "CREATE TABLE IF NOT EXISTS `users` (`id` int NOT NULL AUTO_INCREMENT,`fname` varchar(45) NOT NULL,`city` varchar(45) NOT NULL,`phone` int NOT NULL,`height` float NOT NULL,`is_married` tinyint NOT NULL, PRIMARY KEY (`id`))"
)

var (
	UserDB userDBinterface = &userDBstruct{}
)

type userDBinterface interface {
	SaveUserDetails(*userpb.User) (int64, error)
	GetDetailsByID(int64) (*userpb.User, error)
}

type userDBstruct struct {
}

func makeTable() error {
	stmt, err := Client.Prepare(createTable)
	if err != nil {
		log.Println("Unable To Prepare The createTable Statement", err)
		return errors.New("unable to save the user details")
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println("error while creating the table", err)
		return errors.New("error while creating the table")
	}
	return nil
}

//SaveUSerDetials: Save The User Details Into The Database and returns userID or error
func (db *userDBstruct) SaveUserDetails(details *userpb.User) (int64, error) {
	stmt, err := Client.Prepare(insertUser)
	if err != nil {
		log.Println("Unable To Prepare The insertUser Statement", err)
		return -1, errors.New("unable to save the user details")
	}
	defer stmt.Close()
	insert, err := stmt.Exec(details.GetFname(), details.GetCity(), details.GetPhone(), details.GetHeight(), details.GetIsMarried())
	if err != nil {
		return -1, errors.New("unable to save the details")
	}
	userID, _ := insert.LastInsertId()
	return userID, nil

}

//GetDetialsByID: To Fetch The Details Of The Given User ID
func (db *userDBstruct) GetDetailsByID(userID int64) (*userpb.User, error) {
	details := &userpb.User{}
	stmt, err := Client.Prepare(getUserDetailsByID)
	if err != nil {
		log.Println("Unable to prepare the getDetialsByID statement", err)
		return nil, errors.New("unable to fetch the user details by the given id")
	}
	defer stmt.Close()
	result := stmt.QueryRow(userID)
	if err := result.Scan(&details.Fname, &details.City, &details.Phone, &details.Height, &details.IsMarried); err != nil {
		return nil, errors.New("unable to fetch the user details")
	}
	return details, nil
}
