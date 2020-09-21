module api-gateway

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/smallnest/rpcx v0.0.0-20200917102714-42a82be8f8ab
	utils v0.0.0-00010101000000-000000000000
	proto v0.0.0-00010101000000-000000000000
)

replace utils => ../utils

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

replace proto => ../proto
