package main

import (
	"encoding/xml"
)

type WxAutoMsg struct {
	XMLName xml.Name `xml:"xml"`

	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgID        string `xml:"MsgId"`
}

type StoreMsg struct {
	Symbol string  `json:"symbol"`
	Curr   float64 `json:"curr"`
	God    float64 `json:"god"`
	Step   float64 `json:"step"`
	Action string  `json:"action"`
	Diff   float64 `json:"diff"`
	Now    string  `json:"now"`
	Buy    int     `json:"buy"`
	Sell   int     `json:"sell"`
}

type BuzzStoreMsg struct {
	Symbol     string  `json:"symbol"`
	GodPrice   float64 `json:"god_price"`
	LastPrice  float64 `json:"last_price"`
	ChangeRate float64 `json:"rate"` //default 5%
	Amount     float64 `json:"amount"`
	BuyCount   int64   `json:"buy_count"`
	SellCount  int64   `json:"sell_count"`
}
