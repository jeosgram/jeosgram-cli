package main

import (
	"github.com/jeosgram/jeosgram-cli/cmd"
)

/*

https://words.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/

go build -ldflags "-s -w" -o jeosgram main.go
upx --brute ./jeosgram

GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o jeosgram.exe main.go

*/

func main() {
	cmd.Execute()
}
