package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"reflect"
 )
 
type Config struct {
	//变量名和顺序 必须与 config.toml文件一致
	// web配置
	AppName				string
	Port 				int
	Env					string
	TimeoutThreshold	int     // 超时告警时间，单位:ms

	// 日志配置
	LogFile string

	// 告警配置
	BotWarningEnv		string
	RobotURL       		string
	MentionWhoList 		[]string

	// database配置
	Uri					string
	Database 			string
	Table 				string
}

var config = InitConfig()

func InitConfig() Config{
	f := "./config/config.toml"
	if _, err := os.Stat(f); err != nil {
	   panic(err)
	}
	var conf Config
	if _, err := toml.DecodeFile(f, &conf); err != nil {
	   panic(err)
	}
 
	valueOf := reflect.ValueOf(conf)
	if valueOf.Kind() == reflect.Ptr {
        valueOf = valueOf.Elem()
    }
    typeOf := valueOf.Type()
	fmt.Println("*************** Config Begin ***************")
    if valueOf.Kind() == reflect.Struct {
        for i := 0; i < valueOf.NumField(); i++ {
            // f := v.Field(i)
			fieldName := typeOf.Field(i).Name
        	fieldValue := valueOf.Field(i).Interface()
            fmt.Printf("[%-20s]: %v\n", fieldName, fieldValue)
        }
    }
	fmt.Println("*************** Config End ***************")
	return conf
}


func GetConfig() *Config{
	return &config
}