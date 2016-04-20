package handlers

import (
	"log"

	"code.iguiyu.com/parking/wechat/handlers"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"github.com/gin-gonic/gin"
)

func TextMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)
	// re, err := json.Marshal(ctx)
	// if err != nil {
	//  fmt.Println(err)
	// }

	// fmt.Println(string(re))
	// handlers.BoradcastMessage(string(re))
	handlers.BoradcastMessage(string(ctx.MsgPlaintext))

	msg := request.GetText(ctx.MixedMsg)

	// resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	articles := []response.Article{{Title: "title1", Description: "description", PicURL: "https://www.baidu.com/img/bd_logo1.png", URL: "https://www.baidu.com/"}}

	resp := response.NewNews(msg.FromUserName, msg.ToUserName, msg.CreateTime, articles)
	ctx.RawResponse(resp) // 明文回复
	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func DefaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)

	ctx.NoneResponse()
}

func MenuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)
	event := menu.GetClickEvent(ctx.MixedMsg)
	var desc string
	switch event.EventKey {
	case "scan":
		desc = "停车"
	case "park":
		desc = "停车"
	default:
		desc = "sorry I don't know"

	}
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件"+desc)

	ctx.RawResponse(resp) // 明文回复
	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func DefaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	event := menu.GetClickEvent(ctx.MixedMsg)
	var desc string
	switch event.EventType {
	case "subscribe":
		subevent := request.GetSubscribeEvent(ctx.MixedMsg)
		desc = "谢谢关注" + subevent.EventKey
	case "unsubscribe":
		desc = "取消关注"
	case "SCAN":
		scanEvent := request.GetScanEvent(ctx.MixedMsg)
		desc = "车位编码：" + scanEvent.EventKey + "\n当前时段：12：00-20：00\n当前时段价格：20元/小时\n当前时段封顶：100元\n停车回复是，否则回复"
	case "scancode_waitmsg":
		scanEvent := request.GetScanEvent(ctx.MixedMsg)
		desc = "车位编码：" + scanEvent.EventKey + "\n当前时段：12：00-20：00\n当前时段价格：20元/小时\n当前时段封顶：100元\n停车回复是，否则回复"
	default:
		desc = "sorry I don't know"

	}
	log.Println(desc)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, desc)

	ctx.RawResponse(resp) // 明文回复
}

func PushMsg(c *gin.Context) {
	// var text = mass2users.NewText([]string{"ogwRov8pR7MzlZf6HeUgPnpjBVOk", "ogwRov8iiVgCt5WO0frm-XQb6vaE"}, "hello I am message from xiaoyu")
	// rslt, err := mass2users.Send(wechatClient, text)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println(rslt.MsgId)
	// }
}
