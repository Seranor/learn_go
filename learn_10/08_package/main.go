package main // 声明包名：main包
import (
	"08_package/utils"
	"fmt"
	"github.com/jinzhu/now"
)

import f "fmt" // 别名

//import . "fmt"  // 省略引用格式
//import _ "fmt"  // 匿名导入

func main() {
	ret1 := utils.Add(2, 3)
	fmt.Println(ret1)
	fmt.Println(now.BeginningOfMinute())
	f.Println("hello world")

}
