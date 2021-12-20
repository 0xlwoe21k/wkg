package models


type Ips struct {
	Id         int     `gorm:"primary_key;column:id"  json:"id"`
	Cid        int     `gorm:"foreignkey;column:cid" json:"cid"`
	Ip         string  `gorm:"column:ip"   json:"ip"`
	Os         string  `gorm:"column:os"   json:"os"`
	Indomains  string  `gorm:"column:indomains"   json:"indomains"`
	Geo        string  `gorm:"column:geo"   json:"geo"`
	UpdateTime string  `gorm:"column:updateTime"   json:"updateTime"`
	IsNew		*bool	`gorm:"column:isNew" json:"isNew"    json:"isNew"`
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}


func (d *Ips)TableName() string {
	return "ips"
}





