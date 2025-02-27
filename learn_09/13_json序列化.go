package main

type Address struct {
	Province, City string
}

type Stu struct {
	// 这边首字母不能用小写 序列化失败 可见性的问题
	// 然后在序列化数据的时候，可以使用标签  `json:"name"`  这样在序列化数据的时候就是小写 name
	Name    string `json:"name"`
	Age     int
	Address Address
}

func main() {
	/*	var map01 = map[int]string{1: "111", 2: "222", 3: "333"}
		var stu01 = Stu{Name: "yy", Age: 18, Address: Address{Province: "sc", City: "cd"}}
	*/

	// 序列化
	/*	map01Json, _ := json.Marshal(map01)
		fmt.Println(string(map01Json))

		stu01Json, _ := json.Marshal(stu01)
		fmt.Println(stu01)
		fmt.Println(string(stu01Json))*/

	// 写入文件
	/*	err := os.WriteFile("data.json", stu01Json, 0666)
		if err != nil {
			fmt.Println(err)
		}*/

	// 反序列化
	/*	bytes, _ := os.ReadFile("data.json")
		fmt.Println(string(bytes))
		// 将json字符串转换成go的数据结构
		var stu Stu
		err := json.Unmarshal(bytes, &stu)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(stu)*/

	/*	var s = `{"name":"oi"}`
		var m map[string]string
		json.Unmarshal([]byte(s), &m)
		fmt.Println(m)*/
}
