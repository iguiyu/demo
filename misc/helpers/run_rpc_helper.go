package helpers

import (
	"log"
	"net"
	"net/rpc"

	"github.com/kobeld/goutils"
)

type rpcEngine struct {
	rpc.Server
	runtime *Runtime // Add runtime support to change env
}

func NewRpcEngine(cs ConfigShell, env string) *rpcEngine {
	// Load the app configuration
	cs.LoadConfigAndSetupEnv(env)

	return &rpcEngine{
		runtime: &Runtime{Config: cs},
	}
}

func (this *rpcEngine) Run() {

	// Register the runtime rpc methods for changing env when runing tests
	this.RegisterName("Runtime", this.runtime)

	var port = this.runtime.Config.GetRpcPort()
	log.Printf("Listening and serving TCP on %s\n", port)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if goutils.HasErrorAndPrintStack(err) {
			continue
		}

		go this.ServeConn(conn)
	}
}

const (
	// RPC method names
	RPC_RUNTIME_SWITCH_ENV = "Runtime.SwitchEnv"
	RPC_RUNTIME_DROP_DB    = "Runtime.DropCurrentDB"
)

// App's configuration which implements this interface can use the Runtime RPC directly
type ConfigShell interface {
	LoadConfigAndSetupEnv(env string)
	GetRpcPort() string
}

// The Runtime rpc object that provides exported methods to change the runtime environments.
// It is useful for changing the env for testing
type Runtime struct {
	Config ConfigShell
}

func (this *Runtime) SwitchEnv(req string, res *string) (err error) {
	this.Config.LoadConfigAndSetupEnv(req)
	return
}

// This is used to drop the test db when tearing down
func (this *Runtime) DropCurrentDB(req string, res *int) (err error) {
	DropDB()
	return
}
