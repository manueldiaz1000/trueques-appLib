package truequeslib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAppOK(t *testing.T) {
	app, err := NewApp("test", "0.0.1")
	assert.NoError(t, err)
	assert.Equal(t, "test", app.Name)
	assert.Equal(t, "0.0.1", app.Version)
	assert.Equal(t, 8870, app.Server.Port)
	assert.Equal(t, "release", app.Server.GinMode)
	assert.Equal(t, ":8870", app.Server.HttpServer.Addr)
	assert.Equal(t, time.Duration(15*time.Second), app.Server.HttpServer.ReadTimeout)
	assert.Equal(t, time.Duration(15*time.Second), app.Server.HttpServer.WriteTimeout)
	assert.NotEmpty(t, app.Server.Router)

	err = app.Run()
	assert.NoError(t, err)
}

func TestNewApp_DefaultConfigOK(t *testing.T) {
	app, err := NewApp("testDefault", "0.0.1")
	assert.NoError(t, err)
	assert.Equal(t, 8080, app.Server.Port)
	assert.Equal(t, ":8080", app.Server.HttpServer.Addr)
	assert.Equal(t, time.Duration(10*time.Second), app.Server.HttpServer.ReadTimeout)
	assert.Equal(t, time.Duration(10*time.Second), app.Server.HttpServer.WriteTimeout)
}

func TestNewApp_LogFileEmptyError(t *testing.T) {
	app, err := NewApp("", "0.0.1")
	assert.Empty(t, app)
	assert.Contains(t, err.Error(), "err-log_filename_empty")
}

func TestNewApp_LogFileInvalidError(t *testing.T) {
	app, err := NewApp("/", "0.0.1")
	assert.Empty(t, app)
	assert.Contains(t, err.Error(), "err-creating_logFile")
}

func TestNewApp_YmlFileNotFoundError(t *testing.T) {
	app, err := NewApp("unknown", "0.0.1")
	assert.Empty(t, app)
	assert.Contains(t, err.Error(), "err-load_config")
	assert.Contains(t, err.Error(), "err-reading_yamlFile")
}

func TestNewApp_YmlFileUnmarshallError(t *testing.T) {
	app, err := NewApp("testError", "0.0.1")
	assert.Empty(t, app)
	assert.Contains(t, err.Error(), "err-load_config")
	assert.Contains(t, err.Error(), "err-unmarshalling_yamlFile")
}

func TestGetBaseResponseOK(t *testing.T) {
	app, err := NewApp("test", "0.0.1")
	assert.NoError(t, err)

	response := app.GetBaseReponse("err-not_found")
	assert.Equal(t, "0.0.1", response.Version)
	assert.Equal(t, "err-not_found", response.Error)
}

func TestGetBaseResponseError(t *testing.T) {
	app, err := NewApp("test", "0.0.1")
	assert.NoError(t, err)

	response := app.GetBaseReponse("")
	assert.NotEqual(t, "0.0.0", response.Version)
	assert.Empty(t, "", response.Error)
}

func TestRunAppOK(t *testing.T) {
	app, err := NewApp("testPortInvalid", "0.0.1")
	assert.NoError(t, err)

	err = app.Run()
	assert.Contains(t, err.Error(), "invalid port")
}
