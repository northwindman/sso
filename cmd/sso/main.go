package main

import (
	"fmt"
	"github.com/northwindman/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: initial logger

	// TODO: initial application

	// TODO: запустить gRPC-сервер приложения
}
