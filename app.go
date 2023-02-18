package truequeslib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Name    string
	Version string
	Config  Config
	Server  *http.Server
	Router  *gin.Engine
}

type BaseResponse struct {
	Version string `json:"version"`
	Error   string `json:"error"`
}

func NewApi(name string, version string) (Api, error) {
	f, err := os.OpenFile(name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return Api{}, fmt.Errorf("err-setting_logFile: %w", err)
	}

	log.SetOutput(f)
	log.Println("Init app", version)

	config, err := LoadConfig(name + ".yml")
	if err != nil {
		return Api{}, fmt.Errorf("err-init_config: %w", err)
	}

	gin.SetMode(config.GinMode)
	router := gin.Default()
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(config.Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
	}

	return Api{
		Name:    name,
		Version: version,
		Router:  router,
		Server:  server,
	}, nil
}

func (api *Api) GetBaseReponse(err string) BaseResponse {
	return BaseResponse{
		Version: api.Version,
		Error:   err,
	}
}
