package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// SystemConfigs get from .yml
type SystemConfigs struct {
	DBStock struct {
		Type     string `yaml:"type" envconfig:"DB_TYPE"`
		Host     string `yaml:"host" envconfig:"DB_HOST"`
		Name     string `yaml:"name" envconfig:"DB_NAME"`
		User     string `yaml:"user" envconfig:"DB_USER"`
		Password string `yaml:"pass" envconfig:"DB_PASSWORD"`
	} `yaml:"db_stock"`
	DBTrend struct {
		Type     string `yaml:"type" envconfig:"DB_TYPE"`
		Host     string `yaml:"host" envconfig:"DB_HOST"`
		Name     string `yaml:"name" envconfig:"DB_NAME"`
		User     string `yaml:"user" envconfig:"DB_USER"`
		Password string `yaml:"pass" envconfig:"DB_PASSWORD"`
	} `yaml:"db_trend"`
	Crawler struct {
		SeleniumPath    string `yaml:"seleniumPath" envconfig:"crawler_seleniumPath"`
		GeckoDriverPath string `yaml:"geckoDriverPath" envconfig:"crawler_geckoDriverPath"`
		Port            int    `yaml:"port" envconfig:"crawler_port"`
		Delay           int    `yaml:"delay" envconfig:"crawler_delay"`
	} `yaml:"crawler"`
	Log struct {
		Path string `yaml:"path" envconfig:"LOG_PATH"`
		Name string `yaml:"name" envconfig:"LOG_NAME"`
	} `yaml:"log"`
}

// WebConfigs get from .yml
type WebConfigs struct {
	Server struct {
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	File struct {
		Billing struct {
			Path string `yaml:"path" envconfig:"FILE_BILLING_PATH"`
		} `yaml:"billing"`
	} `yaml:"file"`
}

// CollectorConfigs get from .yml
type CollectorConfigs struct {
	CloudWatch struct {
		Region string `yaml:"region" envconfig:"CW_REGION"`
		API    string `yaml:"api" envconfig:"CW_API"`
		User   string `yaml:"user" envconfig:"CW_USER"`
	} `yaml:"cloudwatch"`
}

// BillingConfigs get from .yml
type BillingConfigs struct {
	InitState string `yaml:"InitState"`
}

var configs *SystemConfigs

// Load the configurations
func Load(path string) *SystemConfigs {
	cfg := SystemConfigs{}
	f, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Panic(err)
	}

	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Panic(err)
	}
	configs = &cfg
	return configs
}

// Get the configurations
func Get() *SystemConfigs {
	return configs
}
