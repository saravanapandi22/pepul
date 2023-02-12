package controllers

import (
	"example/pepel/models"
	"example/pepel/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func New (UserService services.UserService) UserController {
	return UserController{
		UserService: UserService,
	}
}

func (uc *UserController) CreateUser (ctx *gin.Context) {
	var UserData models.User
	if err := ctx.ShouldBindJSON(&UserData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&UserData)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (uc *UserController) GetAll (ctx *gin.Context) {
	UserData, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, UserData)
}

func (uc *UserController) UpdateUser (ctx *gin.Context) {
	var UserData models.User
	if err := ctx.ShouldBindJSON(&UserData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&UserData)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (uc *UserController) DeleteUser (ctx *gin.Context) {
	UserName := ctx.Param("name")
	err := uc.UserService.DeleteUser(&UserName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (uc *UserController) RegisteredRoutes (rr *gin.RouterGroup) {
	UserRoutes := rr.Group("/userApi")
	UserRoutes.POST("/create", uc.CreateUser)
	UserRoutes.GET("/getAll", uc.GetAll)
	UserRoutes.PATCH("/update", uc.UpdateUser)
	UserRoutes.DELETE("/delete/:name", uc.DeleteUser)
}
