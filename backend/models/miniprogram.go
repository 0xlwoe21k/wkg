package models


type MiniProgram struct {
	Id         int     `gorm:"primary_key;column:id"`
	Cid        int     `gorm:"foreignkey;column:cid"`
	Name       string  `gorm:"column:name"`
	Notice     string  `gorm:"column:notice"`
	UpdateTime string  `gorm:"column:updateTime"`
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}

func (d *MiniProgram)TableName() string {
	return "miniprogram"
}





