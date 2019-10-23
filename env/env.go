package env

import (
	"fmt"
	"github.com/spf13/viper"
	apiv1 "k8s.io/api/core/v1"
	"log"
	"reflect"
)

type Config struct {
	V *viper.Viper
}

var Con *Config

func InitConfig (yamlName string,path string) *Config {
	Con := &Config{
		V:viper.New(),
	}
	//设置配置文件的名字
	Con.V.SetConfigName(yamlName)
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH

	Con.V.AddConfigPath(path)
	//设置配置文件类型
	Con.V.SetConfigType("yaml")
	if err := Con.V.ReadInConfig(); err != nil{
		log.Fatal(err.Error())
	}
	return Con
}

func GenerateEnv(filename string,path string)[]apiv1.EnvVar{
	config := InitConfig(filename,path)
	envList := make([]apiv1.EnvVar,0)
	type EnvType []interface{}
	envTypeList := make(EnvType,0)
	//fmt.Println(reflect.TypeOf(config.V.Get("env")).Kind())
	//fmt.Println(config.V.Get("env"))
	if reflect.TypeOf(config.V.Get("env")).Kind() == reflect.Slice {
		envTypeList = config.V.Get("env").([]interface{})
	}
	for _,m := range envTypeList {
		if reflect.TypeOf(m).Kind() == reflect.Map {
			m2 := m.(map[interface{}]interface{})
			key := m2["name"].(string)
			value := ""
			switch reflect.TypeOf(m2["value"]).Kind() {
			case reflect.Bool:
				value = fmt.Sprintf("%t",m2["value"])
			case reflect.Int , reflect.Int64,reflect.Int32:
				value = fmt.Sprintf("%d",m2["value"])
			case reflect.Float32,reflect.Float64:
				value = fmt.Sprintf("%f",m2["value"])
			case reflect.String:
				value = fmt.Sprintf("%s",m2["value"])
			default:
				value = fmt.Sprintf("%v",m2["value"])
			}

			tmpEnv := apiv1.EnvVar{
				Name:key,
				Value:value,
			}
			envList = append(envList,tmpEnv)
		}

	}
	return envList

}