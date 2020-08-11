package config

type ConfigAPI struct {
	Host string `yaml:"host" env:"API_HOST" env-default:"127.0.0.1"`
	Port uint   `yaml:"port" env:"API_PORT" env-default:"8080"`
}
