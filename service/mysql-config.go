/*
 * #Date: 2021-11-23 02:07:18
 * #Author: Rocky Chen
 * #LastEditors: Rocky Chen
 * #LastEditTime: 2021-11-23 02:11:25
 * #Description:
 */
package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

type MysqlDB struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
	HttpPort int
	Key      string
}

func JsonParse(filename string) (MysqlDB, error) {
	v := MysqlDB{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return v, err
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Println(err)
		return v, err
	}
	return v, nil
}

//全局
var DB *sql.DB
var Port int

func init() {
	//json读取数据
	//v := MysqlDB{}
	// 外部的读取的时候，当前目录是根目录，所以需要完整的路径
	conf, err := JsonParse("service/mysql.json")
	if err != nil {
		log.Fatalln("请创建server/mysql.json 配置文件")
	}
	DB = conf.InitDB()
	Port = conf.HttpPort
}

func (this *MysqlDB) InitDB() *sql.DB {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local",
		this.Username, this.Password,
		this.Host, this.Port, this.Db)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("mysql连接成功")
	return db
}
