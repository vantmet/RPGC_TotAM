package main

import (
	"fmt"
	"github.com/vantmet/RPGC_TotAM/internal/menu"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	s "strings"
)

type Character struct {
	ID   string `yaml:"id"`
	Name string `yaml: "name"`
}

func main() {
	characters := loadCharacters()
	fmt.Println(characters)

	menu := menu.NewMenu("Chose an option:")

	for _, c := range characters {
		menu.AddItem(c.Name, fmt.Sprint(c.ID))
	}

	menu.AddItem("Create a Character", "create")
	menu.AddItem("Exit", "exit")

	choice := menu.Display()

	switch choice {
	case "create":
		fmt.Printf("You wnat to create a character!")
	case "exit":
		return
	}
}

func loadCharacters() []Character {
	var c []Character
	path := os.Getenv("HOME") + "/.RPGC_TotAM"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
	} else {
		os.Chdir(path)
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, entry := range files {
			n := entry.Name()
			if s.HasSuffix(s.ToLower(n), ".yml") {
				log.Println("Found Character: ", n)
				file, err := os.ReadFile(n)
				log.Println("Found Character: ", string(file))
				if err != nil {
					log.Println("File Read Error.")
				} else {
					var char Character
					if err := yaml.Unmarshal(file, &char); err != nil {
						log.Fatal(err)
					}
					log.Println(char)
					c = append(c, char)
					log.Println("Character Imported")
				}
			}
		}
	}
	return c
}
