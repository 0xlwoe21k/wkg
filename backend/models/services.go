package models

type Services struct {
	Id         int     `gorm:"primary_key;column:id"`
	Cid        int     `gorm:"primary_key;column:cid"`
	Service    string  `gorm:"column:service"`
	Ipport     string  `gorm:"column:ipport"`
	Product    string  `gorm:"column:product"`
	UpdateTime string  `gorm:"column:updateTime"`
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}


func (d *Services)TableName() string {
	return "services"
}