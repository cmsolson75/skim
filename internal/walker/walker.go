package walker

import (
	"os"
	"path/filepath"

	"github.com/cmsolson75/skim/internal/config"
	"github.com/cmsolson75/skim/internal/util"
)

type Service struct {
	cfg *config.Config
}

type FileData struct {
	Path     string
	Contents []string
}

func New(cfg *config.Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Walk() ([]FileData, error) {
	allowed := util.NewSet(s.cfg.AllowedExts)
	skipDirs := util.NewSet(s.cfg.SkipDirs)
	var files []FileData

	err := filepath.Walk(s.cfg.InputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && skipDirs.Has(info.Name()) {
			return filepath.SkipDir
		}
		if info.IsDir() {
			return nil
		}
		if !allowed.Has(filepath.Ext(info.Name())) {
			return nil
		}
		content, err := ReadLines(path)
		if err == nil {
			files = append(files, FileData{Path: path, Contents: content})
		}
		return nil
	})
	return files, err
}
