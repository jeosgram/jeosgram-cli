package types

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
