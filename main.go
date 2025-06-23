package main

import (
	"fmt"

	"github.com/chillyweather/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	err = cfg.SetUser("Dima")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	cfg, _ = config.Read()
	fmt.Println(cfg)
}
