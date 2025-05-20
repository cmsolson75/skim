package cmd

import (
	"fmt"
	"os"

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
		c := di.New(cfg)
		files, err := c.Walker.Walk()
		if err != nil {
			return err
		}
		return c.Output.Write(files)
	},
}

func init() {
	rootCmd.PersistentFlags().String("input-dir", ".", "Directory to scan")
	rootCmd.PersistentFlags().String("output-dir", ".", "Output directory")
	rootCmd.PersistentFlags().String("output-name", "context.txt", "Output file name")
	rootCmd.PersistentFlags().StringSlice("allowed-extensions", nil, "Allowed file extensions")
	rootCmd.PersistentFlags().StringSlice("skip-dirs", nil, "Directories to skip")

	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.AutomaticEnv()
}
