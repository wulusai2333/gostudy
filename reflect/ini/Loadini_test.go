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

/*func TestReadConf(t *testing.T) {
	//want:=make(map[string]map[string][]byte, 10)
	want := map[string]map[string][]byte{
		"Mysql": map[string][]byte{
			"ip":       []byte("1.1.1.1"),
			"port":     []byte("3306"),
			"name":     []byte("root"),
			"password": []byte("123456"),
		},
	}
	configFile, err := os.OpenFile("./config.ini", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	//得到一个文件读写对象
	reader := bufio.NewReader(configFile)
	got := readConf(reader)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want:%#v ,got:%#v %#v", want, got, string([]byte{0x33, 0x34, 0x35, 0x36, 0x31, 0x2e, 0x31}))
	}
}*/

/*
want:
map[string]map[string][]uint8{
	"Mysql":map[string][]uint8{
		"ip":[]uint8{0x31, 0x2e, 0x31, 0x2e, 0x31, 0x2e, 0x31},
		"name":[]uint8{0x72, 0x6f, 0x6f, 0x74},
		"password":[]uint8{0x31, 0x32, 0x33, 0x34, 0x35, 0x36},
		"port":[]uint8{0x33, 0x33, 0x30, 0x36}}} ,
got:map[string]map[string][]uint8{
	"Mysql":map[string][]uint8{
		"ip":[]uint8{0x33, 0x34, 0x35, 0x36, 0x31, 0x2e, 0x31},
		"name":[]uint8{0x72, 0x6f, 0x6f, 0x74},
		"password":[]uint8{0x31, 0x32, 0x33, 0x34, 0x35, 0x36},
		"port":[]uint8{0x33, 0x33, 0x30, 0x36}}}

*/
