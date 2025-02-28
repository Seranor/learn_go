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
	Score map[string]float64 `json:"score"`
}
type StudentManagement struct {
	Students map[int]Student
}

func (s StudentManagement) ViewInfo() {
	fmt.Println("学生成绩信息")
	fmt.Println("--------------------------------------------------------------------------")
	for _, stu := range s.Students {
		fmt.Printf("学号:%-8d 姓名:%-8s 年龄:%-8d 语文成绩:%-8.2f 数学成绩:%-8.2f 英语成绩:%-8.2f\n",
			stu.Sid, stu.Name, stu.Age, stu.Score["chinese"], stu.Score["math"], stu.Score["english"])
	}
	fmt.Println("---------------------------------------------------------------------------")
}
func (s StudentManagement) getScore() map[int]Student {
	// 读取文件
	bytes, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var dataMap map[int]Student
	err = json.Unmarshal(bytes, &dataMap)
	if err != nil {
		fmt.Println(err)
	}
	return dataMap

	// 反序列化
}

func NewStudent(sid int, name string, age int, score map[string]float64) Student {
	return Student{
		Sid:   sid,
		Name:  name,
		Age:   age,
		Score: score}
}

func (s StudentManagement) addStudent() {
	var sid, age int
	var name string
	fmt.Println("请输入添加的学号 姓名 年龄，如 1001 xxx 19")
	fmt.Scan(&sid, &name, &age)
	_, isExist := s.Students[sid]
	if isExist {
		fmt.Println("输入的学号已存在")
	} else {
		var chinese, math, english float64
		fmt.Println("请输入学生语数外成绩，如77 88 99")
		fmt.Scan(&chinese, &math, &english)
		s.Students[sid] = NewStudent(sid, name, age,
			map[string]float64{"chinese": chinese, "math": math, "english": english})
	}
}

func (s StudentManagement) delStudent() {
	var sid int
	fmt.Println("请输入学号")
	fmt.Scan(&sid)
	_, isExist := s.Students[sid]
	if isExist {
		delete(s.Students, sid)
	} else {
		fmt.Println("删除的学号不存在")
	}
}
func (s StudentManagement) updateStudent() {
	var sid int
	fmt.Println("请输入学号")
	fmt.Scan(&sid)
	_, isExist := s.Students[sid]
	if isExist {
		var chinese, math, english float64
		fmt.Println("输入新的学生语数外成绩，如77 88 99")
		fmt.Scan(&chinese, &math, &english)
		s.Students[sid].Score["chinese"] = chinese
		s.Students[sid].Score["math"] = math
		s.Students[sid].Score["english"] = english
		fmt.Println("更新成绩成功")
	} else {
		fmt.Println("更新的学号不存在")
	}
}

func (s StudentManagement) saveStudent() {
	// 序列化
	stuJson, _ := json.Marshal(s.Students)
	// 存入文件
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
	var students = StudentManagement{
		Students: make(map[int]Student),
	}
	var ret = students.getScore()
	if ret != nil {
		students.Students = ret
	}

	for true {
		showMenu()
		var choice int
		fmt.Println("请输入选项功能")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			students.addStudent()
		case 2:
			students.ViewInfo()
		case 3:
			students.updateStudent()
		case 4:
			students.delStudent()
		case 5:
			students.saveStudent()
		case 6:
			fmt.Println("退出程序成功")
			os.Exit(0)
		default:
			fmt.Println("输入不正确")
		}
	}
}
