package rest

import (
	"net/http"

	"github.com/alehechka/swagger-api/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {object} types.UsersResponse
// @Failure 404 {object} types.Errors
// @Router /users [get]
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": []types.User{}})
}

// GetUserByID ... Get the user by id
// @Summary Get one user
// @Description get user by ID
// @Tags Users
// @Param id path string true "User ID" Format(uuid)
// @Success 200 {object} types.UserResponse
// @Failure 400,404 {object} types.Errors
// @Router /users/{id} [get]
func getUserByID(c *gin.Context) {
	userID := uuid.MustParse(c.Param("id"))

	user := types.User{
		ID: userID,
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// CreateUser ... Create User
// @Summary Create new user based on parameters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body types.User true "User Data"
// @Success 201 {object} types.UserResponse
// @Failure 400,500 {object} types.Errors
// @Router /users [post]
func createUser(c *gin.Context) {
	c.JSON(http.StatusCreated, types.User{})
}

// UpdateUser ... Update User
// @Summary Update existing user based on parameters
// @Description Update existing user
// @Tags Users
// @Accept json
// @Param id path string true "User ID" Format(uuid)
// @Param user body types.User true "User Data"
// @Success 200 {object} types.UserResponse
// @Failure 400,500 {object} types.Errors
// @Router /users/{id} [put]
func updateUser(c *gin.Context) {
	userID := uuid.MustParse(c.Param("id"))

	user := types.User{
		ID: userID,
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// DeleteUser ... Delete User
// @Summary Delete existing user based on parameters
// @Description Delete existing user
// @Tags Users
// @Param id path string true "User ID" Format(uuid)
// @Success 204
// @Failure 400,500 {object} types.Errors
// @Router /users/{id} [delete]
func deleteUser(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
