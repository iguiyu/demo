package handlers

import (
	"log"
	"net/http"

	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/gin-gonic/gin"
	"github.com/iguiyu/demo/deployment/global"
	"github.com/iguiyu/demo/deployment/model"
	"github.com/kobeld/mgowrap"
	"gopkg.in/mgo.v2/bson"
)

func CreateMenu(c *gin.Context) {
	var btn1 menu.Button
	btn1.SetAsScanCodeWaitMsgButton("扫码", "park")
	var btn2 menu.Button
	btn2.SetAsLocationSelectButton("车位号", "location")
	var btn3 menu.Button
	btn3.SetAsClickButton("充值", "encharge")
	var btn4 menu.Button
	btn4.SetAsClickButton("我的", "my")

	btns := []menu.Button{btn1, btn2, btn3, btn4}
	var menu_ menu.Menu
	menu_.Buttons = btns
	err := menu.Create(global.WechatClient, &menu_)
	var wxMenu = &model.WXMenu{
		Id: bson.NewObjectId(),
		Menu: menu.Menu{
			Buttons: btns,
		},
	}

	log.Println("创建菜单 ", err)
	err = mgowrap.Save(wxMenu)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Println("保存menu")
}
func _createMenu(c *gin.Context) {
	var btn1 menu.Button
	var btn1_sub1 menu.Button
	var btn1_sub2 menu.Button
	btn1_sub1.SetAsScanCodePushButton("push", "push")
	btn1_sub2.SetAsScanCodeWaitMsgButton("waitmsg", "waitmsg")
	var btn1_subs = []menu.Button{btn1_sub1, btn1_sub2}
	btn1.SetAsSubMenuButton("扫毛", btn1_subs)
	var btn2 menu.Button
	btn2.SetAsLocationSelectButton("车位号", "location")
	btns := []menu.Button{btn1, btn2}

	var menu_ menu.Menu
	menu_.Buttons = btns
	err := menu.Create(global.WechatClient, &menu_)
	var wxMenu = &model.WXMenu{
		Id: bson.NewObjectId(),
		Menu: menu.Menu{
			Buttons: btns,
		},
	}

	log.Println("创建菜单 ", err)
	err = mgowrap.Save(wxMenu)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Println("保存menu")

}
