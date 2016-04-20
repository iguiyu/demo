package model

import "gopkg.in/mgo.v2/bson"

type NestedUser struct {
	Id     bson.ObjectId
	Name   string
	Avatar string `bson:",omitempty"`
}

func (this *NestedUser) IdHex() string {
	return this.Id.Hex()
}

type NestedClub struct {
	Id   bson.ObjectId
	Name string
}

func (this *NestedClub) IdHex() string {
	return this.Id.Hex()
}
