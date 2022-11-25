package main

import (
	"fmt"

	"election_algorithm/pkg/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", cfg)
}
