package characters

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	s "strings"
)

type Character struct {
	ID   string `yaml:"id"`
	Name string `yaml: "name"`
}

func LoadCharacters() []Character {
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
				if err != nil {
					log.Println("File Read Error.")
				} else {
					var char Character
					if err := yaml.Unmarshal(file, &char); err != nil {
						log.Fatal(err)
					}
					c = append(c, char)
					log.Println("Character Imported: ", char)
				}
			}
		}
	}
	return c
}
