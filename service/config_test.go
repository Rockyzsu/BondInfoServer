package service

import (
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	//v:=MysqlDB{}
	//fmt.Println(v)
	fmt.Println(os.Getwd())
	conf, err := JsonParse("config.json")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("Mine")
	fmt.Println(conf)
}
