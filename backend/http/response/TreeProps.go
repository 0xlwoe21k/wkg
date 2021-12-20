package response



type TreeProps struct {
	Id		int64	`json:"id,omitempty" gorm:"id"`
	ParentId	int64	`json:"parentIid,omitempty" gorm:"parentIid"`
	Title 	string	`json:"title,omitempty" gorm:"title"`
	Key   	string	`json:"key,omitempty" gorm:"key"`
	IsLeaf	string 	`json:"isLeaf,omitempty" gorm:"isLeaf"`
}


type TreeNode struct {
	Id		int64	`json:"id,omitempty" gorm:"id"`
	ParentId	int64	`json:"parentIid,omitempty" gorm:"parentIid"`
	Title 	string	`json:"title,omitempty" gorm:"title"`
	Key   	string	`json:"key,omitempty" gorm:"key"`
	IsLeaf	string 	`json:"isLeaf,omitempty" gorm:"isLeaf"`
	Children []TreeProps  `json:"children ,omitempty" gorm:"children"`
}



func (d *TreeProps)TableName() string {
	return "knowledgeCategories"
}

