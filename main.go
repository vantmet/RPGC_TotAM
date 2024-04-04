package main

import (
	"fmt"
	c "github.com/vantmet/RPGC_TotAM/internal/characters"
	"github.com/vantmet/RPGC_TotAM/internal/menu"
)

func main() {
	characters := c.LoadCharacters()

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
