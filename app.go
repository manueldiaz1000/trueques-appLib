package truequeslib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type App struct {
	Name    string
	Version string
	Server  Server                 `yaml:"server"`
	Params  map[string]interface{} `yaml:"params"`
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
	if strings.Trim(name, " ") == "" {
		return App{}, ErrLogFilenameEmpty
	}

	f, err := os.OpenFile(name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return App{}, fmt.Errorf(ErrCreatingLogFile.Error()+": %w", err)
	}

	log.SetOutput(f)
	log.Println("Init app", version)

	app := App{
		Name:    name,
		Version: version,
	}

	err = app.LoadConfig()
	if err != nil {
		return App{}, fmt.Errorf(ErrLoadingConfig.Error()+": %w", err)
	}

	return app, nil
}

func (app *App) LoadConfig() error {
	cfgFilename := app.Name + ".yml"
	yamlFile, err := os.ReadFile(cfgFilename)
	if err != nil {
		return fmt.Errorf("err-reading_yamlFile %s: %w", cfgFilename, err)
	}

	err = yaml.Unmarshal(yamlFile, &app)
	if err != nil {
		return fmt.Errorf("err-unmarshalling_yamlFile %s: %w", cfgFilename, err)
	}

	router := gin.Default()
	if app.Server.GinMode == "" {
		app.Server.GinMode = "debug"
		gin.SetMode(app.Server.GinMode)
	}

	if app.Server.Port == 0 {
		app.Server.Port = 8080
	}

	if app.Server.ReadTimeout == 0 {
		app.Server.ReadTimeout = 10
	}

	if app.Server.WriteTimeout == 0 {
		app.Server.WriteTimeout = 10
	}

	httpServer := &http.Server{
		Addr:         ":" + strconv.Itoa(app.Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(app.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.Server.WriteTimeout) * time.Second,
	}

	app.Server.Router = router
	app.Server.HttpServer = httpServer

	return nil
}

func (app *App) Run() error {
	log.Println("Running app in port " + app.Server.HttpServer.Addr)
	err := app.Server.HttpServer.ListenAndServe()
	return err
}

func (app *App) GetBaseReponse(err string) BaseResponse {
	return BaseResponse{
		Version: app.Version,
		Error:   err,
	}
}
