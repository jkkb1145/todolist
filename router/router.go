package router

import (
	"github.com/gin-gonic/gin"
	"todo_list/controller"
	"todo_list/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	fmt.Println("success")
	//})
	v1 := r.Group("/v1/api")
	{
		user := v1.Group("/user")
		{
			ctl := controller.NewUserController()
			user.POST("/register", ctl.UserRegister)
			user.GET("/login", ctl.UserLogin)
		}
		auth := v1.Group("auth")
		auth.Use(middleware.UserIdentity())
		{
			ctl := controller.NewUserController()
			auth.POST("/changepw", ctl.ChangePassword)
			auth.POST("/changeuserinfo", ctl.ChangeUserInfo)
		}
		todo := v1.Group("/todo")
		todo.Use(middleware.UserIdentity())
		{
			ctl := controller.NewTodoController()
			todo.POST("/createlist", ctl.CreateNewList)
			todo.GET("/getlist", ctl.GetList)
			todo.GET("/showall", ctl.ShowAllList)
			todo.POST("/updatelist", ctl.UpdateList)
			todo.DELETE("/deletelist", ctl.DeleteList)
		}
	}
	return r
}
