package model

type Bondinfo struct {
	Code             string  `json:"code"`
	Name             string  `json:"name"`
	Price            float32 `json:"price"`
	PremiumRT        float32 `json:"premium_rt"`   // 溢价率
	RedeemIcon       string  `json:"redeem_icon"`  // 强赎标志
	CurrIssAmt       float32 `json:"curr_iss_amt"` // 剩余规模
	ZgPrice          float32 `json:"zg_price"`
	ConvertPrice     float32 `json:"convert_price"`
	Deadline         string  `json:"dead_line"`
	Volume           float32 `json:"volume"`
	Rating           string  `json:"rating"`
	ConvertStartDate string  `json:"convert_start_date"`
}
