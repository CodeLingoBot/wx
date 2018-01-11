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
	Curr   float64 `json:"curr"`
	God    float64 `json:"god"`
	Step   int     `json:"step"`
	Action string  `json:"action"`
	Diff   float64 `json:"diff"`
	Now    string  `json:"now"`
	Buy    int     `json:"buy"`
	Sell   int     `json:"sell"`
}
