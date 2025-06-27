package main

import (
	"fmt"

	"github.com/chillyweather/gator/internal/cli"
	"github.com/chillyweather/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println(cfg)
}
