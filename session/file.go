package session

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var jeosgramPathConf string

type Token struct {
	AccessToken  string `json:"access_token" yaml:"access_token"`
	RefreshToken string `json:"refresh_token" yaml:"refresh_token"`
}

type Config struct {
	CLI struct { // para cosas de actualizacion
		DisableUpdateCheck bool   `yaml:"disable_update_check"`
		LastVersionCheck   int64  `yaml:"last_version_check"`
		NewerVersion       string `yaml:"newer_version"`
	} `yaml:"cli"`
	Token `yaml:"token"`
}

func homePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return ExecutableDir() // ruta del ejecutable
	}
	return home, nil
}

func EnsureFolder() (string, error) {
	home, err := homePath()
	if err != nil {
		return "", err
	}

	jeosgramDir := filepath.Join(home, ".jeosgram")
	if !IsExistPath(jeosgramDir) {
		if err := os.Mkdir(jeosgramDir, 0755); err != nil {
			return "", err
		}
	}

	return jeosgramDir, nil
}

func Clean() error {
	jeosgramDir, _ := EnsureFolder()
	return os.RemoveAll(jeosgramDir)
}

func ReadConfig() (*Config, error) {
	jeosgramDir, err := EnsureFolder()
	if err != nil {
		return nil, err
	}

	jeosgramPathConf = filepath.Join(jeosgramDir, "config.yml")
	data, _ := os.ReadFile(jeosgramPathConf)

	var conf Config

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func SaveConfig(conf *Config) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return os.WriteFile(jeosgramPathConf, data, 0644)
}

func SaveTokens(token *Token) error {
	conf, _ := ReadConfig()
	conf.Token.AccessToken = token.AccessToken
	conf.Token.RefreshToken = token.RefreshToken
	return SaveConfig(conf)
}
