package main

import "fmt"

func main() {
	/*	stu := make(map[int]map[string]string)
		stu[1001] = map[string]string{"name": "liu", "age": "18"}
		stu[1002] = map[string]string{"name": "zzz", "age": "19"}
		fmt.Println(stu)
		fmt.Println(stu[1002]["age"])*/

	/*	var data = make(map[string][]string)
		data["山东省"] = []string{"济南", "青岛", "菏泽"}
		data["河北"] = []string{"河北", "廊坊", "保定"}
		fmt.Println(data)
		fmt.Println(data["山东省"][1])*/

	var stu = make([]map[string]string, 3)
	stu[0] = map[string]string{"name": "liu", "age": "28"}
	stu[1] = map[string]string{"name": "zzz", "age": "29"}
	stu[2] = map[string]string{"name": "xxx", "age": "29"}
	stu = append(stu, map[string]string{"name": "yyy", "age": "28"})
	fmt.Println(stu)
}
