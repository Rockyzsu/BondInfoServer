package model

type Bondinfo struct {
	Code             string `json:"code"`
	Name             string `json:"name"`
	Price            string `json:"price"`
	PremiumRT        string `json:"premium_rt"`   // 溢价率
	RedeemIcon       string `json:"redeem_icon"`  // 强赎标志
	CurrIssAmt       string `json:"curr_iss_amt"` // 剩余规模
	ZgPrice          string `json:"zg_price"`
	ConvertPrice     string `json:"convert_price"`
	Deadline         string `json:"dead_line"`
	Volume           string `json:"volume"`
	Rating           string `json:"rating"`
	ConvertStartDate string `json:"convert_start_date"`
}
