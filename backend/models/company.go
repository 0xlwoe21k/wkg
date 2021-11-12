package models


type Company struct {
	Id int  			`gorm:"primary_key;column:id" json:"id"`
	ProjectType string 	`gorm:"column:projectType"    json:"projectType"`
	CompanyName string  `gorm:"column:companyName"    json:"companyName"`
	Domain string 		`gorm:"column:domain"         json:"domain"`
	SrcUrl string 		`gorm:"column:srcUrl"         json:"srcUrl"`
	KeyWord string		`gorm:"column:keyWord"        json:"keyWord"`
	MonitorStatus  *bool `gorm:"type:boolean; column:monitorStatus"  json:"monitorStatus"`
	MonitorRate int  	`gorm:"column:monitorRate"    json:"monitorRate"`
	VulnScanStatus *bool `gorm:"type:boolean; column:vulnScanStatus" json:"vulnScanStatus"`
	VulnScanRate   int 	`gorm:"column:vulnScanRate"   json:"vulnScanRate"`
	LastUpdateTime string	`gorm:"column:lastUpdateTime" json:"lastUpdateTime"`
}



func (d *Company)TableName() string {
	return "company"
}



