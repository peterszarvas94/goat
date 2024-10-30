package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileWriter struct {
	FolderPath string
	FileName   string
}

func NewFileWriter(folderPath, fileName string) (*FileWriter, error) {
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return nil, err
	}
	return &FileWriter{FolderPath: folderPath, FileName: fileName}, nil
}

func (fileWriter *FileWriter) Write(data []byte) (int, error) {
	filename := fmt.Sprintf("%s-logs.txt", time.Now().Format("2006-01-02"))
	filePath := filepath.Join(fileWriter.FolderPath, filename)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.Write(data)
}
