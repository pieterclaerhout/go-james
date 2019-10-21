package common

import (
	"path/filepath"
	"strings"
)

// PackageNameToShort converts a package name to it's short version
func PackageNameToShort(packageName string) string {
	name := filepath.Base(packageName)
	name = strings.TrimPrefix(name, "go-")
	name = strings.TrimSuffix(name, "go-")
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ToLower(name)
	return name
}
