package main

import (
	"fmt"
	"log"

	"github.com/G-shiy/GoGanizer/handlers"
)

func main() {
	rules, err := handlers.LoadRules("./rules/rules.json")
	if err != nil {
		log.Fatal("Error loading rules")
	}
	if err := handlers.OrganizeFiles(".", rules); err != nil {
		log.Fatal("Error organize files")
	}
	fmt.Println("Organização concluida com sucesso")
}
