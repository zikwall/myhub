package config

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/zikwall/myhub/pkg/database"
)

type Config struct {
	Server `yaml:"server"`
}

type Server struct {
	Prefork  bool         `yaml:"prefork"`
	Database database.Opt `yaml:"database"`
	Telegram Telegram     `yaml:"telegram"`
}

type Telegram struct {
	BotKey   string `yaml:"bot_key"`
	BotDebug bool   `yaml:"bot_debug"`
}

func New(filepath string) (*Config, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	d := yaml.NewDecoder(bytes.NewReader(content))
	if err = d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
