package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	InputDir    string   `mapstructure:"input_dir"`
	OutputDir   string   `mapstructure:"output_dir"`
	OutputName  string   `mapstructure:"output_name"`
	AllowedExts []string `mapstructure:"allowed_extensions"`
	SkipDirs    []string `mapstructure:"skip_dirs"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.skim")

	setDefaults()

	_ = viper.ReadInConfig() // optional

	cfg := &Config{
		InputDir:    viper.GetString("input-dir"),
		OutputDir:   viper.GetString("output-dir"),
		OutputName:  viper.GetString("output-name"),
		AllowedExts: viper.GetStringSlice("allowed-extensions"),
		SkipDirs:    viper.GetStringSlice("skip-dirs"),
	}

	// Validate required fields
	if cfg.InputDir == "" {
		return nil, fmt.Errorf("input_dir is required")
	}
	if cfg.OutputDir == "" {
		cfg.OutputDir = "./out"
	}
	if cfg.OutputName == "" {
		cfg.OutputName = "context.txt"
	}
	if len(cfg.AllowedExts) == 0 {
		cfg.AllowedExts = []string{".go", ".py", ".yaml", ".txt"}
	}
	if len(cfg.SkipDirs) == 0 {
		cfg.SkipDirs = []string{".git", "__pycache__"}
	}

	return cfg, nil
}

func setDefaults() {
	viper.SetDefault("output-dir", "./out")
	viper.SetDefault("output-name", "context.txt")
	viper.SetDefault("allowed-extensions", []string{".go", ".py", ".yaml", ".txt"})
	viper.SetDefault("skip-dirs", []string{".git", "__pycache__"})
}
