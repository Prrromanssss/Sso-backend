package main

import (
	"fmt"

	"github.com/Prrromanssss/sso-backend/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger

	// TODO: init app

	// TODO: start gRPC-server for the app
}
