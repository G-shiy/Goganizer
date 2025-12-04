package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	rules, err := LoadRules("./rules/rules.json")
	if err != nil {
		log.Fatal("Error loading rules")
	}
	if err := OrganizeFiles(".", rules); err != nil {
		log.Fatal("Error organize files")
	}
	fmt.Println("Organização concluida com sucesso")
}

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
			continue // Skip directories
		}
		ext := filepath.Ext(entry.Name())

		destDir, exists := rules[ext]
		if !exists {
			continue
		}

		if _, err := os.Stat(filepath.Join(dir, destDir)); os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Join(dir, destDir), 0o755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", destDir, err)
			}
			fmt.Printf("Created directory: %s\n", destDir)
		}
		oldPath := filepath.Join(dir, entry.Name())
		newPath := filepath.Join(destDir, entry.Name())
		if err := MoveFile(oldPath, newPath); err != nil {
			log.Printf("failed to move file %s to %s: %v", oldPath, newPath, err)
			continue
		}
		fmt.Printf("Moved file %s to %s\n", oldPath, newPath)
	}

	return nil
}

func MoveFile(src, dst string) error {
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

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

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return os.Remove(src)
}
