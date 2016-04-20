package main

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/iguiyu/demo/deployment/global"
	"github.com/iguiyu/demo/deployment/handlers"
	"github.com/kobeld/mgowrap"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(handlers.DefaultMsgHandler)
	mux.DefaultEventHandleFunc(handlers.DefaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, handlers.TextMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, handlers.MenuClickEventHandler)

	global.MsgHandler = mux
	var wxEncodedAESKey string
	global.MsgServer = core.NewServer(global.WxOriId, global.WxAppId, global.WxToken, wxEncodedAESKey, global.MsgHandler, nil)

}

func getCallBackIp(w http.ResponseWriter, r *http.Request) {
	// log.Println(base.GetCallbackIP(wechatClient))

}
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func wxCallbackHandler(c *gin.Context) {

	global.MsgServer.ServeHTTP(c.Writer, c.Request, nil)
}

func broad(c *gin.Context) {
	handlers.BoradcastMessage("I am message from parking server\n")
}

func main() {
	handlers.LongCon()
	// Init Gin and setup the router
	mgowrap.SetupDatbase("localhost", "userdb")

	router := gin.Default()
	// router.LoadHTMLGlob("template/*")

	{
		// router.POST("/menu", createMenu)
		// router.GET("/menu", listMenu)
		// router.GET("/menu/:id", getOneMenu)
		// router.PATCH("", ...)
		// router.DELETE(relativePath, ...)
		router.GET("/createMenu", handlers.CreateMenu)
		router.GET("/wechatcreateQrcode", handlers.CreateQrcode)
		router.GET("/getQrcode", handlers.GetQrcode)
		router.GET("/toQrcode", handlers.ToQrcode)
		router.GET("/pushMsg", handlers.PushMsg)
		router.GET("/broad", broad)
		router.Any("/wx_callback", wxCallbackHandler)

		// /wechat/
		// router.GET("/", index)
	}

	log.Printf("Listening and serving http on %s\n", ":8080")
	router.Run(":8080")
	// log.Println(http.ListenAndServe(":8080", nil))
}
