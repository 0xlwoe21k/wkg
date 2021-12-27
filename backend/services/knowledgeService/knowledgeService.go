package knowledgeService

import (
	"backend/db"
	"backend/libs/helper"
	"backend/models"
	"fmt"
	"time"
)

type KnowledgeService struct {
}

func GetTopCategories() ([]models.Category, error) {
	categories := []models.Category{}
	err := db.Orm.Model(&models.Category{}).Select("title,ckey,isLeaf,level").Where("parentId=? and level=?", 0, 1).Find(&categories).Error
	if err != nil {

		return nil, err
	}

	return categories, nil
}

func GetSecondCategories(key string) ([]models.Category, error) {

	var id int
	err := db.Orm.Model(&models.Category{}).Select("id").Where("ckey=?", key).Find(&id).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return nil, err
	}

	categories := []models.Category{}
	err = db.Orm.Model(&models.Category{}).Where("parentId=? and level=?", id, 2).Find(&categories).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return nil, err
	}
	return categories, nil
}

func GetKnowledgeCategoriesList(key string) ([]models.Knowledge, error) {
	knowledge := []models.Knowledge{}
	var pid int
	err := db.Orm.Model(&models.Category{}).Select("id").Where("ckey=?", key).Find(&pid).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return nil, err
	}
	err = db.Orm.Model(&models.Knowledge{}).Select("title,ckey,isLeaf,level").Where("parentId=? ", pid).Find(&knowledge).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return nil, err
	}
	return knowledge, nil
}

func GetCategoryTree() ([]models.CategoryTree, error) {
	var tree []models.CategoryTree
	categories := []models.Category{}

	var ToptreeCache = make(map[int]models.CategoryTree)
	//一级菜单
	err := db.Orm.Debug().Model(&models.Category{}).Select("id,title,ckey,isLeaf,level").Where("parentId=? and level=?", 0, 1).Order("id ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var _tree models.CategoryTree
	for _, v := range categories {
		_tree.Id = v.Id
		_tree.ParentId = v.ParentId
		_tree.Key = v.CKey
		_tree.Level = v.Level
		_tree.Title = v.Title
		_tree.IsLeaf = v.IsLeaf
		ToptreeCache[v.Id] = _tree
	}
	_tree = models.CategoryTree{}

	var SecondtreeCache = make(map[int]models.CategoryTree)
	//二级菜单
	err = db.Orm.Model(&models.Category{}).Select("id,parentId,title,ckey,isLeaf,level").Where("level=?", 2).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	for _, v := range categories {
		_tree.Id = v.Id
		_tree.ParentId = v.ParentId
		_tree.Key = v.CKey
		_tree.Title = v.Title
		_tree.Level = v.Level
		_tree.IsLeaf = v.IsLeaf
		SecondtreeCache[v.Id] = _tree
	}
	_tree = models.CategoryTree{}
	//有的文章在一级菜单下，所以也要带上
	//Secondknowledge := []models.Knowledge{}
	//err = db.Orm.Model(&models.Knowledge{}).Select("parentId,title,ckey,isLeaf,level").Find(&Secondknowledge).Error
	//if err != nil {
	//	fmt.Println("knowledgeserver.go err[", err, "]")
	//	return nil, err
	//}
	//for _, v := range Secondknowledge {
	//	if ptree, ok := ToptreeCache[v.ParentId]; ok {
	//		if v.Level == 2 && v.IsLeaf == true {
	//			_tree.Id = v.Id
	//			_tree.ParentId = v.ParentId
	//			_tree.Key = v.CKey
	//			_tree.Title = v.Title
	//			_tree.Level = v.Level
	//			_tree.IsLeaf = v.IsLeaf
	//			ptree.Children = append(ptree.Children, _tree)
	//			ToptreeCache[v.ParentId] = ptree
	//		}
	//	}
	//}
	//_tree = models.CategoryTree{}
	knowledgeCate := []models.Category{}
	//三级菜单
	err = db.Orm.Model(&models.Category{}).Select("parentId,title,ckey,isLeaf,level").Where("level=?",3).Find(&knowledgeCate).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return nil, err
	}
	for _, v := range knowledgeCate {
		if ptree, ok := SecondtreeCache[v.ParentId]; ok {
			_tree.Id = v.Id
			_tree.ParentId = v.ParentId
			_tree.Key = v.CKey
			_tree.Level = v.Level
			_tree.Title = v.Title
			_tree.IsLeaf = v.IsLeaf
			ptree.Children = append(ptree.Children, _tree)
			SecondtreeCache[v.ParentId] = ptree
		}
	}
	_tree = models.CategoryTree{}
	//把二级菜单放到数组里
	for _, v := range SecondtreeCache {
		if ptree, ok := ToptreeCache[v.ParentId]; ok {
			ptree.Children = append(ptree.Children, v)
			ToptreeCache[v.ParentId] = ptree
		}
	}

	for _, v := range ToptreeCache {
		tree = append(tree, v)
	}

	return tree, nil
}

func GetKnowledge(key string) (*models.Knowledge, error) {
	knowledge := models.Knowledge{}

	err := db.Orm.Debug().Model(&models.Knowledge{}).Select("title,content,updateTime").Where("ckey=? ", key).Find(&knowledge).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return nil, err
	}

	return &knowledge, nil
}

type Option struct {
	Value string `json:"value"`
	Lable string `json:"label"`
}

func GetTopSelectOption() ([]Option, error) {

	//var pid int
	//err :=db.Orm.Model(&models.Category{}).Select("id").Where("ckey=?",key).Find(&pid).Error
	//if err != nil{
	//	fmt.Println("knowledgeserver.go err[",err,"]")
	//	return nil,err
	//}
	cate := []models.Category{}
	err := db.Orm.Model(&models.Category{}).Where("level=?", 1).Find(&cate).Error
	if err != nil {
		return nil, err
	}
	var option []Option
	for _, v := range cate {
		option = append(option, Option{Value: v.CKey, Lable: v.Title})
	}
	return option, nil
}

func GetSecodSelectOption(key string) ([]Option, error) {
	var pid int
	err := db.Orm.Model(&models.Category{}).Select("id").Where("ckey=?", key).Find(&pid).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return nil, err
	}

	cate := []models.Category{}
	err = db.Orm.Debug().Model(&models.Category{}).Where("parentId=? and level=?", pid, 2).Find(&cate).Error
	if err != nil {
		return nil, err
	}
	var option []Option
	for _, v := range cate {
		option = append(option, Option{Value: v.CKey, Lable: v.Title})
	}

	return option, nil
}

type KList struct {
	Title      string `json:"title"`
	UpdateTime string `json:"updateTime"`
}

func GetSummary() ([]KList, error) {
	var knowledgeList []KList
	var know []models.Knowledge

	err := db.Orm.Model(&models.Knowledge{}).Select("title,updateTime").Order("updateTime desc").Limit(10).Find(&know).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return nil, err
	}

	for _, v := range know {
		var _klist KList
		_klist.Title = v.Title
		_klist.UpdateTime = v.UpdateTime
		knowledgeList = append(knowledgeList, _klist)
	}

	return knowledgeList, nil
}

func AddTopNode(topNode string) (error) {
	var cate models.Category

	cate.ParentId = 0
	cate.Title = topNode
	cate.Level = 1
	cate.CKey = helper.Md5(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))

	err := db.Orm.Model(&models.Category{}).Create(&cate).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return err
	}
	return nil
}

func AddSecondNode(topKey string,secondNode string) (error) {
	var cate models.Category

	var pid int
	err := db.Orm.Model(&models.Category{}).Select("id").Where("ckey=?", topKey).Find(&pid).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return err
	}
	cate.ParentId = pid
	cate.Title = secondNode
	cate.Level = 2
	cate.CKey = helper.Md5(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))

	err = db.Orm.Model(&models.Category{}).Create(&cate).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return err
	}
	return nil
}




func SaveNewKnowledge(topValue string, secondValue string, title string, content string)(string,error)  {
	var pid int
	var err error
	var tKnow models.Knowledge
	var cate models.Category
	if secondValue != "" {
		err := db.Orm.Model(&models.Category{}).Select("id").Where("ckey=?", secondValue).Find(&pid).Error
		if err != nil {
			fmt.Println("knowledgeserver.go err[", err, "]")
			return "",err
		}
		cate.Level = 3
	} else {
		err := db.Orm.Model(&models.Category{}).Select("id").Where("ckey=?", topValue).Find(&pid).Error
		if err != nil {
			fmt.Println("knowledgeserver.go err[", err, "]")
			return "",err
		}
		cate.Level = 2
	}
	now := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	tKnow.Content = content
	tKnow.Title = title
	tKnow.CKey = helper.Md5(now)
	tKnow.UpdateTime = now


	cate.ParentId = pid
	cate.Title = title
	cate.CKey = tKnow.CKey
	cate.IsLeaf = true

	err = db.Orm.Debug().Model(&models.Knowledge{}).Create(&tKnow).Error
	if err != nil {
		return "",err
	}


	err = db.Orm.Debug().Model(&models.Category{}).Create(&cate).Error
	if err != nil {
		return "",err
	}

	return tKnow.CKey, nil
}

func SaveEditKnowledge(topValue string, secondValue string, title string, content string, key string) error {
	var pid int
	err := db.Orm.Model(&models.Category{}).Select("id").Where("ckey=?", secondValue).Find(&pid).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return err
	}
	now := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	var tKnow models.Knowledge
	tKnow.Content = content
	tKnow.Title = title
	tKnow.UpdateTime = now
	err = db.Orm.Debug().Where("ckey=?", key).Updates(&tKnow).Error
	if err != nil {
		return err
	}
	return nil
}
func DelTreeNode(isLeaf bool, key string) error {
	err := db.Orm.Model(&models.Category{}).Where("ckey=?",key).Delete(&models.Category{}).Error
	if err != nil {
		fmt.Println("knowledgeserver.go err[", err, "]")
		return err
	}

	if(isLeaf){
		err = db.Orm.Model(&models.Knowledge{}).Where("ckey=?",key).Delete(&models.Knowledge{}).Error
		if err != nil {
			fmt.Println("knowledgeServer.go err[", err, "]")
			return err
		}
	}

	return nil
}


