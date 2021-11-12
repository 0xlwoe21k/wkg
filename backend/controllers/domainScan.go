package controllers

import (
	subdomainscan2 "backend/module/subdomainscan"
	requests2 "backend/module/subdomainscan/amass/requests"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type DomainScanController struct {
	beego.Controller
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	HandshakeTimeout:  5 * time.Second,
	// CheckOrigin: 处理跨域问题，线上环境慎用
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端读写消息
type wsMessage struct {
	messageType int
	data        []byte
}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // Mutex互斥锁，避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}


func (wsConn *wsConnection) wsReadLoop() {
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		req := &wsMessage{}
		req = &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			goto closed
		}
	}
	error:
		wsConn.wsClose()
	closed:
		fmt.Println("websocket is closed.")
}

func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *wsConnection) wsClose() {
	wsConn.wsSocket.Close()
	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}

func (wsConn *wsConnection) wsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				goto error
			}
		case <-wsConn.closeChan:
			goto closed
		}
	}
	error:
		wsConn.wsClose()
	closed:
		fmt.Println("websocket is closed.")
}

func (wsConn *wsConnection) procLoop() {
	// 启动一个gouroutine发送心跳
	resOutput := make(chan *requests2.Output,10)


	//go func() {
	//	for {
	//		time.Sleep(3 * time.Second)
	//		if err := wsConn.wsWrite(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
	//			fmt.Println("heartbeat fail")
	//			wsConn.wsClose()
	//			break
	//		}
	//	}
	//}()
	//go func() {
	//	for  {
	//		wsConn.outChan <- &wsMessage{messageType: 1,data: []byte("test")}
	//		time.Sleep(3*time.Second)
	//	}
	//}()

	rootCtx ,cancel := context.WithCancel(context.Background())

	type DomainScanInfo struct {
		Rootdomain string `json:"rootdomain"`
		Passive bool 	`json:"passive"`
		Active bool		`json:"active"`
	}
	var dsi DomainScanInfo
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			fmt.Println("read fail")
			break
		}
		//rootdomain := string(msg.data)
		fmt.Println(string([]byte(msg.data)))
		err = json.Unmarshal(msg.data,&dsi)
		if err != nil{
			continue
		}
		//收到消息后开始暴破域名信息
		go func(output chan *requests2.Output,cancel context.CancelFunc) {
			for {
				select {
				case op := <- output:
					if op != nil{
						var data string
						if dsi.Passive {
							data += op.Name
							err = wsConn.wsWrite(msg.messageType, []byte(data))
							if err != nil {
								fmt.Println("write fail")
								cancel()
								break
							}
							continue
						}else{
							for _,v  :=range op.Sources {
								data += v
							}
							data += "  |   "+op.Name+"     |     "
							for _,v:=range op.Addresses {
								data += v.Address.String()+"  |  "//v.Netblock.String()+"   |   "//+v.Description
							}

							err = wsConn.wsWrite(msg.messageType, []byte(data))
							if err != nil {
								fmt.Println("write fail")
								cancel()
								break
							}
						}
					}

				}
			}
		}(resOutput,cancel)

		if dsi.Passive{
			subdomainscan2.PassiveDomainBrute(rootCtx,resOutput,dsi.Rootdomain)
		}
		if dsi.Active{
			subdomainscan2.ActiveDomainBrute(rootCtx,resOutput,dsi.Rootdomain)
		}
	}
}

func (this *DomainScanController) DomainScan()   {
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
	go wsConn.procLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()

}