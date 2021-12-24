package knowledgeService

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetCategoryTree(t *testing.T) {
	x,_:=GetCategoryTree()

	b,_:=json.Marshal(x)
	fmt.Println(string(b))
}