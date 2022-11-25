package main

import (
	"context"
	"fmt"

	"election_algorithm/pkg/global"
)

func main() {
	fmt.Printf("%#v\n", global.GetConfig())
	cli := global.GetRedisCli()
	fmt.Println(cli.Ping(context.Background()).Result())
}
