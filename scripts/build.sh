#!bin/bash
export GO111MODULE=on

#Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/gexpose-linux-amd64 ./main.go
#Linux ARM
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/gexpose-linux-arm64 ./main.go
#Mac OS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/gexpose-darwin-amd64 ./main.go
#Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/gexpose-windows-amd64.exe ./main.go
#OpenWrt
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o ./bin/gexpose-openwrt-amd64 ./main.go

echo "DONE!!!"
