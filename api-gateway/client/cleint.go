package client

import (
	"api-gateway/config"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	ImClient   client.XClient
	UserClient client.XClient
	BaseClient client.XClient
	BookClient client.XClient
)

func Init() {
	clients := []*client.XClient{&ImClient, &UserClient, &BaseClient, &BookClient}

	for i, v := range config.Config.Gateway.List {
		dis := client.NewConsulDiscovery(v, v, config.Config.Registry.Adders, nil)
		option := client.DefaultOption
		option.SerializeType = protocol.ProtoBuffer
		*(clients[i]) = client.NewXClient(v, client.Failtry, client.RoundRobin, dis, option)
	}
}
