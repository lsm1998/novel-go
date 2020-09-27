#!/usr/bin/env bash
protoc -I=./ -I=../ --proto_path=C:/Users/Administrator/go/novel-go/proto --go_out=. ./*.proto
#protoc -I=. -I=../../../ --js_out=import_style=commonjs,binary:. *.proto
#pbjs -p=. -p=../ -t json common.proto > common.json
#不需要空
sed -i.bak 's/,omitempty//g' *.pb.go
