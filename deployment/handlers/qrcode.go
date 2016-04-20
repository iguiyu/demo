package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chanxuehong/wechat.v2/mp/qrcode"
	"github.com/gin-gonic/gin"
	"github.com/iguiyu/demo/deployment/global"
)

func CreateQrcode(c *gin.Context) {
	qrcode, err := qrcode.CreateStrScenePermQrcode(global.WechatClient, "1026")
	// qrcode, err := qrcode.CreateStrScenePermQrcode(wechatClient, "I am qrcode string")
	fmt.Println("qrcode.url = ", qrcode.URL, "   qrcode.ticket = ", qrcode.Ticket)
	if err != nil {
		log.Println(err)
	}
}
func GetQrcode(c *gin.Context) {
	var tmp = qrcode.QrcodePicURL(global.Ticket[1])
	fmt.Println(tmp)
	c.Redirect(http.StatusFound, tmp)
}

func ToQrcode(c *gin.Context) {

	c.Redirect(http.StatusFound, "http://weixin.qq.com/q/gkTjREPmFe9sdFqnzGxH")
}
