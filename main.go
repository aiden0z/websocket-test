package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// Event describe the client control message.
const (
	EventClose = "event.close"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Read
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Errorf("read error: %s", err)
			break
		}
		c.Logger().Infof("recv: %s", msg)

		if string(msg) == EventClose {
			c.Logger().Info("send close message called by client")
			ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		} else {
			// Write
			err = ws.WriteMessage(messageType, msg)
			if err != nil {
				c.Logger().Errorf("wirte error: %s", err)
				break
			}
		}
	}

	return nil
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "./views")
	e.Static("/static", "./static")
	e.GET("/ws", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
