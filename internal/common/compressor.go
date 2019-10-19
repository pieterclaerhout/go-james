package common

import (
	"github.com/pieterclaerhout/go-james/internal/common/compressors"
)

// Compressor is what can be injected into a subcommand when you need zip / tgz functions
type Compressor struct{}

// CreateTgz creates a tgz file from a path
func (c Compressor) CreateTgz(tgzPath string) compressors.Compressor {
	return compressors.NewTarball(tgzPath)
}

// CreateZip creates a zip file from a path
func (c Compressor) CreateZip(zipPath string) compressors.Compressor {
	return compressors.NewZip(zipPath)
}
