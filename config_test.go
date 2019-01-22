package goverseer

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := ReadConfig("daemon/goverseer.toml")
	if err != nil {
		t.Error(err)
	}
	if cfg.Title != "My Awesome Project" {
		t.Error("Title invalid")
	}
	if cfg.Http.Host != "" {
		t.Error("Http.Host invalid")
	}
	if cfg.Http.Port != 8176 {
		t.Error("Http.Port invalid")
	}
	fmt.Printf("%#v\n", cfg)
}
