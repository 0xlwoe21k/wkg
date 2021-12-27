package models

type Knowledge struct {
	Id         int    `gorm:"primary_key;column:id"  json:"id"`
	Title      string `gorm:"column:title"    json:"title"`
	Content    string `gorm:"column:content" json:"content"`
	CKey       string `gorm:"column:ckey" json:"key"`
	UpdateTime string `gorm:"column:updateTime" json:"updateTime"`
}

func (d *Knowledge) TableName() string {
	return "knowledge"
}

type Category struct {
	Id       int    `gorm:"primary_key;column:id" json:"id"`
	ParentId int    `gorm:"column:parentId" json:"parentId"`
	Title    string `gorm:"column:title" json:"title"`
	Level    int    `gorm:"column:level" json:"level"`
	IsLeaf   bool   `gorm:"column:isLeaf" json:"isLeaf"`
	CKey     string `gorm:"column:ckey" json:"key"`
}

type CategoryTree struct {
	Id       int            `json:"-"`
	ParentId int            `json:"-"`
	Level    int			`json:"level"`
	Title    string         `json:"title,omitempty"`
	IsLeaf   bool           `json:"isLeaf,omitempty1"`
	Key      string         `json:"key,omitempty"`
	Children []CategoryTree `json:"children,omitempty"`
}

func (d *Category) TableName() string {
	return "category"
}
