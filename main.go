package main

import (
	"fmt"

	"github.com/WadeGulbrandsen/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	cfg.SetUser("wade")
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(cfg)
}
