package split

import (
	"reflect"
	"testing"
)

/*
	单元测试:
	好处 测试代码直接上传到代码仓库中,避免了不同人都要写自己的测试和测试用例丢失带来的麻烦
	一个好的测试标准:
	测试函数覆盖率 100%
	代码覆盖率 60%
	如果达不到,要么是测试用例写的不够周全,要么代码问题太多
*/
func TestSplit(t *testing.T) {
	got := Split("abc", "b")
	want := []string{"a", "c"}
	if !reflect.DeepEqual(got, want) {
		//测试失败了
		t.Errorf("want: %#v ,got: %#v", want, got)
	}
}

//测试组
func TestSplitGroup(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := []test{
		{input: "abc", sep: "b", want: []string{"a", "c"}},
		{input: "asdfg", sep: "d", want: []string{"as", "fg"}},
		{input: "qqwweerrsa", sep: "w", want: []string{"qq", "", "eerrsa"}},
	}
	for _, tr := range tests {
		got := Split(tr.input, tr.sep)
		if !reflect.DeepEqual(got, tr.want) {
			t.Fatalf("want:%#v ,got:%#v", tr.want, got)
		}
	}
}

/*
	子测试
	在控制台可以
		命令 go test -run=TestSplitSon/simple 来启动指定测试用例
		命令 go test -cover 测试用例的覆盖率,通常要求覆盖大部分代码,最好能覆盖所有代码
		命令 go test -cover -coverprofile=c.out 将测试覆盖率相关信息打印到文件中
		命令 go tool cover -html=c.out 用html的方式打开刚才生成的文件
*/

func TestSplitSon(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple": {input: "abc", sep: "b", want: []string{"a", "c"}},
		"wrong":  {input: "asdfg", sep: "d", want: []string{"as", "fg"}},
		"more":   {input: "qqwweerrsa", sep: "w", want: []string{"qq", "", "eerrsa"}},
	}
	for name, tr := range tests {
		//子测试,即在所有测试用例中只执行指定名称测试
		t.Run(name, func(t *testing.T) {
			got := Split(tr.input, tr.sep)
			if !reflect.DeepEqual(got, tr.want) {
				t.Fatalf("want:%#v ,got:%#v", tr.want, got)
			}
		})

	}
}
