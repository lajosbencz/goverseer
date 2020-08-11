package config

type ConfigLog struct {
	Path  string `yaml:"path" env:"LOG_PATH" env-default:"goverseer.log"`
	Level string `yaml:"level" env:"LOG_LEVEL" env-default:"debug"`
	Size  string `yaml:"size" env:"LOG_SIZE" env-default:"5M"`
}
