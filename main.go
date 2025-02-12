package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
    Id string
    Name string
    Dept string
    Email string
}

var Users []User

func main() {

  Users = append(Users, User{
	Id: "1",
	Name: "harshi",
	Dept: "CSE",
	Email: "harshi123@gmail.com",
})
  router := gin.Default()
  router.GET("/users", GetUsers)
  router.POST("/users", AddUsers)
  router.GET("/users/:id", FindUsers)
  router.PUT("/users/:id", UpdateUser)
  router.Run(":8080")
  router.DELETE("/users/:id", DeleteUser)
}

func GetUsers(c *gin.Context){
    c.JSON(http.StatusOK,Users)
}

func AddUsers(c *gin.Context) {
    var userFromFrontend User
	c.ShouldBindJSON(&userFromFrontend)
	Users = append(Users, userFromFrontend)
	c.JSON(http.StatusCreated, userFromFrontend)
}

func FindUsers(c *gin.Context) {
	id := c.Param("id")

	for _, user := range Users {
		if user.Id == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var UpdatedUserfromFronted User
	if err := c.ShouldBindJSON(&UpdatedUserfromFronted); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"invalid request"})
		return

	}

	for i, user := range Users {
		if user.Id == id {
			Users[i] = UpdatedUserfromFronted
			c.JSON(http.StatusOK, UpdatedUserfromFronted)
			return
		}
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range Users{
	    if user.Id == id {
		 Users = append(Users[:i],Users[i+1])
		 c.JSON(http.StatusOK, gin.H{"message":"User Deleted"})
	  return
	}
}
c.JSON(http.StatusNotFound, gin.H{"message":"User not found"})
}

