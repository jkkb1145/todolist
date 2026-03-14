package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/model"
	"todo_list/service"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func (u *UserController) UserRegister(c *gin.Context) {
	var user model.Userregistermodel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When ShouldBindJSON",
			"detail": err.Error(),
		})
		return
	}
	err, userinfo := u.UserService.UserRegister(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The UserRegister function.",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"msg":    "The Program Ran Successfully Without Error.",
		"detail": *userinfo,
	})
}

func (u *UserController) UserLogin(c *gin.Context) {
	var user model.UserLogIn
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When ShouldBindJSON",
			"detail": err.Error(),
		})
		return
	}
	err, userinfo := u.UserService.UserLogIn(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The UserLogIn function.",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"msg":    "The Program Ran Successfully Without Error.",
		"detail": userinfo,
	})
}

func (u *UserController) ChangePassword(c *gin.Context) {
	userid, exisit := c.Get("user_id")
	if !exisit {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   -1,
			"msg":    "Please Login",
			"detail": nil,
		})
		return
	}
	var userCP model.ModelForCP
	if err := c.ShouldBindJSON(&userCP); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When ShouldBindJSON",
			"detail": err.Error(),
		})
		return
	}
	err := u.UserService.ChangePassword(userid.(int), userCP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The ChangePassword function.",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"msg":    "The Program Ran Successfully Without Error.",
		"detail": nil,
	})
}

func (u *UserController) ChangeUserInfo(c *gin.Context) {
	userid, exisit := c.Get("user_id")
	if !exisit {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   -1,
			"msg":    "Please Login",
			"detail": nil,
		})
		return
	}
	var nuserinfo model.ModelForCI
	if err := c.ShouldBindJSON(&nuserinfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When ShouldBindJSON",
			"detail": err.Error(),
		})
		return
	}
	err := u.UserService.ChangeUserInfo(userid.(int), nuserinfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The ChangeUserInfo function.",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"msg":    "The Program Ran Successfully Without Error.",
		"detail": nil,
	})
}
