package main

import (
	"fmt"
	"github.com/mini-ecs/back-end/internal/router"
	"github.com/mini-ecs/back-end/pkg/config"
	"github.com/mini-ecs/back-end/pkg/log"
)

var logger = log.GetGlobalLogger()

func main() {
	logger.Infof("Initializing project, config: %+v", config.GetConfig())
	logger.Infof("Starting server...")
	r := router.NewRouter()
	err := r.Run(fmt.Sprintf(":%v", config.GetConfig().Service.Port))
	if err != nil {
		return
	}
}
