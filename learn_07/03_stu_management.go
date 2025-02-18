package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	info := `
	1. 添加学生成绩
	2. 查看学生成绩
	3. 更新学生成绩
	4. 删除学生信息
	5. 退出
`

	students := make(map[int]map[string]string)
	reader := bufio.NewReader(os.Stdin) // 创建一个读取器

	for {
		fmt.Println(info)
		fmt.Print("请输入选项>>>")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("添加学生成绩")
			fmt.Println("请输入学生信息（学号 姓名 语文 数学 英语，例如：1001 aaa 90 90 90）：")
			studentInput, _ := reader.ReadString('\n')     // 读取整行输入，包含空格
			studentInput = strings.TrimSpace(studentInput) // 去掉末尾的换行符

			// 处理输入的学生信息
			studentInfo := strings.Fields(studentInput) // 用空格分割输入
			if len(studentInfo) != 5 {
				fmt.Println("输入格式错误，请按照正确格式输入（学号 姓名 语文 数学 英语）。")
				break
			}

			// 转换学号为整数
			studentID := studentInfo[0]
			var studentIDInt int
			_, err := fmt.Sscanf(studentID, "%d", &studentIDInt)
			if err != nil {
				fmt.Println("学号格式不正确，请输入整数学号。")
				break
			}

			// 将学生成绩存入 students 字典
			students[studentIDInt] = map[string]string{
				"name":    studentInfo[1],
				"chinese": studentInfo[2],
				"math":    studentInfo[3],
				"english": studentInfo[4],
			}
			fmt.Println("学生成绩添加成功!")

		case 2:
			fmt.Println("查看学生成绩")
			if len(students) == 0 {
				fmt.Println("暂无学生信息。")
			} else {
				for id, student := range students {
					fmt.Printf("学号: %d\n", id)
					for subject, score := range student {
						fmt.Printf("%s: %s\n", subject, score)
					}
					fmt.Println()
				}
			}

		case 3:
			fmt.Println("更新学生成绩")
			var studentIDInt int
			fmt.Print("请输入要更新成绩的学生学号：")
			fmt.Scan(&studentIDInt)

			if student, exists := students[studentIDInt]; exists {
				fmt.Printf("当前信息: 学生 %s 成绩: 语文: %s 数学: %s 英语: %s\n", student["name"], student["chinese"], student["math"], student["english"])
				fmt.Print("请输入更新后的成绩（语文 数学 英语）：")
				var chinese, math, english string
				fmt.Scan(&chinese, &math, &english)

				// 更新学生成绩
				students[studentIDInt]["chinese"] = chinese
				students[studentIDInt]["math"] = math
				students[studentIDInt]["english"] = english
				fmt.Println("学生成绩更新成功!")
			} else {
				fmt.Println("未找到该学生信息!")
			}

		case 4:
			fmt.Println("删除学生成绩")
			var studentIDInt int
			fmt.Print("请输入要删除的学生学号：")
			fmt.Scan(&studentIDInt)

			if _, exists := students[studentIDInt]; exists {
				delete(students, studentIDInt)
				fmt.Println("学生成绩删除成功!")
			} else {
				fmt.Println("未找到该学生信息!")
			}

		case 5:
			fmt.Println("退出成功")
			return

		default:
			fmt.Println("输入选项错误，请重新输入。")
		}
	}
}
