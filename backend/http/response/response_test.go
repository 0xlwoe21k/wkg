package response

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestTree(T *testing.T)  {

	dsn := "root:root@(localhost)/wkg?charset=utf8&parseTime=true&loc=Local&charset=utf8"
	Orm,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println("[!] db.go line:18 error:"+err.Error())
		os.Exit(0)
	}
	var tp []TreeProps

	err = Orm.Model(&TreeProps{}).Find(&tp).Error
	if err !=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(tp)

	//td := TreeProps{Id:0,Title: "cicso",Key: "1"}
	//header.Children = append(header.Children,td)
}

//第一层,知识库的

func buildTree(treeList []TreeProps) []TreeNode {

	//for _,v:=range treeList {
	//	if v.ParentId
	//}

return nil
}