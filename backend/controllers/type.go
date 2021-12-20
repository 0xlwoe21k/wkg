package controllers


type PageParam struct {
	Page int 		`json:"page"`
	PageSize int 	`json:"pagesize"`
	Type 	string	`json:"type"`
	Keyword	string	`json:"keyword"`
}

type SearchStrut struct {
	Type string `json:"type"`
	KeyWord string `json:"keyWord"`
}