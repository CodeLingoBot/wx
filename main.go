package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
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

func Parse(m *StoreMsg, val string, buy, sell string) {
	err := json.Unmarshal([]byte(val), m)
	if err != nil {
		log.InfoErrorf(err, "%s", val)
	}
	buyInt, _ := strconv.Atoi(buy)
	sellInt, _ := strconv.Atoi(sell)

	m.Buy = buyInt
	m.Sell = sellInt
}

func GetLastStatus(symbol string) string {
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	var m = &StoreMsg{}

	if symbol == "" {
		return "unknown symbol"
	}
	key_last := fmt.Sprintf("%s_last", symbol)
	key_buy := fmt.Sprintf("%s_buy", symbol)
	key_sell := fmt.Sprintf("%s_sell", symbol)

	jsonData, _ := c.Get(key_last).Result()
	buy, _ := c.Get(key_buy).Result()
	sell, _ := c.Get(key_sell).Result()

	Parse(m, jsonData, buy, sell)

	bindata, _ := json.MarshalIndent(m, "", "  ")
	return string(bindata)
}

func DumpObj(obj interface{}) {
	data, _ := json.Marshal(obj)
	log.Infof("%s", data)
}

func OnContent(user string, content string) (reply string) {
	if !IsValidUser(user) {
		reply = "bye bye"
		return
	}
    i, err := strconv.ParseInt(content, 10, 32)
    if err ==nil {

	content = strings.ToLower(content)
	switch {
	case strings.Contains(content, "ltc"):
		reply = GetLastStatus("ltc_usdt")

	case strings.Contains(content, "bch"):
		reply = GetLastStatus("bch_usdt")

	default:
		reply = GetLastStatus("eos_usdt")
	}
	return
}

func OnTextMsg(recv *WxAutoMsg) (send *WxAutoMsg) {
	send = &WxAutoMsg{}
	send.FromUserName = recv.ToUserName
	send.ToUserName = recv.FromUserName
	send.MsgID = recv.MsgID
	send.CreateTime = recv.CreateTime
	send.MsgType = recv.MsgType

	send.Content = OnContent(recv.FromUserName, recv.Content)
	return
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
