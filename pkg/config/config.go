package config

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"os"
)

type ScriptConfig struct {
	Host     string `config:"host"`
	Port     int    `config:"port"`
	Database string `config:"database"`
	Username string `config:"username"`
	Password string `config:"password"`
	CAFile   string `config:"caFile"`
}

func LoadScriptConfig() *ScriptConfig {
	configPath := fmt.Sprintf("%s/.ggi/config.json", os.Getenv("HOME"))
	loader := confita.NewLoader(file.NewBackend(configPath))
	cfg := &ScriptConfig{}
	err := loader.Load(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
