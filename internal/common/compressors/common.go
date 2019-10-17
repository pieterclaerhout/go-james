package compressors

// Compressor defines the interface a compressor needs to implement
type Compressor interface {
	Path() string                     // Path returns the path to the archive
	AddFile(name string, path string) // AddFile adds the file from path as name to the archive
	Close() error                     // Close creates and closes the archive
}

// fileEntry defines a file which needs to be added to an archive
type fileEntry struct {
	Name string // The name of the file in the archive
	Path string // The path to the file
}
