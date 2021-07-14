package routers

import (
	"github.com/gin-gonic/gin"
	"go_web_demo/controller"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	//设置静态文件路径
	r.Static("/static", "static")
	//解析静态模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	group := r.Group("/v1")
	{
		//添加待办事项
		group.POST("/todo", controller.CreatTodo)

		group.GET("/todo", controller.GetTodoList)

		group.PUT("/todo/:id", controller.UpdateTodo)

		group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
