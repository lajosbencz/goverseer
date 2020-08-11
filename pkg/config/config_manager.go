package config

type ConfigManager struct {
	Umask string `yaml:"umask" env:"MANAGER_UMASK" env-default:"022"`
	User  string `yaml:"user" env:"MANAGER_USER" env-default:""`
}
