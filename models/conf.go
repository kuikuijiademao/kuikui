package models

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type MysqlConf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"post"`
}

func (c *MysqlConf) getConf() *MysqlConf {
	//读取resources/application.yaml文件
	yamlFile, err := ioutil.ReadFile("./configs/mysql.yaml")
	//若出现错误，打印错误提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func (c *MysqlConf) GetConfStr() (string, string, string) {
	//读取resources/application.yaml文件
	yamlFile, err := ioutil.ReadFile("./configs/mysql.yaml")
	//若出现错误，打印错误提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.UserName,
		c.Password,
		c.Url,
		c.Port,
		c.DbName,
	)
	cdsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.UserName,
		c.Password,
		c.Url,
		c.Port,
		c.DbName,
	)

	driverName := "mysql"
	return dsn, driverName, cdsn
}
