module user

go 1.15

require (
	github.com/garyburd/redigo v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo v1.8.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/smallnest/rpcx v0.0.0-20200917102714-42a82be8f8ab
	github.com/stvp/tempredis v0.0.0-20181119212430-b82af8480203 // indirect
	gopkg.in/redsync.v1 v1.0.1
	proto v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

replace proto => ../proto
