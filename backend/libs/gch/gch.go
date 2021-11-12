package gch


var (
	 GChan chan string
)

func init()  {
	GChan = make(chan string,10)
}