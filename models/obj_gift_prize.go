package models

// ObjGiftPrize 奖品数据
type ObjGiftPrize struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	PrizeNum     int    `json:"-"`
	LeftNum		 int    `json:"-"`
	PrizeCodeA   int    `json:"-"`
	PrizeCodeB   int    `json:"-"`
	Img          string `json:"img"`
	Displayorder int    `json:"displayorder"`
	Gtype        int    `json:"gtype"`
	Gdata        string `json:"gdata"`
}
