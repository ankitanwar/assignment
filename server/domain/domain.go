package domain

import (
	"errors"
	"math"

	userpb "github.com/ankitanwar/assignment/proto"
)

//CheckDetials : To verify whether the given details are correct or not
func CheckDetails(details *userpb.User) error {
	if details.GetFname() == "" {
		return errors.New("first name cannot be empty")
	} else if details.GetCity() == "" {
		return errors.New("city cannot be empty")
	} else if math.Ceil(math.Log10(float64(details.GetPhone()))) < 10 { //phone number should be 10 digit long
		return errors.New("phone number should be atleast 10 digit long")
	} else if details.GetHeight() == 0 {
		return errors.New("height cannot be zero")
	}
	return nil
}
