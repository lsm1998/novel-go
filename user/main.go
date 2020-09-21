package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
	"user/config"
	"user/service"
	"utils"
)

func main() {
	if err := utils.ScanConfig(&config.Config); err != nil {
		panic(err)
	}
	server.UsePool = true
	newServer := server.NewServer()
	addRegistryPlugin(newServer)
	_ = newServer.RegisterName(config.Config.Rpc.Server, new(service.UserServer), "")
	listenRpc(newServer)
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + fmt.Sprintf(`:%d`, config.Config.Rpc.Port),
		ConsulServers:  config.Config.Registry.Adders,
		BasePath:       config.Config.Rpc.Server,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		panic(err)
	}
	s.Plugins.Add(r)
}

func listenRpc(s *server.Server) {
	_ = s.Serve("tcp", fmt.Sprintf(`:%d`, config.Config.Rpc.Port))
}
