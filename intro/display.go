package intro

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	_ "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	_ "image/color"
	"log"
	"time"
)

const (
	Title       = "Xiaohan's Introduction"
	TextDelay   = 5
	WindowEdgeH = 400
	WindowScale = 16 / 9.0
	Start       = "Start"
	End         = "End"
)

var (
	WindowDefaultSize = fyne.Size{Width: WindowEdgeH * WindowScale, Height: WindowEdgeH}
	quitEvent         = make(chan bool, 1)
	// define button color
	buttonColor = canvas.NewRectangle(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
)

func Exec(a fyne.App) {
	// read introduction
	text, _ := Read()
	// define window
	w := a.NewWindow(Title)
	// get canvas
	//c := w.Canvas()
	// define content button
	content := widget.NewButton(text[0], WarnMsg())
	// define start button
	startButton := widget.NewButton(Start, Display(a, w, content, text))
	// define container
	//c := container.New(layout.NewStackLayout(), buttonColor, startButton)
	// set container
	w.SetContent(startButton)
	// resize window
	w.Resize(WindowDefaultSize)
	w.CenterOnScreen()
	w.ShowAndRun()
}

func WarnMsg() func() {
	return func() {
		log.Println("wait display completed")
	}
}

func Display(a fyne.App, w fyne.Window, button *widget.Button, text []string) func() {
	return func() {
		log.Println("start to display content")
		w.SetContent(button)
		go delay(a, w, button, text)
	}
}

func delay(a fyne.App, w fyne.Window, content *widget.Button, text []string) {
	for i := 1; i < len(text); i++ {
		content.SetText(text[i])
		time.Sleep(time.Second * TextDelay)
	}
	//// send quit event
	//quitEvent <- true
	//close(quitEvent)
	log.Println("display completed")
	quitButton := widget.NewButton(text[len(text)-1], a.Quit)
	w.SetContent(quitButton)
}

func triggerQuit(a fyne.App, w fyne.Window, text []string) func() {
	return func() {
		go func() {
			flag, ok := <-quitEvent
			if flag || !ok {
				quitButton := widget.NewButton(text[len(text)-1], a.Quit)
				w.SetContent(quitButton)
			}
		}()
	}
}
