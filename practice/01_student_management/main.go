package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Sid   int                `json:"sid"`
	Name  string             `json:"name"`
	Age   int                `json:"age"`
	Score map[string]float64 `json:"score"` // 保留为 map 类型
}

type StuManagement struct {
	StudentMap map[int]Student `json:"students"` // 结构体中的学生数据存储在 StudentMap 中
}

func (s *StuManagement) viewInfo() {
	fmt.Println("-----------------------------------查看成绩----------------------------------------")
	for _, stu := range s.StudentMap {
		fmt.Printf("学号:%d 姓名:%s 年龄:%d 语文成绩:%f 数学成绩:%f 英语成绩:%f\n", stu.Sid, stu.Name, stu.Age, stu.Score["chinese"], stu.Score["math"], stu.Score["english"])
	}
	fmt.Println("---------------------------------------------------------------------------------")
}

func showMenu() {
	fmt.Print(`
**************************************
  1 添加学生成绩  
  2 查看学生成绩  
  3 更新学生成绩  
  4 删除学生成绩
  5 保存
  6 退出程序    
**************************************
`)
}

func main() {
	var StudentManager = StuManagement{
		StudentMap: make(map[int]Student),
	}

	// 初始化文件，文件是否存在 --> 读取，无 --> 创建
	file, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	content, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	// 检查文件内容是否为空，并进行反序列化
	if len(content) != 0 {
		err := json.Unmarshal(content, &StudentManager) // 文件不为空，反序列化
		if err != nil {
			fmt.Println("unmarshal err:", err)
			return
		}
	} else {
		fmt.Println("文件为空，初始化数据为空。")
	}

	// 打印读取到的数据
	fmt.Println("读取的学生数据：", StudentManager)

	for true {
		showMenu()
		var choice int
		fmt.Print("请输入选项>>>")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// 添加学生成绩的代码
			fmt.Println("添加学生成绩功能尚未实现")
		case 2:
			StudentManager.viewInfo()
		case 3:
			// 更新学生成绩的代码
			fmt.Println("更新学生成绩功能尚未实现")
		case 4:
			// 删除学生成绩的代码
			fmt.Println("删除学生成绩功能尚未实现")
		case 5:
			// 保存学生数据到文件
			updatedContent, err := json.Marshal(StudentManager)
			if err != nil {
				fmt.Println("保存数据到文件失败:", err)
				return
			}
			err = os.WriteFile("data.json", updatedContent, 0644)
			if err != nil {
				fmt.Println("写入文件失败:", err)
			} else {
				fmt.Println("数据已保存")
			}
		case 6:
			fmt.Println("退出程序成功")
			os.Exit(0)
		default:
			fmt.Println("输入的选择不合法")
		}
	}
}
