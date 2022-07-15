package controllers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/macduyhai/loadBalancingGrpcApi/dtos"
	"github.com/macduyhai/loadBalancingGrpcApi/models"
	"github.com/macduyhai/loadBalancingGrpcApi/services"
	"github.com/macduyhai/loadBalancingGrpcApi/utilitys"
)

type Controller struct {
	userService services.UserService
}

func NewController(provider services.Provider) Controller {
	return Controller{
		userService: provider.GetUserService(),
	}
}
func (ctl *Controller) Login(context *gin.Context) {
	// log.Println(context.Request.Header)
	var request dtos.LoginRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Decode json login request error: " + err.Error())
		utilitys.ResponseError400(context, "Login error")
		return
	}

	token, err := ctl.userService.Login(request)
	if err != nil {
		log.Println("Create token login Error:" + err.Error())
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, token, "Login success")
	}
}

func (ctl *Controller) Create(context *gin.Context) {
	var request dtos.AddRequest

	err := context.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Decode json create request error: " + err.Error())
		utilitys.ResponseError400(context, err.Error())
		return
	}
	// timeNow := utilitys.TimeIn("Asia/Ho_Chi_Minh")
	timeNow := time.Now()
	us := models.User{
		Username:   request.Username,
		Password:   request.Password,
		Fullname:   request.Fullname,
		Salary:     request.Salary,
		Active:     1,
		CreateTime: &timeNow,
	}
	user, err := ctl.userService.Create(us)
	if err != nil {
		log.Println("Create user error:" + err.Error())
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, user, "Create success")
	}
}

func (ctl *Controller) Edit(context *gin.Context) {
	var request dtos.EditRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Decode json edit request error: " + err.Error())
		utilitys.ResponseError400(context, "Edit error")
		return
	}
	// timeNow := utilitys.TimeIn("Asia/Ho_Chi_Minh")
	timeNow := time.Now()
	us := models.User{
		Username:   request.Username,
		Fullname:   request.Fullname,
		Salary:     request.Salary,
		UpdateTime: &timeNow,
	}
	user, err := ctl.userService.Edit(us, request.Token)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	} else {
		utilitys.ResponseSuccess200(context, user, "Edit success")
	}

}
func (ctl *Controller) Delete(context *gin.Context) {
	var request dtos.DeleteRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Decode json edit request error: " + err.Error())
		utilitys.ResponseError400(context, "Edit error")
		return
	}
	// timeNow := utilitys.TimeIn("Asia/Ho_Chi_Minh")
	timeNow := time.Now()
	us := models.User{
		Username:   request.Username,
		UpdateTime: &timeNow,
	}
	err = ctl.userService.Delete(us, request.Token)
	if err != nil {
		log.Println("Delete user error:" + err.Error())
		utilitys.ResponseError400(context, err.Error())
		return
	} else {
		utilitys.ResponseSuccess200(context, "", "Delete success")
	}

}

func (ctl *Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong Pong",
	})
}
func stringYYYYMMDD2Time(input string) (*time.Time, error) {
	layout := "2006-01-02"
	result, err := time.Parse(layout, input)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
