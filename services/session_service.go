package services

import (
	"os"
	"path/filepath"

	"github.com/jeosgram/jeosgram-cli/types"
	"github.com/jeosgram/jeosgram-cli/utils"
	"gopkg.in/yaml.v3"
)

type SessionService interface {
	Clean() error
	ReadConfig() (*types.Config, error)
	SaveConfig(conf *types.Config) error
	SaveTokens(token *types.Token) error
}

type FileBasedAuthentication struct {
	configPath string
}

func (service FileBasedAuthentication) homePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return utils.ExecutableDir() // ruta del ejecutable
	}
	return home, nil
}

func (service FileBasedAuthentication) ensureFolder() (string, error) {
	home, err := service.homePath()
	if err != nil {
		return "", err
	}

	jeosgramDir := filepath.Join(home, ".jeosgram")
	if !utils.IsExistPath(jeosgramDir) {
		if err := os.Mkdir(jeosgramDir, 0755); err != nil {
			return "", err
		}
	}

	return jeosgramDir, nil
}

func (service FileBasedAuthentication) Clean() error {
	jeosgramDir, _ := service.ensureFolder()
	return os.RemoveAll(jeosgramDir)
}

func (service FileBasedAuthentication) ReadConfig() (*types.Config, error) {
	jeosgramDir, err := service.ensureFolder()
	if err != nil {
		return nil, err
	}

	service.configPath = filepath.Join(jeosgramDir, "config.yml")
	data, _ := os.ReadFile(service.configPath)

	var conf types.Config

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func (service FileBasedAuthentication) SaveConfig(conf *types.Config) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return os.WriteFile(service.configPath, data, 0644)
}

func (service FileBasedAuthentication) SaveTokens(token *types.Token) error {
	conf, _ := service.ReadConfig()
	conf.Token.AccessToken = token.AccessToken
	conf.Token.RefreshToken = token.RefreshToken
	return service.SaveConfig(conf)
}
