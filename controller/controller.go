package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_web_demo/dao"
	"go_web_demo/models"
	"net/http"
)

func IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreatTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	if err := dao.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}


func GetTodoList(c *gin.Context) {
	//查询所有数据
	todoList, err := dao.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todoList)
	}
}


func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := dao.GetTodoById(id)
	if  err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(todo)
	//将修改的字段的值设置到todo中
	c.BindJSON(todo)
	fmt.Println(todo)
	if err := dao.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}


func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err :=  dao.DeleteTodoById(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, nil)
	}
}
