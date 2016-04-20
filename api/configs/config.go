package configs

import (
	"fmt"

	"github.com/iguiyu/demo/misc/global"
	"github.com/iguiyu/demo/misc/helpers"
	"github.com/iguiyu/demo/misc/oauth"
)

var (
	AppConf = new(appConf)
)

type appConf struct {
	helpers.BaseConfig `,inline`
}

func (this *appConf) LoadConfigAndSetupEnv(env string) {
	this.loadConfig(env)
	this.setOAuthAddress(env)
}

func (this *appConf) loadConfig(env string) {
	switch env {
	case helpers.ENV_DEV:
		this.BaseConfig = helpers.BaseConfig{
			ServiceName: "api_dev",
			Env:         env,
			HttpPort:    ":3003",
		}

	case helpers.ENV_TEST:
		this.BaseConfig = helpers.BaseConfig{
			ServiceName: "api_test",
			Env:         env,
			HttpPort:    ":3003",
		}

	case helpers.ENV_PROD:
		this.BaseConfig = helpers.BaseConfig{
			ServiceName: "api_prod",
			Env:         env,
			HttpPort:    ":8000",
		}

	case helpers.ENV_CI:
		// TODO:
	default:
		panic("=== Please provide the runtime env (dev/prod/test/ci) ===")
	}

	fmt.Printf("The config is:\n%+v\n\n", this)
}

func (this *appConf) setOAuthAddress(env string) {
	switch env {
	case helpers.ENV_DEV, helpers.ENV_TEST:
		authAddr := global.ServiceMap["auth"]
		oauth.SetOauthAddress(authAddr)
	case helpers.ENV_PROD:
		oauth.SetOauthAddress("auth" + AppConf.HttpPort)
	}
}
