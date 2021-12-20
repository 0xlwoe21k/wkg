package global

import "fmt"

var(
	HasNewCompanyFlag = false
)

func init()  {
	fmt.Println("[*] init global env.")
}