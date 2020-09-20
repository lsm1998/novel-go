package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"proto/example"
)

var host = flag.String("s", "127.0.0.1:8972", "server ip and port")
var etcdAddr = "127.0.0.1:2379"

func main() {
	servicePath := "Hello"
	serviceMethod := "SayHello"

	flag.Parse()

	// dis := client.NewPeer2PeerDiscovery(*host, "")
	dis := client.NewEtcdDiscovery("Hello", servicePath, []string{etcdAddr}, nil)

	option := client.DefaultOption
	option.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient(servicePath, client.Failtry, client.RoundRobin, dis, option)
	defer xclient.Close()

	var req = example.HelloMessageReq{Name: "刘时明"}
	var rsp example.HelloMessageRsp
	err := xclient.Call(context.Background(), serviceMethod, &req, &rsp)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Result)
}
