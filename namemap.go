package main

var (
    IDMap = map[int]string{
        1:"eos_usdt",
        2:"ltc_usdt",
        3:"bch_usdt",
        4:"etc_usdt",
    }
)
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

