package goverseer

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title    string
	Http     ConfigHttp
	Include  string
	Programs []ConfigProgram
}

type ConfigHttp struct {
	Host string
	Port uint16
}

type ConfigProgram struct {
	Title        string
	Command      string
	ProcessName  string
	NumProcs     uint
	Directory    string
	Prioirity    uint
	AutoStart    bool
	StartSecs    uint
	StartRetries uint
	User         string
	Group        string
	StopSignal   string
}

func ReadConfig(path string) (*Config, error) {
	var err error
	if blob, err := ioutil.ReadFile(path); err == nil {
		var cfg Config
		if _, err := toml.Decode(string(blob), &cfg); err != nil {
			return nil, err
		}

		if cfg.Include != "" {
			var includeDir string
			if !filepath.IsAbs(cfg.Include) {
				dir := filepath.Dir(path)
				includeDir = fmt.Sprintf("%s/%s/*.toml", cfg.Include, dir)
			} else {
				includeDir = fmt.Sprintf("%s/*.toml", cfg.Include)
			}
			files, err := filepath.Glob(includeDir)
			if err != nil {
				return nil, err
			}
			for _, v := range files {
				cfgProgram, err := readConfigProgram(v)
				if err != nil {
					return nil, err
				}
				cfg.Programs = append(cfg.Programs, *cfgProgram)
			}
		}
		return &cfg, nil
	}
	return nil, err

}

func readConfigProgram(path string) (*ConfigProgram, error) {
	var err error
	if blob, err := ioutil.ReadFile(path); err != nil {
		var cfg ConfigProgram
		if _, err := toml.Decode(string(blob), &cfg); err != nil {
			return nil, err
		}
		return &cfg, nil
	}
	return nil, err
}
