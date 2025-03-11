package ch02

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*func TestAdd(t *testing.T) {
	ret := add(1, 2)
	fmt.Println("test1")
	if ret != 3 {
		t.Errorf("expect %d, actual %d", 3, ret)
	}
}

func TestAdd2(t *testing.T) {
	// go test -short 会跳过
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	ret := add(1, 5)
	fmt.Println("test2")
	if ret != 6 {
		t.Errorf("expect %d, actual %d", 6, ret)
	}
}*/

func TestAdd(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3},
		{12, 12, 24},
		{-9, 8, -1},
		{0, 0, 0},
	}
	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect %d, actual %d", value.out, re)
		}
	}
}

func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 123
	b = 456
	c = 579
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("%d + %d expect %d, actual %d", a, b, c, actual)
		}
	}
}

const numbers = 100

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str += fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
