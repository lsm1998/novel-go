package main

import (
	"context"
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"proto/example"
	"time"
)

type Hello int

func (t *Hello) SayHello(ctx context.Context, req *example.HelloMessageReq, rsp *example.HelloMessageRsp) error {
	rsp.Result = "hello:" + req.Name
	return nil
}

var (
	host = flag.String("s", "127.0.0.1:8972", "listened ip and port")

	etcdAddr = "127.0.0.1:2379"
)

func main() {
	flag.Parse()
	server.UsePool = true
	newServer := server.NewServer()
	addRegistryPlugin(newServer)
	_ = newServer.RegisterName("Hello", new(Hello), "")
	_ = newServer.Serve("tcp", *host)
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + *host,
		EtcdServers:    []string{etcdAddr},
		BasePath:       "Hello",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		panic(err)
	}
	s.Plugins.Add(r)
}
