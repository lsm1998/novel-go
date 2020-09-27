package main

import (
	"book/config"
	"book/service"
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
	"utils"
)

func main() {
	if err := utils.ScanConfig(&config.Config); err != nil {
		panic(err)
	}
	config.Init()
	server.UsePool = true
	newServer := server.NewServer()
	addRegistryPlugin(newServer)
	_ = newServer.RegisterName(config.Config.Rpc.Server, new(service.BookServer), "")
	_ = newServer.Serve("tcp", fmt.Sprintf(`:%d`, config.Config.Rpc.Port))
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: fmt.Sprintf(`tcp@:%d`, config.Config.Rpc.Port),
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
