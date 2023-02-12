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
func main() {
	defer mongoClient.Disconnect(ctx)

	path := Server.Group("/p1")
	UserController.RegisteredRoutes(path)
}