package main

import (
	"fmt"
	"github.com/vantmet/RPGC_TotAM/internal/menu"
)

func main() {
	menu := menu.NewMenu("Chose an option:")

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
