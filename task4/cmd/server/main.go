package main

import (
	"blog/internal/app"
	"blog/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("../../configs")
	if err != nil {
		log.Fatal("加载配置失败:", err)
	}

	app.Run(cfg)
}
