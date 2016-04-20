package model

import (
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"gopkg.in/mgo.v2/bson"
)

type WXMenu struct {
	Id bson.ObjectId `bson:"_id"`
	menu.Menu
}

func (this *WXMenu) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}
func (this *WXMenu) CollectionName() string {
	return "menus"
}
