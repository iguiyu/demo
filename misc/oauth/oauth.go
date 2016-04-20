package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/RangelReale/osin"
	"github.com/iguiyu/microservices/misc/global"
	"github.com/kobeld/goutils"
)

var oauthAddress = ""

// The offical client
var OfficalClient osin.Client = &osin.DefaultClient{
	Id:          "4fb4b98c6fece5320cbbdae4",
	Secret:      "5018d345558fbe46c4000001",
	RedirectUri: "https://club.iguiyu.com",
}

func SetOauthAddress(addr string) {
	oauthAddress = addr
}

func GrantAccessTokenWithClientCredentials(userId string) (responseData osin.ResponseData, err error) {

	url := fmt.Sprintf("http://%s/token?grant_type=client_credentials&user_id=%s", oauthAddress, userId)

	request, err := http.NewRequest("GET", url, nil)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	request.SetBasicAuth(OfficalClient.GetId(), OfficalClient.GetSecret())

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	if response.StatusCode != 200 {
		err = errors.New("Invalid status code " + response.Status)
		return
	}

	responseData = make(osin.ResponseData)

	jdec := json.NewDecoder(response.Body)
	err = jdec.Decode(&responseData)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	return
}

func GetAccess(token string) (access *osin.AccessData, err error) {
	url := fmt.Sprintf("http://%s/access?token=%s", oauthAddress, token)

	request, err := http.NewRequest("GET", url, nil)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	switch response.StatusCode {
	case 200:
		// Correct, do nothing.
	case 400:
		return nil, global.ErrInvalidToken
	default:
		return nil, errors.New("Invalid status code: " + response.Status)
	}

	access = new(osin.AccessData)

	jdec := json.NewDecoder(response.Body)
	err = jdec.Decode(&access)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	return
}
