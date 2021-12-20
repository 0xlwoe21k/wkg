package models

type Apps struct {
	Id         int     `gorm:"primary_key;column:id"`
	Cid        int     `gorm:"primary_key;column:cid"`
	Appname    string  `gorm:"column:appname"`
	Notice     string  `gorm:"column:notice"`
	UpdateTime string  `gorm:"column:updateTime"`
	IsNew		*bool	`gorm:"column:isNew" json:"isNew"`
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}


func (d *Apps)TableName() string {
	return "apps"
}