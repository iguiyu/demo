package reqres

import (
	"github.com/iguiyu/microservices/misc/model"
	"gopkg.in/mgo.v2/bson"
)

type NestedUserReq struct {
	Id     bson.ObjectId
	Name   string
	Avatar string
}

func (this NestedUserReq) ToNestedUser() model.NestedUser {
	return model.NestedUser{
		Id:     this.Id,
		Name:   this.Name,
		Avatar: this.Avatar,
	}
}
