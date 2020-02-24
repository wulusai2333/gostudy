package ini

import (
	"reflect"
	"testing"
)

/*
	go test -cover
	go test -cover -coverprofile=c.out
	go tool cover -html=c.out
*/
type Mysql struct {
	Ip       string `ini:"ip"`
	Port     int64  `ini:"port"`
	Name     string `ini:"name"`
	Password string `ini:"password"`
}

func TestLoadIni(t *testing.T) {
	var want = Mysql{
		Ip:       "1.1.1.1",
		Port:     3306,
		Name:     "root",
		Password: "123456",
	}
	var got = Mysql{}
	LoadIni(&got, "ini")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want:%#v ,got:%#v", want, got)
	}
}

func BenchmarkLoadIni(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var want = Mysql{
			Ip:       "1.1.1.1",
			Port:     3306,
			Name:     "root",
			Password: "123456",
		}
		var got = Mysql{}
		LoadIni(&got, "ini")
		if !reflect.DeepEqual(want, got) {
			b.Fatalf("want:%#v ,got:%#v", want, got)
		}
	}
}
