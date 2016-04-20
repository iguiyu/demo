package model

import (
	"time"

	"github.com/kobeld/goutils"

	"github.com/kobeld/mgowrap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BaseRepository interface {
	Get(query map[string]interface{}, result interface{}) error
	GetById(id string, result interface{}) error
	GetFieldsById(id string, fields []string, result interface{}) error
	GetAndUpdateById(id string, changer map[string]interface{}, result interface{}) error
	GetAll(result interface{}, sortFields ...string) error
	GetSome(query map[string]interface{}, result interface{}, sortFields ...string) error
	GetSomeByIds(ids []string, result interface{}, sortFields ...string) error
	Save(po mgowrap.PersistentObject, funcs ...func()) error
	Delete(po mgowrap.PersistentObject) error
	HasAny(po mgowrap.PersistentObject, query map[string]interface{}) (bool, error)
	HasAnyById(po mgowrap.PersistentObject, id string) (bool, error)
}

// The Mongo db implementation of BaseRepository
type BaseRepo struct {
	DB *mgowrap.Database
}

func (this *BaseRepo) Get(query map[string]interface{}, result interface{}) (err error) {
	var selector = bson.M{}
	for key, value := range query {
		selector[key] = value
	}
	return this.DB.Find(query, result)
}

func (this *BaseRepo) GetAll(result interface{}, sortFields ...string) (err error) {
	return this.DB.FindAll(bson.M{}, result, sortFields...)
}

func (this *BaseRepo) GetSome(query map[string]interface{}, result interface{}, sortFields ...string) (err error) {
	var q = bson.M{}
	for key, value := range query {
		q[key] = value
	}

	return this.DB.FindAll(q, result, sortFields...)
}

func (this *BaseRepo) GetSomeByIds(ids []string, result interface{}, sortFields ...string) (err error) {

	oIds, err := goutils.TurnPlainIdsToObjectIds(ids)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	query := bson.M{"_id": bson.M{"$in": oIds}}

	return this.DB.FindAll(query, result, sortFields...)
}

func (this *BaseRepo) GetById(id string, result interface{}) (err error) {
	return this.DB.FindByIdHex(id, result)
}

func (this *BaseRepo) GetFieldsById(id string, fields []string, result interface{}) (err error) {

	var (
		query    = bson.M{"_id": bson.ObjectIdHex(id)}
		selector = bson.M{}
	)

	for _, field := range fields {
		selector[field] = 1
	}

	err = this.DB.FindAndSelect(query, selector, result)

	return
}

func (this *BaseRepo) Save(po mgowrap.PersistentObject, funcs ...func()) (err error) {
	return this.DB.Save(po, funcs...)
}

func (this *BaseRepo) Delete(po mgowrap.PersistentObject) (err error) {
	return this.DB.DeleteInstance(po)
}

func (this *BaseRepo) GetAndUpdateById(id string, changer map[string]interface{}, result interface{}) (err error) {

	var (
		selector = bson.M{"_id": bson.ObjectIdHex(id)}
		setValue = bson.M{"updatedAt": time.Now()}
	)

	for key, value := range changer {
		setValue[key] = value
	}

	change := mgo.Change{
		Update:    bson.M{"$set": setValue},
		ReturnNew: true,
	}

	_, err = this.DB.FindAndApply(selector, change, result)

	return
}

func (this *BaseRepo) HasAny(po mgowrap.PersistentObject, query map[string]interface{}) (bool, error) {
	selector := bson.M{}
	for key, value := range query {
		selector[key] = value
	}

	return this.DB.HasAny(po, selector)
}

func (this *BaseRepo) HasAnyById(po mgowrap.PersistentObject, id string) (bool, error) {
	return this.DB.HasAny(po, bson.M{"_id": bson.ObjectIdHex(id)})
}
