package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"todo_list/model"
	"todo_list/service"
)

type TodoController struct {
	Todocontroller *service.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{
		Todocontroller: service.NewTodoService(),
	}
}

func (t *TodoController) CreateNewList(c *gin.Context) {
	var info *model.CreatNewList
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When ShouldBindJSON",
			"detail": err.Error(),
		})
		return
	}
	info.CreatedAt = time.Now()
	info.UpdatedAt = time.Now()
	err := t.Todocontroller.CreateNewList(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The CreateNewList function.",
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

func (t *TodoController) GetList(c *gin.Context) {
	listID, err := strconv.Atoi(c.Query("list_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When Turn String Into Int",
			"detail": err.Error(),
		})
		return
	}
	list, err := t.Todocontroller.GetList(listID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The GetList function.",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"msg":    "The Program Ran Successfully Without Error.",
		"detail": *list,
	})
}

func (t *TodoController) ShowAllList(c *gin.Context) {
	userID, exisit := c.Get("user_id")
	if !exisit {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   -1,
			"msg":    "Please Login",
			"detail": nil,
		})
		return
	}
	lists, err := t.Todocontroller.ShowAllList(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The ShowAllList function.",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"msg":    "The Program Ran Successfully Without Error.",
		"detail": *lists,
	})
}

func (t *TodoController) UpdateList(c *gin.Context) {
	listID, err := strconv.Atoi(c.Query("list_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When Turn String Into Int",
			"detail": err.Error(),
		})
		return
	}
	var info *model.UpdateList
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When ShouldBindJSON",
			"detail": err.Error(),
		})
		return
	}
	info.UpdatedAt = time.Now()
	err = t.Todocontroller.UpdateList(listID, info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The UpdateList function.",
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

func (t *TodoController) DeleteList(c *gin.Context) {
	listID, err := strconv.Atoi(c.Query("list_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred When Turn String Into Int",
			"detail": err.Error(),
		})
		return
	}
	err = t.Todocontroller.DeleteList(listID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   -1,
			"msg":    "Error Occurred During The Execution Of The DeleteList function.",
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
