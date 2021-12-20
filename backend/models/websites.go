package models

type Websites struct {
	Id         int     `gorm:"primary_key;column:id" json:"id"`
	Cid        int     `gorm:"foreignkey;column:cid"  json:"cid"`
	Domain	   string  `gorm:"column:domain"  json:"domain"`
	Website    string  `gorm:"column:website"  json:"website"`
	Ips		   string  `gorm:"column:ips"  json:"ips"`
	Favicon	   string  `gorm:"column:favicon"  json:"favicon"`
	FaviconUrl	   string  `gorm:"column:faviconUrl"  json:"faviconUrl"`
	Title      string  `gorm:"column:title"  json:"title"`
	Headers    string  `gorm:"column:headers"  json:"headers"`
	Cert    string  `gorm:"column:cert"  json:"cert"`
	Finger     string  `gorm:"column:finger"  json:"finger"`
	Screenshot string  `gorm:"column:screenshot"  json:"screenshot"`
	UpdateTime string  `gorm:"column:updateTime" json:"updateTime"`
	IsNew		*bool	`gorm:"column:isNew" json:"isNew"`
	CDN			*bool
	Companys   Company `gorm:"ForeignKey:Id;AssociationForeignKey:Cid"`
}


func (d *Websites)TableName() string {
	return "websites"
}



