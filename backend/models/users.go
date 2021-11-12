package models

type Users struct {
	Id         int     `gorm:"primary_key;column:id"`
	Username   int     `gorm:"column:username"`
	Password   string  `gorm:"column:password"`

}


func (d *Users)TableName() string {
	return "users"
}