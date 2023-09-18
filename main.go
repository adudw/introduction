package main

import (
	"fyne.io/fyne/v2/app"
	"introduction/intro"
)

func main() {
	a := app.New()
	intro.Exec(a)
}
