package controllers

import (
	"backend/libs/gch"
	nuclei2 "backend/module/vulnscan/nuclei/v2/cmd/nuclei"
	"fmt"
	"github.com/astaxie/beego"
	"os"
)

type VulnController struct {
	beego.Controller
}



func (wsConn *wsConnection) vulnScanSingleProcLoop() {
	// 启动一个gouroutine发送心跳
	//resOutput := make(chan *requests.Output,10)
	//rootCtx ,cancel := context.WithCancel(context.Background())

	//收到消息后开始暴破域名信息
	go func() {
		for {
			select {
			case result := <-gch.GChan:
				if len(result) > 0{
					err := wsConn.wsWrite(1, []byte(result))
					if err != nil {
						fmt.Println("write fail")
						continue
					}
				}
			}
		}
	}()

	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			fmt.Println("read fail")
			break
		}
		fmt.Println(string([]byte(msg.data)))
		target := string(msg.data)

		nuclei2.NucleiScanSingle(target)
		fmt.Println("123")
	}
}

func (this *VulnController) VulnScanSingle()   {
	wsSocket, err := wsUpgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		return
	}
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}
	// 处理器
	go wsConn.vulnScanSingleProcLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}




func (wsConn *wsConnection) vulnScanMultiProcLoop() {
	// 启动一个gouroutine发送心跳
	//收到消息后开始暴破域名信息
	go func() {
		for {
			select {
			case result := <-gch.GChan:

				if len(result) > 0{
					err := wsConn.wsWrite(1, []byte(result))
					if err != nil {
						fmt.Println("write fail")
						continue
					}
				}
			}
		}
	}()

	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			fmt.Println("read fail")
			break
		}
		fmt.Println(string([]byte(msg.data)))
		if msg.data != nil{
			targetList := string(msg.data)
			fp ,err:= os.OpenFile("./upload/target.txt",os.O_RDWR|os.O_CREATE,644)
			if err != nil{
				fmt.Println("open file error.")
				continue
			}
			fp.WriteString(targetList)
			fp.Close()
		}
		nuclei2.NucleiScanMulti("./upload/target.txt")
		fmt.Println("123")
	}
}

func (this *VulnController) VulnScanMulti()   {
	wsSocket, err := wsUpgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		return
	}
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}
	// 处理器
	go wsConn.vulnScanMultiProcLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}