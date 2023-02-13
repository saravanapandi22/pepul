package main

import (
	"context"
	"example/pepel/controllers"
	"example/pepel/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	Server *gin.Engine
	UserService services.UserService
	UserController controllers.UserController
	ctx context.Context
	MongoCollection *mongo.Collection
	mongoClient *mongo.Client
	err error
)

func init() {
	ctx = context.TODO()

	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatalln(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("mongo db connection success")

	MongoCollection = mongoClient.Database("userDetails").Collection("usersData")
	UserService = services.NewUserService(MongoCollection, ctx)
	UserController = controllers.New(UserService)
	Server = gin.Default()
}

func UserMiddlewareValidator() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		var UserInput models.User
		if err := gctx.ShouldBindJSON(&UserInput); err == nil {
			UserValidate := validator.New()
			if err := UserValidate.Struct(&UserInput); err != nil {
				gctx.JSON(http.StatusBadRequest, gin.H{
					"error":err.Error(),
				})
				gctx.Abort()
				return
			}
		}
		gctx.Next()
	}
}

func main() {
	defer mongoClient.Disconnect(ctx)
	Server.Use(UserMiddlewareValidator())
	Server.MaxMultipartMemory = 8 << 20
	path := Server.Group("/p1")
	UserController.RegisteredRoutes(path)
}
