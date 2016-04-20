package main

import (
	"flag"
	"log"

	"github.com/iguiyu/demo/api/configs"
	"github.com/iguiyu/microservices/api/routers"
)

var (
	env = flag.String("env", "dev", "Running Environment")
)

func main() {
	flag.Parse()
	configs.AppConf.LoadConfigAndSetupEnv(*env)
	router := routers.MakeHandlersWithRouter()
	log.Printf("Listening and serving HTTP on %s\n", configs.AppConf.HttpPort)
	router.Run(configs.AppConf.HttpPort)
}
