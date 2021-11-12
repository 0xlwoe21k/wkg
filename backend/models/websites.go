package models



type Websites struct {
	Id         int     `gorm:"primary_key;column:id"`
	Cid        int     `gorm:"foreignkey;column:cid"`
	Website    string  `gorm:"column:website"`
	Title      string  `gorm:"column:title"`
	Headers    string  `gorm:"column:headers"`
	Finger     string  `gorm:"column:finger"`
	Screenshot string  `gorm:"column:screenshot"`
	UpdateTime string  `gorm:"column:updateTime"`
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}


func (d *Websites)TableName() string {
	return "websites"
}



