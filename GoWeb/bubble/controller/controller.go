package controller

import (
	"gin_webapp/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

/*增加待办*/
func CreateATodo(c *gin.Context) {
	// 获取请求数据
	var todo models.Todo
	c.BindJSON(&todo)
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

/*查看所有todo*/
func GetTodoList(c *gin.Context) {
	todoList, err := models.GetTodoList()
	print(todoList)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
	} else {
		todo, err := models.GetATodo(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.Bind(&todo)
		if models.UpdateATodo(todo); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}

	}
}

func DeleteATodo(c *gin.Context) {
	c.Request.ParseForm()
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"eroor": "无效的id"})
	} else {
		if err := models.DeleteATodo(id); err != nil {
			c.JSON(http.StatusOK, gin.H{"eroor": err.Error()})
		}else {
			c.JSON(http.StatusOK, gin.H{id:"deleted"})
		}
	}
}
