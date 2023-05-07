package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tarunrana0222/social-site-go/config"
	"github.com/tarunrana0222/social-site-go/helpers"
	"github.com/tarunrana0222/social-site-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.BindJSON(&newUser); err != nil {
			fmt.Println("Sign Up", err)
			c.JSON(500, gin.H{"err": err.Error(), "message": "Error while binding json"})
			return
		}
		newUser.ID = primitive.NewObjectID()
		newUser.UserID = newUser.ID.Hex()
		if len(newUser.Password) < 8 || len(newUser.Password) > 16 {
			c.JSON(400, gin.H{"message": "Password should be of length 8-16", "passwordLength": (len(newUser.Password))})
			return
		}
		password, err := helpers.CreatePasswordHash(newUser.Password)
		if err != nil {
			fmt.Println("Sign Up", err)
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}

		newUser.Password = password
		if err := validate.Struct(newUser); err != nil {
			c.JSON(400, gin.H{"err": err.Error(), "message": "Error on validating inputs"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		insertedId, err := config.UsersColl.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}
		c.JSON(200, gin.H{"success": true, "data": insertedId})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		type LoginReqBody struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required"`
		}
		var request LoginReqBody

		if err := c.BindJSON(&request); err != nil {
			c.JSON(500, gin.H{"err": err.Error(), "message": "Error while binding json"})
			return
		}

		if err := validate.Struct(request); err != nil {
			c.JSON(400, gin.H{"err": err.Error(), "message": "Error on validating inputs"})
			return
		}

		var user models.User

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := config.UsersColl.FindOne(ctx, bson.M{"email": request.Email}).Decode(&user); err != nil {
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}
		validPassword := helpers.VerifyPasswordHash(user.Password, request.Password)

		if !validPassword {
			c.JSON(401, gin.H{"message": "Email / Password incorrect"})
			return
		}

		loginToken, err := helpers.CreateToken(user)
		if err != nil {
			c.JSON(500, gin.H{"err": err.Error(), "message": "while creating token"})
			return
		}

		c.JSON(200, gin.H{"message": "Login Success", "AuthToken": loginToken})
	}
}
