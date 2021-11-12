package controllers


type PageParam struct {
	Page int `json:"page"`
	PageSize int `json:"pagesize"`
}

type DomaSearchStrut struct {
	Type string `json:"type"`
	KeyWord string `json:"keyWord"`
}