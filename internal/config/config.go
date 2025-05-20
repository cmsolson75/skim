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

	_ = viper.ReadInConfig()

	cfg := &Config{
		InputDir:    viper.GetString("input_dir"),
		OutputDir:   viper.GetString("output_dir"),
		OutputName:  viper.GetString("output_name"),
		AllowedExts: viper.GetStringSlice("allowed_extensions"),
		SkipDirs:    viper.GetStringSlice("skip_dirs"),
	}

	// Validate required fields
	if cfg.InputDir == "" {
		return nil, fmt.Errorf("input_dir is required")
	}
	if cfg.OutputDir == "" {
		cfg.OutputDir = "./"
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
	viper.SetDefault("input_dir", ".")
	viper.SetDefault("output_dir", ".")
	viper.SetDefault("output_name", "context.txt")
	viper.SetDefault("allowed_extensions", []string{".go", ".py", ".yaml", ".txt"})
	viper.SetDefault("skip_dirs", []string{".git", "__pycache__"})
	viper.SetDefault("cloc", false)
	viper.SetDefault("tree", false)
}
