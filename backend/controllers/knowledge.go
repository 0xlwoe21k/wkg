package controllers

import (
	"backend/services/knowledgeService"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type KnowledgeController struct {
	beego.Controller
}

func (c *KnowledgeController) GetTopCategories() {
	rootCategories, err := knowledgeService.GetTopCategories()
	if err != nil {
		fmt.Println("knowledge.go err[", err, "]")
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	catedata, err := json.Marshal(rootCategories)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}

	type Res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}
	res, _ := json.Marshal(Res{Code: 200, Msg: "success", Data: string(catedata)})

	c.Ctx.WriteString(string(res))

}

type KnowStrut struct {
	Content     string `json:"content"`
	SecondValue string `json:"secondValue"`
	Title       string `json:"title"`
	TopValue    string `json:"topValue"`
	Key         string `json:"key"`
}

func (c *KnowledgeController) SaveNewKnowledge() {
	var ks = &KnowStrut{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, ks)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	if ks.Title == "" || ks.Content == ""{
		res, _ := json.Marshal(Result{Code: 400, Msg: "title or content is null."})
		c.Ctx.WriteString(string(res))
		return
	}

	key, err := knowledgeService.SaveNewKnowledge(ks.TopValue, ks.SecondValue, ks.Title, ks.Content)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	} else {
		type KK struct {
			Key string `json:"key"`
		}

		var kx []KK
		x := KK{Key: key}
		kx = append(kx, x)
		kd, _ := json.Marshal(kx)
		res, _ := json.Marshal(Res{Code: 200, Msg: "save success.", Data: string(kd)})
		c.Ctx.WriteString(string(res))
		return
	}
}

func (c *KnowledgeController) SaveEditKnowledge() {
	var ks = &KnowStrut{}
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, ks)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	if ks.Title == "" || ks.Content == ""{
		res, _ := json.Marshal(Result{Code: 400, Msg: "title or content is null."})
		c.Ctx.WriteString(string(res))
		return
	}

	err = knowledgeService.SaveEditKnowledge(ks.TopValue, ks.SecondValue, ks.Title, ks.Content, ks.Key)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	} else {
		res, _ := json.Marshal(Result{Code: 200, Msg: "save success."})
		c.Ctx.WriteString(string(res))
		return
	}
}

func (c *KnowledgeController) GetSecondCategories() {
	key := c.Ctx.Request.FormValue("key")

	rootCategories, err := knowledgeService.GetSecondCategories(key)
	if err != nil {
		fmt.Println("knowledge.go err[", err, "]")
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	catedata, err := json.Marshal(rootCategories)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}

	type Res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}
	res, _ := json.Marshal(Res{Code: 200, Msg: "success", Data: string(catedata)})

	c.Ctx.WriteString(string(res))

}

func (c *KnowledgeController) GetKnowledgeCategories() {
	key := c.Ctx.Request.FormValue("key")

	KnowledgeCategories, err := knowledgeService.GetKnowledgeCategoriesList(key)
	if err != nil {
		fmt.Println("knowledge.go err[", err, "]")
		res, _ := json.Marshal(Result{Code: 400, Msg: "GetKnowledgeCategoriesList error."})
		c.Ctx.WriteString(string(res))
		return
	}

	catedata, err := json.Marshal(KnowledgeCategories)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}

	type Res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}
	res, _ := json.Marshal(Res{Code: 200, Msg: "success", Data: string(catedata)})

	c.Ctx.WriteString(string(res))

}

type Res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (c *KnowledgeController) GetKnowledge() {
	key := c.Ctx.Request.FormValue("key")

	KnowledgeCategories, err := knowledgeService.GetKnowledge(key)
	if err != nil {
		fmt.Println("knowledge.go err[", err, "]")
		res, _ := json.Marshal(Result{Code: 400, Msg: "GetKnowledgeCategoriesList error."})
		c.Ctx.WriteString(string(res))
		return
	}

	catedata, err := json.Marshal(KnowledgeCategories)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}

	if len(catedata) > 0 {
		res, _ := json.Marshal(Res{Code: 200, Msg: "success", Data: string(catedata)})
		c.Ctx.WriteString(string(res))
	} else {
		res, _ := json.Marshal(Res{Code: 200, Msg: "success", Data: "[]"})
		c.Ctx.WriteString(string(res))
	}

}

func (c *KnowledgeController) GetTopSelectOption() {

	topOtions, err := knowledgeService.GetTopSelectOption()
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "GetTopSelectOption error."})
		c.Ctx.WriteString(string(res))
		return
	}
	data, _ := json.Marshal(topOtions)
	res, _ := json.Marshal(Result{Code: 200, Msg: string(data)})
	c.Ctx.WriteString(string(res))
	return

}

func (c *KnowledgeController) GetSecodSelectOption() {
	key := c.Ctx.Request.FormValue("key")

	secondOption, err := knowledgeService.GetSecodSelectOption(key)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "GetSecodSelectOption error."})
		c.Ctx.WriteString(string(res))
		return
	}

	knowData, err := json.Marshal(secondOption)
	if err != nil {
		res, _ := json.Marshal(Result{Code: 400, Msg: "json marshal error.."})
		c.Ctx.WriteString(string(res))
	}

	if len(secondOption) > 0 {
		res, _ := json.Marshal(Res{Code: 200, Msg: "success", Data: string(knowData)})
		c.Ctx.WriteString(string(res))
	} else {
		res, _ := json.Marshal(Res{Code: 200, Msg: "success", Data: "[]"})
		c.Ctx.WriteString(string(res))
	}
	return

}

func (c *KnowledgeController) GetSummary() {

	knowTitleList, err := knowledgeService.GetSummary()
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "GetSecodSelectOption error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data, _ := json.Marshal(knowTitleList)
	res, _ := json.Marshal(Result{Code: 200, Msg: string(data)})
	c.Ctx.WriteString(string(res))
	return

}

func (c *KnowledgeController) AddTopNode() {
	topNode := c.Ctx.Request.FormValue("topNode")

	err := knowledgeService.AddTopNode(topNode)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Res{Code: 400, Msg: "AddTopNode error.",Data: "[]"})
		c.Ctx.WriteString(string(res))
		return
	}

	res, _ := json.Marshal(Res{Code: 200, Msg: "sucess",Data: "[]"})
	c.Ctx.WriteString(string(res))
	return

}

func (c *KnowledgeController) AddSecondNode() {
	SecondNode := c.Ctx.Request.FormValue("topNode")
	parentKey := c.Ctx.Request.FormValue("parentKey")


	err := knowledgeService.AddSecondNode(parentKey,SecondNode,)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Res{Code: 400, Msg: "AddTopNode error.",Data: "[]"})
		c.Ctx.WriteString(string(res))
		return
	}

	res, _ := json.Marshal(Res{Code: 200, Msg: "sucess",Data: "[]"})
	c.Ctx.WriteString(string(res))
	return

}

func (c *KnowledgeController) GetTree() {

	tree, err := knowledgeService.GetCategoryTree()
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "GetSecodSelectOption error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data, _ := json.Marshal(tree)
	res, _ := json.Marshal(Result{Code: 200, Msg: "success", Data: string(data)})
	c.Ctx.WriteString(string(res))
	return

}



func (c *KnowledgeController) DelTreeNode() {
	type delOption struct {
		IsLeaf bool `json:"isLeaf"`
		Key string `json:"key"`
	}

	var op delOption

	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &op)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "json unmarshal error."})
		c.Ctx.WriteString(string(res))
		return
	}

	err = knowledgeService.DelTreeNode(op.IsLeaf,op.Key)
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Res{Code: 400, Msg: "GetSecodSelectOption error.",Data: "[]"})
		c.Ctx.WriteString(string(res))
		return
	}

	res, _ := json.Marshal(Result{Code: 200, Msg: "success", Data: "[]"})
	c.Ctx.WriteString(string(res))
	return

}


func (c *KnowledgeController) DelKnowledgeByKey() {

	tree, err := knowledgeService.GetCategoryTree()
	if err != nil {
		fmt.Println(err)
		res, _ := json.Marshal(Result{Code: 400, Msg: "GetSecodSelectOption error."})
		c.Ctx.WriteString(string(res))
		return
	}

	data, _ := json.Marshal(tree)
	res, _ := json.Marshal(Result{Code: 200, Msg: "success", Data: string(data)})
	c.Ctx.WriteString(string(res))
	return

}