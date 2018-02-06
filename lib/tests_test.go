package lib

import (
	"testing"
	"fmt"
)

//斐波那契数列
//求出第n个数的值
func Fibonacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// 测试程序
func TestFibonacci(t *testing.T) {
	r := Fibonacci(10)
	if r != 55 {
		t.Errorf("Fibonacci(10) failed. Got %d, expected 55.", r)
	}
}

// 性能测试
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkAddStringWithSprintf(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%s", hello, world)
	}
}