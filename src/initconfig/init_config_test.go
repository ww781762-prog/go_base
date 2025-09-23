package initconfig

import (
	"fmt"
	"os"
	"testing"
)

type Config struct {
	ServerConfig `ini:"server"`
	MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Password string  `ini:"password"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

func TestMarshal(t *testing.T) {
	config := Config{}

	fData, err := os.ReadFile("/Users/wallace/GolandProjects/go_base/src/initconfig/config.ini")
	if err != nil {
		fmt.Println("读取文件失败,", err)
	}

	err = UnMarshal(fData, &config)
	if err != nil {
		fmt.Println("解析配置出错", err)
		return
	}
	fmt.Printf("%#v\n", config)

	buf, err := Marshal(config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}

func TestMarshalFile(t *testing.T) {
	config := Config{}
	fileName := "/Users/wallace/GolandProjects/go_base/src/initconfig/config.ini"
	err := UnMarshalFromFile(fileName, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", config)
	fileName2 := "/Users/wallace/GolandProjects/go_base/src/initconfig/config.ini_2"
	err = MarshalToFile(fileName2, config)
	if err != nil {
		fmt.Println(err)
	}
}
