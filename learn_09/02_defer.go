package main

import "fmt"

func f1() int {
	i := 5
	defer func() {
		i++
	}()
	return i
	/*
		rval = 5
		defer
		ret
	*/
}
func f2() *int {
	i := 5
	defer func() {
		i++
		fmt.Printf(":::%p\n", &i)
	}()
	fmt.Printf(":::%p\n", &i)
	return &i
}
func f3() (result int) {
	defer func() {
		result++
	}()
	return 5 // result = 5  defer_func  ret result
}
func f4() (result int) {
	defer func() {
		result++
	}()
	return result // result = result(0) defer_func ret result
}
func f5() (r int) {
	t := 5
	defer func() {
		t = t + 1
	}()
	return t // r=t=5 defer_func t=6 ret r
}
func f6() (r int) {
	defer func(r int) {
		r = r + 1
	}(r)
	return 5
}
func f7() (r int) {
	defer func(x int) {
		r = x + 1
	}(r)
	return 5
}
func main() {
	//println(f1())
	//println(*f2())
	//println(f3())
	//println(f4())
	//println(f5())
	//println(f6())
	println(f7())

	/*
		return 语句
		rval = xxx
		ret

		defer 语句发生在这俩条之间
		rval = xxx
		defer_func
		ret
	*/
}
