package utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func CopyFile(srcPath, dstPath string) error {
	inFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	return err
}

func WalkAndCopy(srcDir, dstDir string) fs.WalkDirFunc {
	return func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dstDir, relPath)

		if d.IsDir() {
			if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			if err := CopyFile(path, dstPath); err != nil {
				return err
			} else {
				fmt.Printf("🗿 Copied file: %s\n", dstPath)
			}
		}
		return nil
	}
}

func CopyDir(srcDir, dstDir string) error {
	return filepath.WalkDir(srcDir, WalkAndCopy(srcDir, dstDir))
}
