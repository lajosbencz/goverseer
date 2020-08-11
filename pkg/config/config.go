package config

type Config struct {
	API     ConfigAPI     `yaml:"api"`
	Log     ConfigLog     `yaml:"log"`
	Manager ConfigManager `yaml:"manager"`
	Include ConfigInclude `yaml:"include"`
}
