package controllers

import (
	"fmt"
	nethttp "net/http"

	"github.com/goravel/framework/contracts/http"
	"github.com/gorilla/websocket"
)

type WebsocketController struct {
	//Dependent services
}

func NewWebsocketController() *WebsocketController {
	return &WebsocketController{
		//Inject services
	}
}

func (r *WebsocketController) Server(ctx http.Context) http.Response {
	upGrader := websocket.Upgrader{
		ReadBufferSize:  4096, // Specify the read buffer size
		WriteBufferSize: 4096, // Specify the write buffer size
		// Detect request origin
		CheckOrigin: func(r *nethttp.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				fmt.Println("path error")
				return false
			}
			return true
		},
	}

	ws, err := upGrader.Upgrade(ctx.Response().Writer(), ctx.Request().Origin(), nil)
	if err != nil {
		return ctx.Response().String(http.StatusInternalServerError, err.Error())
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		fmt.Println("Received:", string(message))
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("asede")
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}

	return nil
}
