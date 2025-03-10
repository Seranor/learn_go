package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Book struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Title      string `json:"title" gorm:"type:varchar(255);unique;not null"`
	Price      int64  `json:"price"`
	Is_publish byte   `json:"is_publish" gorm:"default:1"`
}

func main() {
	// 连接MySQL
	dsn := "root:asd123...@tcp(101.126.9.58:3306)/go?charset=utf8mb4&parseTime=True&loc=UTC"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("db:::", db)

	// 创建表
	err = db.AutoMigrate(&Book{})
	if err != nil {
		fmt.Println(err)
	}

	// 添加一条记录
	/*	book01 := Book{Title: "水浒传", Price: 98}
		// 返回的生成对象
		ret := db.Create(&book01)
		fmt.Println(ret)

		fmt.Println(book01.ID) // 接收到了返回来的  ID 等信息都可以拿的到
		fmt.Println(book01.Title)*/

	// 批量插入记录
	/*	books := []Book{
		{Title: "百年孤独", Price: 88},
		{Title: "红楼梦", Price: 89},
		{Title: "西游记", Price: 96}}
	db.Create(&books)*/

	// 每批 500 条创建
	//db.CreateInBatches(books,500)

	// 查询
	// 查询所有
	/*	var book []Book
		db.Find(&book)
		fmt.Println(book)*/

	// 查询第一条或最后一条记录
	/*	var book1 Book
		db.First(&book1)
		fmt.Println(book1)

		var book2 Book
		db.Last(&book2)
		fmt.Println(book2)*/

	// 按条件查询
	// 字符串查询 结构体查询  map查询
	var books []Book
	// db.Where("price > ?", 90).Find(&books)
	//db.Debug().Where("price in ?", []int{98, 99, 100}).Find(&books)
	// db.Debug() 打印SQL执行语句 单条
	//db.Debug().Where("title like ?", "西%").Find(&books)
	//db.Debug().Where("is_publish = ? AND price > ?", "1", 97).Find(&books)

	// 结构体查询
	//db.Debug().Where(&Book{Title: "西游记", Price: 0}).Find(&books) // 只会查询非零的情况

	// map 查询
	db.Debug().Where(map[string]interface{}{"Title": "西游记", "Price": 96}).Find(&books)
	fmt.Println(books)
}
