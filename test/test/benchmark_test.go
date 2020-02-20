package tool

import (
	"fmt"
	"os"
	"testing"
	"time"
)

/*
	基准测试:	通过b.n次测试确定代码的性能,至少测试1s
*/
/*
	基准测试 go test -bench=Str2bytes
goos: windows
goarch: amd64
pkg: github.com/wulusai2333/gostudy/test/test
BenchmarkStr2bytes-4    1000000000               0.347 ns/op
PASS
ok      github.com/wulusai2333/gostudy/test/test        0.875s

	包含内存申请的测试 go test -bench=Str2bytes -benchmem
goos: windows
goarch: amd64
pkg: github.com/wulusai2333/gostudy/test/test
#这些参数的意思
#	使用的cpu核心数 		执行次数 				每次操作使用时间 		占用字节数 			申请内存次数
BenchmarkStr2bytes-4    1000000000               0.307 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/wulusai2333/gostudy/test/test        0.891s

*/
func BenchmarkStr2bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Str2bytes("1234567890")
	}
}

/*
	性能比较测试
	go test -bench=Fib1 就会匹配所有包含 Fib1 字段的函数并执行
	go test -bench=Fib. 就会匹配所有包含 Fib 字段的函数并执行
	go test -bench=. 执行所有bench测试
	go test -bench=. -benchtime=10s 设定性能测试的时间
*/
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}
func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}
func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}
func BenchmarkFib30(b *testing.B) {
	benchmarkFib(b, 30)
}
func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}
func BenchmarkFib45(b *testing.B) {
	benchmarkFib(b, 45)
}

/*
	重置时间 排除一些准备工作非代码效率的时间
*/
func BenchmarkBytes2str(b *testing.B) {
	data := []byte("123")
	time.Sleep(time.Second)
	b.ResetTimer() //重置时间 以上代码准备耗费的时间不算在基准测试中
	for i := 0; i < b.N; i++ {
		Bytes2str(data)
	}
}

/*
	并行测试 测试在不同goroutine个数的场景下耗费的时间
	go test -bench=Str2bytes2cpu -cpu=1 指定执行使用的cpu个数
*/
func benchmarkStr2bytes2(b *testing.B, cpuNum int) {
	b.SetParallelism(cpuNum)             //设置cpu数
	b.RunParallel(func(pb *testing.PB) { //调用内置函数
		for pb.Next() { //循环
			Str2bytes("123")
		}
	})
}
func BenchmarkStr2bytes2g1(b *testing.B) {
	benchmarkStr2bytes2(b, 1)
}
func BenchmarkStr2bytes2g2(b *testing.B) {
	benchmarkStr2bytes2(b, 2)
}
func BenchmarkStr2bytes2g3(b *testing.B) {
	benchmarkStr2bytes2(b, 3)
}
func BenchmarkStr2bytes2g4(b *testing.B) {
	benchmarkStr2bytes2(b, 4)
}

/*
	Setup TearDown 设置的恢复
*/
func TestMain(m *testing.M) {
	fmt.Println("测试之前的准备工作")
	//如果TestMain用了flag 则这里应该加上flag.Parse()
	retCode := m.Run()
	fmt.Println("测试之后的收尾工作")
	os.Exit(retCode)
}

/*
	示例函数 用于生成文档
*/
func ExampleFib() {
	fmt.Println(Fib(2))
	fmt.Println(Fib(3))
	//Output: Fib2 Fib3

}
