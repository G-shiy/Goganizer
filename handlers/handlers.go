package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func LoadRules(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	var rules map[string]string
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON from %s: %w", path, err)
	}

	return rules, nil
}

func OrganizeFiles(dir string, rules map[string]string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := filepath.Ext(entry.Name())

		destDir, exists := rules[ext]
		if !exists {
			continue
		}

		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			if err := os.MkdirAll(destDir, 0o755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", destDir, err)
			}
			fmt.Printf("Created directory: %s\n", destDir)
		}

		srcPath := filepath.Join(dir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if err := moveFile(srcPath, destPath); err != nil {
			return fmt.Errorf("failed to move file from %s to %s: %w", srcPath, destPath, err)
		}
		fmt.Printf("Moved file: %s to %s\n", srcPath, destPath)

	}
	return nil
}

func moveFile(srcPath, destPath string) error {
	if err := os.Rename(srcPath, destPath); err == nil {
		return nil
	}

	inFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		return err
	}

	return os.Remove(srcPath)
}
