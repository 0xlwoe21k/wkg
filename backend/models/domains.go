package models

type Domain struct {
	Id         	int     `gorm:"primary_key;column:id" json:"id"`
	Cid        	int     `gorm:"foreignkey;column:cid" json:"cid"`
	Domain     	string  `gorm:"column:domain" json:"domain"`
	Type       	string  `gorm:"column:type" json:"type"`
	Source 		string	`gorm:"column:source" json:"source"`
	UpdateTime 	string  `gorm:"column:updateTime"json:"updateTime"`
	IsNew		*bool	`gorm:"column:isNew" json:"isNew"`
	Companys   	Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}


func (d *Domain)TableName() string {
	return "domains"
}