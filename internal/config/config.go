package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	LinkedIn struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"linkedin"`
	App struct {
		Headless bool   `yaml:"headless"`
		LogLevel string `yaml:"log_level"`
		DataDir  string `yaml:"data_dir"`
	} `yaml:"app"`
	Stealth struct {
		MinDelay int `yaml:"min_delay"`
		MaxDelay int `yaml:"max_delay"`
	} `yaml:"stealth"`
	Search struct {
		Keywords string `yaml:"keywords"`
		Limit    int    `yaml:"limit"`
	} `yaml:"search"`
}

func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{}

	// Defaults
	cfg.App.Headless = false
	cfg.App.LogLevel = "info"
	cfg.App.DataDir = "./data"
	cfg.Stealth.MinDelay = 2
	cfg.Stealth.MaxDelay = 5

	// Attempt to load from config.yaml if it exists
	f, err := os.Open("config.yaml")
	if err == nil {
		defer f.Close()
		decoder := yaml.NewDecoder(f)
		if err := decoder.Decode(cfg); err != nil {
			return nil, fmt.Errorf("failed to decode config.yaml: %w", err)
		}
	}

	// Override with Environment Variables
	if v := os.Getenv("LINKEDIN_USERNAME"); v != "" {
		cfg.LinkedIn.Username = v
	}
	if v := os.Getenv("LINKEDIN_PASSWORD"); v != "" {
		cfg.LinkedIn.Password = v
	}

	if v := os.Getenv("HEADLESS"); v != "" {
		cfg.App.Headless = (v == "true")
	}
	if v := os.Getenv("LOG_LEVEL"); v != "" {
		cfg.App.LogLevel = v
	}
	if v := os.Getenv("DATA_DIR"); v != "" {
		cfg.App.DataDir = v
	}
	if v := os.Getenv("MIN_DELAY"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			cfg.Stealth.MinDelay = i
		}
	}
	if v := os.Getenv("MAX_DELAY"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			cfg.Stealth.MaxDelay = i
		}
	}
	if v := os.Getenv("SEARCH_KEYWORDS"); v != "" {
		cfg.Search.Keywords = v
	}
	if v := os.Getenv("SEARCH_LIMIT"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			cfg.Search.Limit = i
		}
	}

	// Defaults if missing
	if cfg.Search.Keywords == "" {
		cfg.Search.Keywords = "Software Engineer"
	}
	if cfg.Search.Limit == 0 {
		cfg.Search.Limit = 10
	}

	return cfg, nil
}
