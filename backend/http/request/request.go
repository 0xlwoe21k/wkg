package request

type Query struct {
	Page int		`json:"page"`
	PageSize int `json:"pagesize"`
	Type string		`json:"type"`
	Keyword string	`json:"keyword"`
	Cid int		`json:"cid"`
}