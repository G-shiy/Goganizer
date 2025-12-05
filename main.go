package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/G-shiy/goganizer/handlers"
)

//go:embed rules/rules.json
var defaultRules []byte

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	rules, err := handlers.LoadRules("./rules/rules.json", defaultRules)
	if err != nil {
		log.Fatalf("\033[31mError loading rules: %v\033[0m\n", err)
	}
	fmt.Printf("\033[36mOrganizing files in: %s\033[0m\n", dir)

	if err := handlers.OrganizeFiles(dir, rules); err != nil {
		log.Fatalf("\033[31mError organizing files: %v\033[0m\n", err)
	}
	fmt.Println("\033[32mOrganization completed successfully!\033[0m")
}
