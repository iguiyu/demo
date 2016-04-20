package global

import "github.com/chanxuehong/wechat.v2/mp/core"

const (
	WxAppId     = "wxe41d25cf540e80be"
	WxAppSecret = "35c40cf7e02bb00fff7622b03f8ef458"

	WxOriId = "gh_3dc347709f3d"
	WxToken = "park"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	MsgHandler core.Handler
	MsgServer  *core.Server
)
var (
	AccessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(WxAppId, WxAppSecret, nil)
	WechatClient      *core.Client           = core.NewClient(AccessTokenServer, nil)
)
var Ticket []string = []string{
	"gQE_8DoAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2drVGpSRVBtRmU5c2RGcW56R3hIAAIE58AMVwMEAAAAAA==",
	"gQHQ8DoAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL1pVVFRiZVBtTmU5TUdyMmUtR3hIAAIEmB8XVwMEAAAAAA==",
}
