package main

import (
	"log"
	"send-sms/internal/config"
	"send-sms/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	r := routes.SetupRouter(cfg)
	addr := ":" + cfg.Port
	log.Printf("✅ Send-SMS microservice listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("❌ Failed to run server: %v", err)
	}
}
