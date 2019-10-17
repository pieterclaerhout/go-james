package compressors

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-log"
	"github.com/pkg/errors"
)

// ZipCompressor is used to create zip files
type ZipCompressor struct {
	path    string
	entries []fileEntry
}

// NewZip returns a new ZipCompressor instance
func NewZip(path string) *ZipCompressor {
	return &ZipCompressor{
		path:    path,
		entries: []fileEntry{},
	}
}

// Path returns the path to the zip file
func (archive *ZipCompressor) Path() string {
	return archive.path
}

// AddFile adds the file from path as name to the archive
//
// If name is not specified, the basename of path is used
func (archive *ZipCompressor) AddFile(name string, path string) {
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
func (archive *ZipCompressor) Close() error {

	if len(archive.entries) == 0 {
		return errors.New("No files found to compress")
	}

	file, err := os.Create(archive.path)
	if err != nil {
		return err
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	for _, entry := range archive.entries {
		if err := archive.addFileToZipWriter(entry, zipWriter); err != nil {
			return err
		}
	}

	return nil

}

func (archive *ZipCompressor) addFileToZipWriter(entry fileEntry, zipWriter *zip.Writer) error {

	file, err := os.Open(entry.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(stat)
	if err != nil {
		return err
	}

	header.Name = entry.Name
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	if _, err := io.Copy(writer, file); err != nil {
		return err
	}

	return nil

}
