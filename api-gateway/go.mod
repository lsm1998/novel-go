module api-gateway

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/smallnest/rpcx v0.0.0-20200917102714-42a82be8f8ab
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	proto v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

replace proto => ../proto
