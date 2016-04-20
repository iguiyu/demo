package model

import (
	"time"

	"github.com/iguiyu/microservices/misc/global"
	"github.com/kobeld/goutils"
	"gopkg.in/mgo.v2/bson"
)

type Base struct {
	Id        bson.ObjectId `bson:"_id" json:"-"`
	CreatedAt time.Time     `bson:",omitempty" json:"-"`
	UpdatedAt time.Time     `bson:",omitempty" json:"-"`
	DeletedAt time.Time     `bson:",omitempty" json:"-"`
}

func (this *Base) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func (this *Base) SetTime() {
	current := time.Now()
	if this.CreatedAt.IsZero() {
		this.CreatedAt = current
		this.UpdatedAt = current
	} else {
		this.UpdatedAt = current
	}
}

func (this *Base) IsFound() bool {
	return this.Id.Valid()
}

// Fields mapping for copying
func (this *Base) IdHex() string {
	return this.Id.Hex()
}

func (this *Base) CreatedAtStr() string {
	return goutils.FormatTime(this.CreatedAt, global.TIME_LAYOUT_DEFAULT)
}

func (this *Base) UpdatedAtStr() string {
	return goutils.FormatTime(this.UpdatedAt, global.TIME_LAYOUT_DEFAULT)
}

func (this *Base) DeletedAtStr() string {
	return goutils.FormatTime(this.DeletedAt, global.TIME_LAYOUT_DEFAULT)
}

func (this *Base) CreatedAtTimestamp() string {
	return goutils.TimeToMillisecond(this.CreatedAt)
}
