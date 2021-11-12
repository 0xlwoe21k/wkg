package models



type SystemConfig struct {
	Id int 						`gorm:"primary_key;column(id)"`
	EmailNotifyEnable bool 	`gorm:"column:emailNotifyEnable"`
	EmailServerAddr string 	`gorm:"column:email_ServerAddr"`
	EmailUserName string 		`gorm:"column:emailUserName"`
	EmailPassword string 		`gorm:"column:emailPassword"`
	WeChatNotifyEnable bool 	`gorm:"column:weChatNotifyEnable"`
	WeChatKey string 			`gorm:"column:weChatKey"`
	DingtalkNotfyEnable bool 	`gorm:"column:dingtalkNotfyEnable"`
	DingtalkAccessToken string `gorm:"column:dingtalkAccessToken"`
}



func (SystemConfig)TableName() string {
	return "systemConfig"
}