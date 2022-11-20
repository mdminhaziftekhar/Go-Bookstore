package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

type conf struct {
	Encoding  string `yaml:"encoding"`
	Reconnect bool   `yaml:"reconnect"`
	TCP       string `yaml:"tcp"`
	Database  string `yaml:"database"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// fmt.Println("DEBUG ...")
	// fmt.Println("Database : ", c.Database)
	// fmt.Println("Username : ", c.Username)
	// fmt.Println("Password : ", c.Password)
	// fmt.Println("Reconnect : ", c.Reconnect)
	// fmt.Println("Encoding : ", c.Encoding)

	return c
}

func Connect() {

	var c conf

	c.getConf()

	// USERNAME:PASSWORD@tcp(127.0.0.1:3306)/database_name?charset=utf8&parseTime=True&loc=Local

	loginCred := c.Username + ":" + c.Password + "@" + c.TCP + "/" + c.Database + "?" + c.Encoding

	d, err := gorm.Open("mysql", loginCred)

	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
