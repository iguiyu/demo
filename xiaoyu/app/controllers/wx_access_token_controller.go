package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/iguiyu/demo/xiaoyu/app/consts"
	"github.com/iguiyu/demo/xiaoyu/app/model"
	"github.com/kobeld/mgowrap"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WXAccessToken struct {
	*revel.Controller
}

func (c WXAccessToken) Fuck() revel.Result {
	fmt.Println(" Fuck")
	return c.RenderJson("FUCK")
}

func (w WXAccessToken) AccessToken() revel.Result {

	fmt.Println("I am app AccessToken")
	var result = w.GetAccessToken()
	fmt.Println("______________", result.ToString())
	w.RenderArgs["accesstoken"] = result

	return w.RenderJson(result)
}
func (w WXAccessToken) CheckAccessToken() revel.Result {
	w.AccessToken()
	return nil
}
func (w WXAccessToken) generateAccessToken() model.WXAccessToken {
	var url = strings.Replace(strings.Replace(consts.Url_get_accesstoken, "APPID", consts.AppId, 1), "APPSECRET", consts.Secret, 1)
	fmt.Println(">>>>>>", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(">>>>>>", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(">>>>>>", err)
	}
	respStr := string(body)
	fmt.Println(">>>>>>", respStr)
	var at model.WXAccessToken
	err = json.Unmarshal([]byte(respStr), &at)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(">>>>>> accesstoken", at.Access_token)
	fmt.Println(">>>>>> expire in", at.Expires_in)
	at.Created_at = time.Now()
	err = mgowrap.Save(&at)
	if err != nil {
		fmt.Println("error:", err)
	}
	return at
}
func (w WXAccessToken) GetAccessToken() model.WXAccessToken {
	var (
		err error
		at  model.WXAccessToken
	)
	err = mgowrap.FindWithLimit(bson.M{}, &at, 1, "created_at")
	if err != nil {
		at = w.generateAccessToken()
		fmt.Println("new token = ", at.ToString())

	}
	fmt.Println("already has token , and it is ", at.ToString())
	return at
}
