package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 上下文对象 *gin.Context
func index(c *gin.Context) {
	var name = "huahua"

	ip := c.Request.RemoteAddr

	c.HTML(http.StatusOK, "index.html", gin.H{
		"user": name,
		"ip":   ip,
	})
}
func login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", nil)
	} else if c.Request.Method == "POST" {
		// 获取用户数据
		name := c.PostForm("user")
		password := c.PostForm("pwd")
		//fmt.Println(name, password)
		if name == "liu" && password == "123456" {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "login success",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "login fail"})
		}
	}

}

type Books struct {
	Title   string `json:"title"`
	Price   int    `json:"price"`
	Publish string `json:"publish"`
}

func books(c *gin.Context) {
	var book01 = Books{"西游记", 55, "清华出版社"}
	var book02 = Books{"三国演义", 56, "北大出版社"}
	var book03 = Books{"水浒传", 57, "交通大学出版社"}
	var book04 = Books{"红楼梦", 58, "中华出版社"}

	var bookList = []Books{book01, book02, book03, book04}
	fmt.Println(bookList)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": bookList,
	})
}

func main() {
	// 获取默认路由对象
	router := gin.Default()

	// 加载所有的模板文件
	router.LoadHTMLGlob("./templates/*")

	// 构建路由分发函数
	router.GET("/index", index)
	router.GET("/login", login)
	router.POST("/login", login)
	router.GET("/books", books)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
