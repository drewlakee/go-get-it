package mongo

import (
	"fmt"
	"go-get-it/pkg/config"
)

func BuildConnectionUrl(cfg *config.ScriptConfig) string {
	exit := "exit"

	if cfg.Host == "" {
		fmt.Println("Add 'host' to config file")
		return exit
	}

	if cfg.Port == 0 {
		fmt.Println("Add 'port' to config file")
		return exit
	}

	if cfg.Username == "" {
		fmt.Println("Add 'username' to config file")
		return exit
	}

	if cfg.Password == "" {
		fmt.Println("Add 'password' to config file")
		return exit
	}

	url := fmt.Sprintf("mongodb://%s:%s@%s:%d/",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port)

	if cfg.Database != "" {
		url = fmt.Sprintf("%s%s", url, cfg.Database)
	}

	if cfg.CAFile != "" {
		url = fmt.Sprintf("%s?tls=true&tlsCaFile=%s", url, cfg.CAFile)
	}

	return url
}
