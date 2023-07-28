package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	middleware "example/socket/Middleware"
	model "example/socket/Model"
)

type User struct {
	db map[string]model.User
}

func (uc *User) Init(group *gin.RouterGroup, db map[string]model.User) {
	uc.db = db
	v1 := group.Group("/v1")
	users := v1.Group("/users")
	users.Use(middleware.GetAuthMiddleware())
	users.GET("/", uc.getUsers)

	v1.POST("/register", uc.postUser)
	v1.PUT("/login", uc.login)
}

func (uc *User) getUsers(c *gin.Context) {
	users := make([]model.User, 0, len(uc.db))

	for _, user := range uc.db {
		users = append(users, user)
	}

	fmt.Printf("users: %d %d", len(users), len(uc.db))

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"users":   users,
		"count":   len(users),
		"size":    len(uc.db),
	})
}

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uc *User) postUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
				"err":     validationErrors.Error(),
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
		}
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
		Key:	  model.GenerateKey(10),
	}

	uc.db[user.Key] = user

	c.JSON(http.StatusOK, gin.H{
		"message": "posted",
	})
}

func (uc *User) login(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err,
		})

		return
	}

	user, ok := uc.db[input.Username]
	if !ok || user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, invalid login or password",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logged in",
		"secret":  "secret-token",
	})
}
