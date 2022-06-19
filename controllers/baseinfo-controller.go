package controllers

import (
	"bondinfoserver/cache"
	"bondinfoserver/model"
	"bondinfoserver/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseStruct struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data []model.Bondinfo `json:"data"`
}

func parseArgs(strs string) []string {
	str_list := strings.Split(strs, "|")
	return str_list

}

var cached *cache.Cache

func init() {
	cached = new(cache.Cache)
	cached.CacheInit()
}

func BondBaseInfo(ctx *gin.Context) {
	//可转债基本信息

	if ctx.Request.Method == "POST" {

		sign := ctx.PostForm("sign") //鉴权使用
		if !cached.CheckUserExist(sign) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  "sign error",
				"data": []model.Bondinfo{},
			})
			return
		}
		var result string
		currentTime := time.Now()
		currentTimeStr := currentTime.Format("15:04:05")
		if currentTimeStr >= "16:00:00" {
			//获取当天的数据 key
			oldTime := currentTime.Format("20060102")
			_result, found := cached.Get(oldTime)
			if !found {
				//获取json
				status, bondList := service.BaseInfoQuery("test")
				if !status {
					fmt.Println("Error")
					ctx.JSON(http.StatusBadRequest, gin.H{
						"code": -1,
						"msg":  "数据库出错",
						"data": []model.Bondinfo{},
					})
				}
				data_str, _ := json.Marshal(&bondList)

				cached.Set(oldTime, string(data_str))
				ctx.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "",
					"data": bondList,
				})
				return
			}
			fmt.Println("from cached")
			result = _result

		} else {
			//获取前一天的数据 key
			newTime := currentTime.AddDate(0, 0, -1).Format("20060102") //若要获取3天前的时间，则应将-2改为-3
			_result, found := cached.Get(newTime)
			if !found && _result == "" {
				//获取json
				status, bondList := service.BaseInfoQuery("test")
				if !status {
					fmt.Println("Error")
				}
				data_str, _ := json.Marshal(&bondList)

				cached.Set(newTime, string(data_str))
				ctx.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "",
					"data": bondList,
				})
				return
			}
		}

		fmt.Println("from cached")
		var resp_body []model.Bondinfo

		json.Unmarshal([]byte(result), &resp_body)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "",
			"data": resp_body,
		})
		//fmt.Println("current is ", sign)
		//
		//args := ctx.PostForm("args")
		//fmt.Println("args is ", args)
		//args_list := parseArgs(args)
		//fmt.Println(args_list)

	}

}
