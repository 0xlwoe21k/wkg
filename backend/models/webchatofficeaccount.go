package models

type WebChatOfficeAccount struct {
	Id         int     `gorm:"primary_key;column:id"`
	Cid        int     `gorm:"foreignkey;column:cid"`
	Name       string  `gorm:"column:name"`
	Notice     string  `gorm:"column:notice"`
	UpdateTime string  `gorm:"column:updateTime"`
	IsNew		*bool	`gorm:"column:isNew" json:"isNew"`
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}

func (d *WebChatOfficeAccount)TableName() string {
	return "webchatofficeaccount"
}



