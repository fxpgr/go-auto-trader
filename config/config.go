package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/airking05/go-auto-trader/models"

	"gopkg.in/yaml.v2"
)

type DB struct {
	UserID   string `yaml:"user_id"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type Config struct {
	Debug        bool              `yaml:"debug"`
	DBConnection string            `yaml:"db_connection"`
	TraderConfig models.TraderGorm `yaml:"trader_config"`
}

func ReadConfig(path string) *Config {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open config: %s", err))
	}
	defer f.Close()

	return ReadConfigReader(f)
}

func ReadConfigReader(reader io.Reader) *Config {
	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(fmt.Sprintf("failed to read config: %s", err))
	}

	var config Config
	err = yaml.Unmarshal(bs, &config)
	if err != nil {
		panic(fmt.Sprintf("failed to parse config: %s", err))
	}

	return &config
}
