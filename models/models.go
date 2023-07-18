package models

type Transactions struct {
	//gorm.Model

	Tx_id        string `gorm:"primary_key;type:varchar(40)" json:"Tx_id"`
	Acc_id       string `gorm:"type:varchar(40)" json:"Acc_id"`
	Tx_ts        string `gorm:"column:tx_ts" json:"Tx_ts"`
	Status       string `gorm:"type:varchar(40)" json:"Status"`
	Amount       uint64 `json:"Amount"`
	Merchantname string `json:"Merchantname" gorm:"text`
	Merchant_id  string `gorm:"type:varchar(40)" json:"Merchant_id"`
	Tx_type      string `json:"Tx_type" gorm:"text`
}
