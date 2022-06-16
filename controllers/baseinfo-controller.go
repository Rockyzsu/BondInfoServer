package controllers

import (
	"bondinfoserver/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func parseArgs(strs string) []string {
	str_list := strings.Split(strs, "|")
	return str_list

}
func BondBaseInfo(ctx *gin.Context) {
	//可转债基本信息

	if ctx.Request.Method == "POST" {

		sign := ctx.PostForm("sign")
		fmt.Println("current is ", sign)

		args := ctx.PostForm("args")
		fmt.Println("args is ", args)
		args_list := parseArgs(args)
		fmt.Println(args_list)
		status, bondList := service.BaseInfoQuery("test")
		if !status {
			fmt.Println("Error")
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"res":  "",
			"data": bondList,
		})
	}

}
