package main

import (
	"fmt"
	_ "github.com/mini-ecs/back-end/docs"
	"github.com/mini-ecs/back-end/internal/router"
	"github.com/mini-ecs/back-end/pkg/config"
	"github.com/mini-ecs/back-end/pkg/log"
)

var logger = log.GetGlobalLogger()

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	logger.Infof("Initializing project, config: %+v", config.GetConfig())
	logger.Infof("Starting server...")
	r := router.NewRouter()
	err := r.Run(fmt.Sprintf(":%v", config.GetConfig().Service.Port))
	if err != nil {
		return
	}
}
