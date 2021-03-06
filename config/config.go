package config

import (
	"encoding/json"
	"os"
)

/**
 * 服务端配置
 */
type AppConfig struct {
	AppName   string   `json:"app_name"`
	Host      string   `json:"host"`
	Port      string   `json:"port"`
	Mode      string   `json:"mode"`
	Static    string   `json:"static"`
	FilePath  string   `json:"file_path"`
	FontPath  string   `json:"font_path"`
	ImgPath   string   `json:"img_path"`
	JwtSecret string   `json:"jwt_secret"`
	DataBase  DataBase `json:"data_base"`
	Email     Email    `json:"email"`
}

/**
 * mysql配置
 */
type DataBase struct {
	Drive    string `json:"drive"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

/*
* email配置
 */
type Email struct {
	ServerHost string `json:"serverHost"`
	ServerPort int    `json:"serverPort"`
	FromEmail  string `json:"fromEmail"`
	FromPasswd string `json:"fromPasswd"`
}

func InitConfig() (config *AppConfig) {
	file, err := os.Open("./config.json") //当前目录指的是当前项目的跟目录
	// file, err := os.Open("/Users/yangxianglong/go/Vue_Iris/back-end/config.json")
	if err != nil {
		panic((err.Error()))
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	return
}
