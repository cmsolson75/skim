package output

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cmsolson75/skim/internal/analyze"
	"github.com/cmsolson75/skim/internal/config"
	"github.com/cmsolson75/skim/internal/walker"
	"github.com/spf13/viper"
)

type Service struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Write(files []walker.FileData) error {
	if err := os.MkdirAll(s.cfg.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output_dir: %v", err)
	}
	outPath := filepath.Join(s.cfg.OutputDir, s.cfg.OutputName)
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	// Write root
	fmt.Fprintf(w, "Project Root: %s\n\n", s.cfg.InputDir)
	w.Flush()

	if viper.GetBool("cloc") {
		fmt.Fprintln(w, "---- CLOC ----")
		analyze.RunCloc(s.cfg.InputDir, w)
		fmt.Fprintln(w)
	}
	if viper.GetBool("tree") {
		fmt.Fprintln(w, "---- TREE ----")
		analyze.RunTree(s.cfg.InputDir, w)
		fmt.Fprintln(w)
	}
	for _, file := range files {
		fmt.Fprintf(w, "---- %s ----\n", file.Path)
		for _, line := range file.Contents {
			fmt.Fprintln(w, line)
		}
		fmt.Fprintln(w)
	}
	return w.Flush()
}
