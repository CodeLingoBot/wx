package main

import (
	"github.com/morya/utils/log"
	"testing"
)

func TestSendMsg(t *testing.T) {
    var bindata = `<xml><ToUserName><![CDATA[gh_24bcfce41811]]></ToUserName>
<FromUserName><![CDATA[o-ABCDEmCLRjTP09Z6UZdOVLUBLs]]></FromUserName>
<CreateTime>1514900785</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[eos]]></Content>
<MsgId>6506449328693625983</MsgId>
</xml>`
	recv := &WxAutoMsg{}
	err := xml.Unmarshal(bindata, recv)
    send := OnTextMsg(recv)
    DumpObj(send)
}
