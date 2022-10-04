package users

import (
	"net/http"
	"strconv"

	"github.com/JayCardenas/bookstore-users-api/domain/users"
	"github.com/JayCardenas/bookstore-users-api/services"
	"github.com/JayCardenas/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	// // one way of handling the json request from the controller
	// // create struct object
	// var user users.User

	// // rea
	// bytes, err := ioutil.ReadAll(c.Request.Body)

	// if err != nil {
	// 	//TODO: Handle error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Print(err.Error())
	// 	//TODO: Handle json error
	// 	return
	// }

	// result, saveErr := services.CreateUser(user)

	// if saveErr != nil {
	// 	//TODO: Handle user creation error
	// 	return
	// }

	// c.JSON(http.StatusCreated, result)

	// c.String(http.StatusNotImplemented, "implement me!")

	// **************************
	// Another way of handling
	//
	// Controller does not contain any business logic for scalability
	// **************************
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {

		// TODO (COMPLETED): Return bad request to the caller
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)

		return
	}

	result, saveErr := services.CreateUser(user)

	// if theres any error tyring to process the user i.e bad request, internal server error, or any
	// tyoe of error that may occur in the servicing the user
	if saveErr != nil {
		// TODO: handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context) {
	// c.String(http.StatusNotImplemented, "implement me!")

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
