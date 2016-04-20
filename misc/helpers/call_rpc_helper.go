package helpers

import (
	"net/http"
	"net/rpc"

	"github.com/gin-gonic/gin"
	"github.com/iguiyu/microservices/misc/global"
	"github.com/kobeld/goutils"
)

type rpcResponse interface {
	HasError() bool
	ReturnErrorMap() map[string]string
}

func CallRpcService(env Env, name, method string, req interface{}, res rpcResponse) (err error) {
	// Get the service TCP client
	client, err := getEnvTcpClient(name, env)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	err = client.Call(method, req, res)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	return
}

func CallRpcServiceWithContext(c *gin.Context, env Env, name, method string, req interface{}, res rpcResponse) bool {

	// Get the service TCP client
	client, err := getEnvTcpClient(name, env)
	if goutils.HasErrorAndPrintStack(err) {
		c.AbortWithError(http.StatusInternalServerError, err)
		return true
	}

	err = client.Call(method, req, res)
	if err != nil {

		// Check if it is not found error
		if err.Error() == "not found" {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			goutils.PrintStackAndError(err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return true
	}

	if res.HasError() {
		c.JSON(http.StatusBadRequest, res.ReturnErrorMap())
		return true
	}

	return false
}

func getEnvTcpClient(serviceName string, env Env) (client *rpc.Client, err error) {

	var addr string

	switch {
	case env.IsDev(), env.IsTest():

		addr, _ = global.ServiceMap[serviceName]
		if addr == "" {
			err = global.ErrServiceNotFound
			return
		}

	case env.IsProd():
		addr = serviceName + ":80"
	default:
		err = global.ErrInvalidEnv
		return
	}

	client, err = rpc.Dial("tcp", addr)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	return
}
