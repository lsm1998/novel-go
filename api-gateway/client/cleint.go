package client

import (
	"api-gateway/config"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	ImClient client.XClient
)

func Init() {
	list := []client.XClient{ImClient}

	for i, v := range config.Config.Gateway.ServiceList {
		temp := client.NewConsulDiscovery(v, v, config.Config.Registry.Adders, nil)
		option := client.DefaultOption
		option.SerializeType = protocol.ProtoBuffer
		list[i] = client.NewXClient(v, client.Failtry, client.RoundRobin, temp, option)
	}
}
