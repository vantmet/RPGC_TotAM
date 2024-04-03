package main

import (
	"fmt"
	"github.com/vantmet/RPGC_TotAM/internal/menu"
)

func main() {
	menu := menu.NewMenu("Chose a colour")

	menu.AddItem("Red", "red")
	menu.AddItem("Blue", "blue")
	menu.AddItem("Green", "green")
	menu.AddItem("Yellow", "yellow")
	menu.AddItem("Cyan", "cyan")

	choice := menu.Display()

	fmt.Printf("Choice: %s\n", choice)
}
