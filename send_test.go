package main

import (
    "fmt"
	"encoding/xml"
	"testing"

    "github.com/morya/utils/log"
)

func TestSendMsg(t *testing.T) {
	var bindata = fmt.Sprintf(`<xml><ToUserName><![CDATA[gh_24bcfce41811]]></ToUserName>
<FromUserName><![CDATA[%s]]></FromUserName>
<CreateTime>1514900785</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[eos]]></Content>
<MsgId>6506449328693625983</MsgId>
</xml>`, myself)
	recv := &WxAutoMsg{}
	err := xml.Unmarshal(([]byte)(bindata), recv)
	if err != nil {
		log.Info(err)
		return
	}
	send := OnTextMsg(recv)
	DumpObj(send)
}
