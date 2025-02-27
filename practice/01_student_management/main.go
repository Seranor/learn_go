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
	StudentMap map[int]Student
}

func (s StuManagement) viewInfo() {
	fmt.Println("-----------------------------------查看成绩----------------------------------------")
	for _, stu := range s.StudentMap {
		fmt.Printf("学号:%-6d 姓名:%-6s 年龄:%-6d 语文成绩:%-6f 数学成绩:%-6f 英语成绩:%-6f\n",
			stu.Sid, stu.Name, stu.Age, stu.Score["chinese"], stu.Score["math"], stu.Score["english"])
	}
	fmt.Println("---------------------------------------------------------------------------------")
}
func NewStudent(sid int, name string, age int, score map[string]float64) Student {
	return Student{
		Sid:   sid,
		Name:  name,
		Age:   age,
		Score: score,
	}
}
func (s StuManagement) addStudent() {
	var sid, age int
	var name string
	fmt.Println("请分别输入学号,姓名，年龄以空格隔开.如:1001 yuan 18")
	fmt.Scan(&sid, &name, &age)
	_, isExist := s.StudentMap[sid]
	if isExist {
		fmt.Println("学号已存在")
	} else {
		var chinese, math, english float64
		fmt.Println("请分别输入语数英成绩,以空格隔开.如:100 90 80")
		fmt.Scan(&chinese, &math, &english)
		s.StudentMap[sid] = NewStudent(sid, name, age,
			map[string]float64{"chinese": chinese, "math": math, "english": english})
	}
	fmt.Println("添加成功")
}

func (s StuManagement) delStudent() {
	var sid int
	fmt.Println("输入删除的学号")
	fmt.Scan(&sid)
	delete(s.StudentMap, sid)
	fmt.Printf("%d 的学生已删除", sid)
}

func (s StuManagement) getScore() map[int]Student {
	// 读取数据
	bytes, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 反序列化
	var dataMap = make(map[int]Student)
	json.Unmarshal(bytes, &dataMap)
	//fmt.Println("dataMap", dataMap)
	return dataMap
}

func (s StuManagement) updateScore() {
	var sid int
	fmt.Println("请输入你需要更改成绩的学号")
	fmt.Scan(&sid)
	_, isExist := s.StudentMap[sid] // 判断学生是否存在
	if isExist {
		// 更新成绩
		var chinese, math, english float64
		fmt.Println("请分别输入新的语数英成绩,以空格隔开.如:100 90 80")
		fmt.Scan(&chinese, &math, &english)
		s.StudentMap[sid].Score["chinese"] = chinese
		s.StudentMap[sid].Score["math"] = math
		s.StudentMap[sid].Score["english"] = english

	} else {
		fmt.Println("学生不存在")
	}
}

func (s StuManagement) saveInfo() {
	// 序列化
	stuJson, _ := json.Marshal(s.StudentMap)
	//fmt.Println(string(stuJson))
	// 保存数据

	err := os.WriteFile("data.json", stuJson, 0644)
	if err != nil {
		fmt.Println(err)
	}
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
	var studentManager = StuManagement{
		StudentMap: make(map[int]Student),
	}
	var ret = studentManager.getScore()
	if ret != nil {
		studentManager.StudentMap = ret
	}
	// 打印读取到的数据
	//fmt.Println("读取的学生数据：", studentManager)

	for true {
		showMenu()
		var choice int
		fmt.Print("请输入选项>>>")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// 添加学生成绩的代码
			studentManager.addStudent()
		case 2:
			studentManager.viewInfo()
		case 3:
			// 更新学生成绩的代码
			studentManager.updateScore()
		case 4:
			// 删除学生成绩的代码
			studentManager.delStudent()
		case 5:
			// 保存学生数据到文件
			studentManager.saveInfo()
		case 6:
			fmt.Println("退出程序成功")
			os.Exit(0)
		default:
			fmt.Println("输入的选择不合法")
		}
	}
}
