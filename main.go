package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/morya/utils/log"
)

var (
	flagListenAddr = flag.String("listen", "0.0.0.0:7400", "listen address")
)

func GetHandler(c echo.Context) error {
	echostr := c.FormValue("echostr")
	return c.String(http.StatusOK, echostr)
}

func DumpObj(obj interface{}) {
	data, _ := json.Marshal(obj)
	log.Infof("%s", data)
}

func PostHandler(c echo.Context) error {
	bindata, _ := ioutil.ReadAll(c.Request().Body)

	recv := &WxAutoMsg{}
	err := xml.Unmarshal(bindata, recv)
	if err != nil {
		log.Info(err)
		return c.String(http.StatusOK, "success")
	}
	DumpObj(recv)
	send := OnTextMsg(recv)
	DumpObj(send)
	return c.XML(http.StatusOK, send)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", GetHandler)
	e.GET("/wx", GetHandler)
	e.POST("/wx", PostHandler)

	// Start server

	e.Logger.Fatal(e.Start(*flagListenAddr))
}
