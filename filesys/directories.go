package filesys

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrPathNotExists      = errors.New("path not exists")
	ErrPathIsNotDirectory = errors.New("provided path is not a directory")
)

// IsDirectory return nil if path exists and is a directory.
func IsDirectory(path string) error {
	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return ErrPathNotExists
	}

	if err != nil {
		return fmt.Errorf("unexpected error with path: %w", err)
	}

	if !fi.IsDir() {
		return ErrPathIsNotDirectory
	}

	return nil
}
