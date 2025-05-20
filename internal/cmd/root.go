package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cmsolson75/skim/internal/config"
	"github.com/cmsolson75/skim/internal/di"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "skim",
	Short: "Skim aggregates project context for easier ingestion",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}
		if val, err := cmd.Flags().GetString("input-dir"); err == nil && val != "" {
			cfg.InputDir = val
		}
		if val, err := cmd.Flags().GetString("output-dir"); err == nil && val != "" {
			cfg.OutputDir = val
		}
		if val, err := cmd.Flags().GetString("output-name"); err == nil && val != "" {
			cfg.OutputName = val
		}
		if val, err := cmd.Flags().GetString("allowed-extensions"); err == nil && val != "" {
			cfg.AllowedExts = splitAndTrim(val)
		}
		if val, err := cmd.Flags().GetString("skip-dirs"); err == nil && val != "" {
			cfg.SkipDirs = splitAndTrim(val)
		}

		c := di.New(cfg)
		files, err := c.Walker.Walk()
		if err != nil {
			return err
		}
		return c.Output.Write(files)
	},
}

func init() {
	rootCmd.PersistentFlags().String("input-dir", "", "Directory to scan")
	rootCmd.PersistentFlags().String("output-dir", "", "Output directory")
	rootCmd.PersistentFlags().String("output-name", "", "Output file name")
	rootCmd.PersistentFlags().String("allowed-extensions", "", "Allowed file extensions (comma-separated)")
	rootCmd.PersistentFlags().String("skip-dirs", "", "Directories to skip")

	rootCmd.PersistentFlags().Bool("cloc", false, "Include cloc output")
	rootCmd.PersistentFlags().Bool("tree", false, "Include tree output")
	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.AutomaticEnv()
}

func splitAndTrim(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
