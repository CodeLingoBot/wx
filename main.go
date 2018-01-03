package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/morya/utils/log"
)

var (
	flagListenAddr = flag.String("listen", "0.0.0.0:7400", "listen address")
)

const (
	myself = "o-A1l03mCLRjTP09Z6UZdOVLUBLs"
)

func GetHandler(c echo.Context) error {
	echostr := c.FormValue("echostr")
	return c.String(http.StatusOK, echostr)
}

func Parse(m *StoreMsg, val string, buy, sell string) {
	err := json.Unmarshal([]byte(val), m)
	if err != nil {
		log.Info(err)
	}
	buyInt, _ := strconv.Atoi(buy)
	sellInt, _ := strconv.Atoi(sell)

	m.Buy = buyInt
	m.Sell = sellInt
}

func GetLastStatus() string {
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	var m = &StoreMsg{}
	jsonData, _ := c.Get("bch_usdt_last").Result()
	buy, _ := c.Get("bch_usdt_buy").Result()
	sell, _ := c.Get("bch_usdt_sell").Result()

	Parse(m, jsonData, buy, sell)

	bindata, _ := json.MarshalIndent(m, "", "  ")
	return string(bindata)
}

func PostHandler(c echo.Context) error {
	bindata, _ := ioutil.ReadAll(c.Request().Body)

	recv := &WxAutoMsg{}
	send := &WxAutoMsg{}

	err := xml.Unmarshal(bindata, recv)
	if err != nil {
		log.Info(err)
		return c.String(http.StatusOK, "success")

	}
	log.Infof("xml unmashal succ %v", recv)
	send.FromUserName = recv.ToUserName
	send.ToUserName = recv.FromUserName
	send.MsgId = recv.MsgId
	send.CreateTime = recv.CreateTime
	send.MsgType = recv.MsgType

	if recv.FromUserName != myself {
		send.Content = fmt.Sprintf("bye bye")
	} else {
		send.Content = GetLastStatus()
	}

	log.Infof(" build resp %#v", send)
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
