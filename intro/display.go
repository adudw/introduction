package intro

import (
	"fyne.io/fyne/v2"
	_ "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "image/color"
	"log"
	"time"
)

const (
	Title = "Introduction"
	EdgeH = "     "
	EdgeV = "   "
	Delay = 1
	Count = 1
)

var (
	quitChan = make(chan bool, 1)
)

func Exec(a fyne.App) {

	text, _ := Read()

	// define window
	w := a.NewWindow(Title)

	// define container
	//V := canvas.NewText(EdgeV, color.White)
	//H := canvas.NewText(EdgeH, color.White)
	//c := container.New(layout.NewBorderLayout(V, V, H, H))
	c := container.New(layout.NewCenterLayout())

	// define start button
	startButton := widget.NewButton("Start", Display(a, c, widget.NewLabel(text[0]), text))
	c.Add(startButton)

	w.SetContent(c)
	w.CenterOnScreen()
	w.ShowAndRun()
}

func Display(a fyne.App, c *fyne.Container, label *widget.Label, text []string) func() {
	return func() {
		log.Println("display content")
		c.RemoveAll()
		c.Add(label)
		go delay(text, label, c, widget.NewButton("Quit", QuitClick(a)))
	}
}

func QuitClick(a fyne.App) func() {
	return func() {
		//log.Printf("chan %v", quitChan)
		go func() {
			flag, ok := <-quitChan
			log.Println(flag, ok)
			if !ok {
				log.Println("wait completing, can't quit.")
			}
			if flag {
				a.Quit()
			}
		}()
	}
}

func delay(text []string, content *widget.Label, c *fyne.Container, object fyne.Widget) {
	for i := 1; i < len(text); i++ {
		content.SetText(text[i])
		time.Sleep(time.Second * Delay)
	}
	// 传递可关闭信号
	quitChan <- true
	// 更新 object
	//wg.Done()
	c.RemoveAll()
	c.Add(object)
}
