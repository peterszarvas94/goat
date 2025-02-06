package utils

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func CopyDir(srcDir, dstDir string) error {
	// Get properties of the source directory
	srcInfo, err := os.Stat(srcDir)
	if err != nil {
		return err
	}

	// If the source is a directory, create the destination directory
	if err := os.MkdirAll(dstDir, srcInfo.Mode()); err != nil {
		return err
	}

	// Walk through the source directory and copy files/subdirectories
	return filepath.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Construct the destination path
		relPath, err := filepath.Rel(srcDir, srcPath)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			// Create the directory at the destination
			return os.MkdirAll(dstPath, info.Mode())
		}

		// Copy the file
		return CopyFile(srcPath, dstPath)
	})
}
