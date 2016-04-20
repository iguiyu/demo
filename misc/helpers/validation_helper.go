package helpers

import (
	"fmt"
	"time"

	"github.com/iguiyu/microservices/misc/global"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
)

func init() {
	// Setup validator custom validation functions
	validator.SetValidationFunc("bsonId", bsonId)
	validator.SetValidationFunc("blankId", blankId)
	validator.SetValidationFunc("timeFormat", timeFormat)
}

type ValidatorErr struct {
	ErrorMap map[string]string
}

func (this *ValidatorErr) ValidField(value interface{}, name string, tags string) {
	err := validator.Valid(value, tags)
	if err != nil {
		this.AddError(name, err)
	}
}

func (this *ValidatorErr) HasError() bool {
	return len(this.ErrorMap) > 0
}

func (this *ValidatorErr) HasErrorOn(name string) (yes bool) {
	if this.ErrorMap == nil {
		return false
	}
	_, yes = this.ErrorMap[name]
	return
}

func (this *ValidatorErr) AddError(name string, err error) {
	if err == nil {
		return
	}

	if this.ErrorMap == nil {
		this.ErrorMap = make(map[string]string)
	}

	_, ok := this.ErrorMap[name]
	if !ok {
		this.ErrorMap[name] = err.Error()
	}
	return
}

func (this *ValidatorErr) AddNotFoundError() {
	this.AddError("id", global.ErrNotFound)
}

func (this *ValidatorErr) ClearErrors() {
	this.ErrorMap = map[string]string{}
}

func (this *ValidatorErr) HasErrorOnAndReset(name string) (yes bool) {
	yes = this.HasErrorOn(name)
	this.ClearErrors()
	return
}

func (this *ValidatorErr) ToString() string {
	return fmt.Sprintf("%+v", this.ErrorMap)
}

func (this *ValidatorErr) ReturnErrorMap() map[string]string {
	return this.ErrorMap
}

// ----- Custom validation functions -----
func bsonId(v interface{}, param string) error {
	id := v.(bson.ObjectId)
	if !id.Valid() {
		return global.ErrInvalidId
	}
	return nil
}

func blankId(v interface{}, param string) error {
	id := v.(bson.ObjectId)
	if id != "" {
		return global.ErrShouldBeBlank
	}
	return nil
}

func timeFormat(v interface{}, param string) error {
	timeStr := v.(string)
	_, err := time.Parse(param, timeStr)
	if err != nil {
		return global.ErrInvalidTimeFormat
	}

	return nil
}
