package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"请求方式": c.Request.Method,
	})
}
func noRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"error": "not found",
	})
}

type Books struct {
	Title   string `json:"title"`
	Price   int    `json:"price"`
	Publish string `json:"publish"`
}

func books(c *gin.Context) {
	// 获取路径参数
	fmt.Println("year", c.Param("year"))
	fmt.Println("month", c.Param("month"))

	var book01 = Books{"西游记", 55, "清华出版社"}
	var book02 = Books{"三国演义", 56, "北大出版社"}
	var book03 = Books{"水浒传", 57, "交通大学出版社"}
	var book04 = Books{"红楼梦", 58, "中华出版社"}
	var bookLists = []Books{book01, book02, book03, book04}
	c.JSON(200, gin.H{
		"bookLists": bookLists,
	})
}

func articles(c *gin.Context) {
	// 获取请求参数
	// GET
	fmt.Println("GET:::", c.Query("year"))
	fmt.Println("GET:::", c.Query("month"))
	// POST
	fmt.Println("POST:::", c.PostForm("year"))
	fmt.Println("POST:::", c.PostForm("month"))

	// 请求体带 json 格式数据 自己解析
	fmt.Println("JSON:::", c.Request.Body)
	dataJson, _ := io.ReadAll(c.Request.Body)
	var data map[string]int
	fmt.Println(string(dataJson))
	json.Unmarshal(dataJson, &data)
	fmt.Println(data)

	c.JSON(200, gin.H{
		"bookLists": "",
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
			// 重定向
			c.Redirect(http.StatusMovedPermanently, "/ping")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
	}

}

func main() {
	r := gin.Default()
	// 接收任何请求方式
	r.LoadHTMLGlob("./templates/*")
	r.Any("/ping", pong)

	r.Any("/login", login)

	// 没有匹配成功时 走NoRoute
	r.NoRoute(noRoute)

	// 路由参数 请求参数 路径参数
	r.GET("/books/:year/:month", books)
	r.Any("/article", articles)

	// 路由组
	g := r.Group("/api")
	g.GET("/books", books)
	g.Any("/articles", articles)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
