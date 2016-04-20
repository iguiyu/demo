package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type WXAccessToken struct {
	Id           bson.ObjectId `bson:"_id"`
	Access_token string
	Expires_in   int
	Created_at   time.Time
}

func (this *WXAccessToken) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func (this *WXAccessToken) CollectionName() string {
	return "wx_access_token"
}

func (this *WXAccessToken) ToString() string {
	return "access_token = " + this.Access_token + " created at" + this.Created_at.Format("2006-01-02 15:04:05")
}
