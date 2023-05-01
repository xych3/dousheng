package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfigWithPath("../../config-dev.yaml")
	fmt.Println(Config)
}
