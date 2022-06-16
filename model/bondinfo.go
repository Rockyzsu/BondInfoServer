package model

type History struct {
	Id        int    `json:"Id"`
	Url       string `json:"Url"`
	Updated   string `json:"Updated"`
	IsDeleted bool   `json:"isDeleted"`
}

type ContentText struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	Updated string `json:"updated"`
}

type Bondinfo struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Price      string `json:"price"`
	PremiumRT  string `json:"premium_rt"`   // 溢价率
	RedeemIcon string `json:"redeem_icon"`  // 强赎标志
	CurrIssAmt string `json:"curr_iss_amt"` // 剩余规模

}
