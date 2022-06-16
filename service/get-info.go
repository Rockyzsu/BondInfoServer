package service

import (
	"bondinfoserver/model"
	"log"
)

type BONDList []model.Bondinfo

func BaseInfoQuery(text string) (bool, BONDList) {
	rows, err := DB.Query("select `可转债代码`,`可转债名称`,`可转债价格`,`溢价率`,`剩余规模`,`正股现价`,`最新转股价`,`到期时间`,`成交额(万元)`,`评级`,`转股起始日` from tb_bond_jisilu")
	if err != nil {
		log.Println(err)
		return false, nil
	}

	var bondsList BONDList
	for rows.Next() {
		var info model.Bondinfo
		err = rows.Scan(&info.Code, &info.Name, &info.Price, &info.PremiumRT, &info.CurrIssAmt, &info.ZgPrice, &info.ConvertPrice, &info.Deadline, &info.Volume, &info.Rating, &info.ConvertStartDate)
		if err != nil {
			log.Println("读取失败")
		}

		bondsList = append(bondsList, info)
	}

	return true, bondsList
}
