package dir

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"

	log "github.com/sirupsen/logrus"
)

// DirEntry is a struct that represents a directory
type DirEntry struct {
	// Path is the path to the directory
	Path string `json:"path"`
	// Name is the name of the directory
	Name string `json:"name"`
}

// FileEntry is a struct that represents a file
type FileEntry struct {
	// Path is the path to the directory
	Path string `json:"path"`
	// Name is the name of the directory
	Name string `json:"name"`
	// SHA256 is the SHA256 hash of the file
	SHA256 string `json:"sha256"`
}

// NewDirEntry creates a new instance of the Dir struct
func NewDirEntry(location string) *DirEntry {
	return &DirEntry{
		Path: location,
		Name: path.Base(location),
	}
}

// Path returns the path to the directory
func (d *DirEntry) Files() []*FileEntry {
	files, err := os.ReadDir(d.Path)
	if err != nil {
		log.WithError(err).Fatal("Failed to read directory")
	}

	var fileEntries []*FileEntry
	for _, e := range files {
		fileEntries = append(fileEntries, &FileEntry{
			Path:   path.Join(d.Path, e.Name()),
			Name:   e.Name(),
			SHA256: shasum(path.Join(d.Path, e.Name())),
		})
	}

	return fileEntries
}

// Path returns the path to the directory
func (d *DirEntry) SnapshotDirs() []*DirEntry {
	matcher := regexp.MustCompile(`^db\d+$`)

	var snapshots []*DirEntry
	err := filepath.WalkDir(d.Path, func(location string, info os.DirEntry, err error) error {
		if !info.IsDir() {
			return nil
		}

		baseName := path.Base(location)
		if !matcher.Match([]byte(baseName)) {
			return nil
		}

		snapshots = append(snapshots, &DirEntry{
			Path: location,
			Name: baseName,
		})
		return nil
	})
	if err != nil {
		log.WithError(err).Fatal("Failed to walk directory")
	}

	return snapshots

}

func shasum(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("0x%x", h.Sum(nil))
}
