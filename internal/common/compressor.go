package common

import (
	"github.com/pieterclaerhout/go-james/internal/common/compressors"
)

// Compressor is what can be injected into a subcommand when you need zip / tgz functions
type Compressor struct {
	Logging
}

// CreateTgz creates a tgz file from a path
func (c Compressor) CreateTgz(srcPaths []string, tgzPath string) error {
	engine := compressors.NewTarball(tgzPath)
	return c.createArchive(engine, srcPaths)
}

// CreateZip creates a zip file from a path
func (c Compressor) CreateZip(srcPaths []string, zipPath string) error {
	engine := compressors.NewZip(zipPath)
	return c.createArchive(engine, srcPaths)
}

func (c Compressor) createArchive(engine compressors.Compressor, srcPaths []string) error {
	for _, srcPath := range srcPaths {
		engine.AddFile("", srcPath)
	}
	c.LogPathCreation(engine.Path())
	return engine.Close()
}
