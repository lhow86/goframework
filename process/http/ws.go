package http

import (
	"github.com/gin-gonic/gin"
	"goframework/process/controller"
	"strconv"
)

/*
本项目的具体处理逻辑都会在process/controller下进行，而接口层面则是在http和rpc目录下去分发。
websocket是一种协议，一种在http上升级的协议，
所以它并不像我们之前开启http服务一样，直接监听端口，而是在http服务上开一个接口，通过接口去实现服务端与客户端的连接
*/
func Ws(ctx *gin.Context) {
	//websocket是一个协议，一个http之上的协议，对http“升级”了。
	//代码中or循环，相当于http的监听。为了统一，调用controller.GetServerTime()来获取系统时间，这就让http接口和ws接口都统一了，真正的实现逻辑都在controller。
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		if string(message) == "server_time" {
			resp, _, _ := controller.GetServerTime()
			serverTime := strconv.FormatInt(resp.ServerTime, 10)
			message = []byte(serverTime)
		}
		//写入数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
