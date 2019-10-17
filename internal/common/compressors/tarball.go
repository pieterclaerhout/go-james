package compressors

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-log"
	"github.com/pkg/errors"
)

// TarballCompressor is used to create tgz files
type TarballCompressor struct {
	path    string
	entries []fileEntry
}

// NewTarball returns a new TarballCompressor instance
func NewTarball(path string) *TarballCompressor {
	return &TarballCompressor{
		path:    path,
		entries: []fileEntry{},
	}
}

// Path returns the path to the tgz file
func (archive *TarballCompressor) Path() string {
	return archive.path
}

// AddFile adds the file from path as name to the archive
//
// If name is not specified, the basename of path is used
func (archive *TarballCompressor) AddFile(name string, path string) {
	if name == "" {
		name = filepath.Base(path)
	}
	entry := fileEntry{
		Name: name,
		Path: path,
	}
	log.Debug(entry, "Adding:")
	archive.entries = append(archive.entries, entry)
}

// Close creates and closes the archive
func (archive *TarballCompressor) Close() error {

	if len(archive.entries) == 0 {
		return errors.New("No files found to compress")
	}

	file, err := os.Create(archive.path)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for _, entry := range archive.entries {
		if err := archive.addFileToTarWriter(entry, tarWriter); err != nil {
			return err
		}
	}

	return nil
}

func (archive *TarballCompressor) addFileToTarWriter(entry fileEntry, tarWriter *tar.Writer) error {

	file, err := os.Open(entry.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    entry.Name,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	if _, err := io.Copy(tarWriter, file); err != nil {
		return err
	}

	return nil

}
