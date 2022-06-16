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
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	accessKey string
	secretKey string
	bucket    string
	space     string
}

type TencentUserInfo struct {
	AccessKey string
	SecretKey string
	Host      []string
}

func NewTencentUserInfo() *TencentUserInfo {
	//print("space name")
	//print(os.Getenv("qiniu_space"))
	host := os.Getenv("tencent_host")
	http_host := fmt.Sprintf("https://%s", host)
	https_host := fmt.Sprintf("https://www.%s", host)
	return &TencentUserInfo{
		AccessKey: os.Getenv("tencent_access_key"),
		SecretKey: os.Getenv("tencent_secret_key"),
		Host:      []string{http_host, https_host},
	}
}

const SOURCE = "webupload"

func NewUserInfo() *UserInfo {
	//print("space name")
	//print(os.Getenv("qiniu_space"))
	return &UserInfo{
		accessKey: os.Getenv("qiniu_access_key"),
		secretKey: os.Getenv("qiniu_secret_key"),
		bucket:    os.Getenv("qiniu_bucket"),
		space:     os.Getenv("qiniu_space"),
	}
}

type MysqlDB struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
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

var DB *sql.DB

func init() {
	//json读取数据
	//v := MysqlDB{}
	// 外部的读取的时候，当前目录是根目录，所以需要完整的路径
	conf, err := JsonParse("service/config.json")
	if err != nil {
		log.Fatalln("请创建server/config.json 配置文件")
	}
	DB = conf.InitDB()
}

func (this *MysqlDB) InitDB() *sql.DB {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local",
		this.Username, this.Password,
		this.Host, this.Port, this.Db)
	log.Println(conn)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("mysql连接成功")
	return db
}
