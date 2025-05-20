package di

import (
	"github.com/cmsolson75/skim/internal/config"
	"github.com/cmsolson75/skim/internal/output"
	"github.com/cmsolson75/skim/internal/walker"
)

type Container struct {
	Config *config.Config
	Walker *walker.Service
	Output *output.Service
}

func New(cfg *config.Config) *Container {
	return &Container{
		Config: cfg,
		Walker: walker.New(cfg),
		Output: output.New(cfg),
	}
}
