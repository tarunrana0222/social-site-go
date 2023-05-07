package controllers

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tarunrana0222/social-site-go/config"
	"github.com/tarunrana0222/social-site-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var post models.Post
		post.ID = primitive.NewObjectID()
		post.UserID = ctx.GetString("userId")

		post.Title = ctx.PostForm("title")
		if post.Title == "" {
			ctx.JSON(400, gin.H{"message": "Title cant be empty"})
			return
		}
		post.Description = ctx.PostForm("description")

		// getting value form file
		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error while retrieving file", "err": err.Error()})
			return
		}

		supportedExtentions := []string{".jpg", ".png"}

		filename := filepath.Base(file.Filename)
		fileExt := filepath.Ext(filename)
		isExtSupported := false
		for _, value := range supportedExtentions {
			if fileExt == value {
				isExtSupported = true
			}
		}
		if !isExtSupported {
			ctx.JSON(400, gin.H{"message": "File ext not supported "})
			return
		}

		if err := ctx.SaveUploadedFile(file, "public/"+filename); err != nil {
			ctx.JSON(500, gin.H{"message": "error while saving file", "err": err.Error()})
			return
		}

		post.ImagePath = "public/" + filename

		mongoCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		insertedId, err := config.PostsColl.InsertOne(mongoCtx, post)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Post not saved", "err": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": post, "insertedId": insertedId})
	}
}

func GetAllPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var posts []models.Post
		mongoCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		id := ctx.GetString("userId")
		fmt.Print(id)

		cursor, err := config.PostsColl.Find(mongoCtx, bson.M{"userId": id})
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Error while retriving post", "err": err.Error()})
			return
		}

		for cursor.Next(context.TODO()) {
			var result models.Post
			if err := cursor.Decode(&result); err != nil {
				fmt.Println(err.Error())
			}
			posts = append(posts, result)
		}
		ctx.JSON(200, gin.H{"data": posts})
	}
}

func DeletePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Error while converting into objectId", "err": err.Error()})
			return
		}
		var post models.Post
		mongoCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		err = config.PostsColl.FindOneAndDelete(mongoCtx, bson.M{"_id": id}).Decode(&post)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Error while Deleting post", "err": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"deleted": post})
	}
}

func GetSinglePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Error while converting into objectId", "err": err.Error()})
			return
		}
		var post models.Post
		mongoCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := config.PostsColl.FindOne(mongoCtx, bson.M{"_id": id}).Decode(&post); err != nil {
			ctx.JSON(500, gin.H{"message": "Error while retriving post", "err": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": post})
	}
}
