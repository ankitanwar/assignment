#### golang assignment

# Steps To Run : 
    go get -u github.com/ankitanwar/assignment
    docker-compose up

# EndPoints

# http://127.0.0.1:8070/ping  -> To check whether the server is running perfectly fine or not
# http://127.0.0.1:8070/detail/:userID -> To fetch the details of the given userID
# http://127.0.0.1:8070/newuser -> to create new user with json data
    {
        "city":"testcity",   city name cannot be empty otherwise it you will receive an error
        "fname":"testfname",  fname cannot be empty
        "phone":1234567899,  phone length should be greater than equal to 10
        "height":12.09,   height cannot be 0
        "is_married":true    bool value
    }


# http://127.0.0.1:8070/details/list -> to fetch the details of all the user given in the list of userID
    pass the json content like this
    {
        "userID":[1,2]
    }