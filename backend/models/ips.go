package models


type Ips struct {
	Id         int     `gorm:"primary_key;column:id"`
	Cid        int     `gorm:"foreignkey;column:cid"`
	Ip         string  `gorm:"column:ip"`
	Os         string  `gorm:"column:os"`
	Indomains  string  `gorm:"column:indomains"`
	Geo        string  `gorm:"column:geo"`
	UpdateTime string  `gorm:"column:updateTime"`
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}


func (d *Ips)TableName() string {
	return "ips"
}





