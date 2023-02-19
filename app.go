package truequeslib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type App struct {
	Name    string
	Version string
	Config  Config
	Server  Server `yaml:"server"`
}

type Server struct {
	HttpServer   *http.Server
	Router       *gin.Engine
	GinMode      string `yaml:"gin_mode"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type BaseResponse struct {
	Version string `json:"version"`
	Error   string `json:"error"`
}

func NewApp(name string, version string) (App, error) {
	f, err := os.OpenFile(name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return App{}, fmt.Errorf("err-setting_logFile: %w", err)
	}

	log.SetOutput(f)
	log.Println("Init app", version)

	app := App{
		Name:    name,
		Version: version,
	}

	err = app.LoadConfig()
	if err != nil {
		return App{}, fmt.Errorf("err-load_config: %w", err)
	}

	return app, nil
}

func (api *App) LoadConfig() error {
	cfgFilename := api.Name + ".yml"
	yamlFile, err := os.ReadFile(cfgFilename)
	if err != nil {
		return fmt.Errorf("err-reading_yamlFile %s: %w", cfgFilename, err)
	}

	err = yaml.Unmarshal(yamlFile, &api)
	if err != nil {
		return fmt.Errorf("err-unmarshalling_yamlFile %s: %w", cfgFilename, err)
	}

	gin.SetMode(api.Server.GinMode)

	router := gin.Default()
	httpServer := &http.Server{
		Addr:         ":" + strconv.Itoa(api.Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(api.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(api.Server.WriteTimeout) * time.Second,
	}

	api.Server.Router = router
	api.Server.HttpServer = httpServer

	return nil
}

func (api *App) GetBaseReponse(err string) BaseResponse {
	return BaseResponse{
		Version: api.Version,
		Error:   err,
	}
}
