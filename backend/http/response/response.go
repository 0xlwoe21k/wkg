package response

type Rdata struct {
	Token string `json:"token"`
}
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Rdata  `json:"data"`
}
