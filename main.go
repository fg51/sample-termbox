package main

import "fmt"
import "time"
import (
	"github.com/nsf/termbox-go"
)


type Message struct {
	x int
	y int
	format string
}


func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	inputmode := 0
	ctrlxpressed := false

	msg0 := "hello"
	x, y := 1, 2
	printTermbox(x, y, termbox.ColorDefault, termbox.ColorDefault, msg0)


	msgStep := newMessage(1, 3, "step : %04d")
	printTermbox(
		msgStep.x,
		msgStep.y,
		termbox.ColorDefault,
		termbox.ColorDefault,
		fmt.Sprintf(msgStep.format, 1),
	)

	now := time.Now()
	x, y = 1, 4
	formatTime := "time : %d/%02d/%02d %02d:%02d"
	msgTime := fmt.Sprintf(
		formatTime,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute())
	printTermbox(x, y, termbox.ColorDefault, termbox.ColorDefault, msgTime)


	msgSpeed := newMessage(1, 5, "speed: %4d [rpm] [#### ###   ]")
	printTermbox(
		msgSpeed.x,
		msgSpeed.y,
		termbox.ColorDefault,
		termbox.ColorDefault,
		fmt.Sprintf(msgSpeed.format, 100),
	)

	msgPower := newMessage(1, 6, "power: %4d [A]   [#### ##    ]")
	printTermbox(
		msgPower.x,
		msgPower.y,
		termbox.ColorDefault,
		termbox.ColorDefault,
		fmt.Sprintf(msgPower.format, 3))

	printTermbox(1, 7, termbox.ColorDefault, termbox.ColorDefault, "step1 - step1.5 - step2 - step3 -step4")

	printTermbox(9, 7, termbox.ColorDefault, termbox.ColorMagenta, "step1.5")

	printTermbox(1, 8, termbox.ColorDefault, termbox.ColorDefault, "status: accel")


	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlC {
				break loop
			}
			if ev.Key == termbox.KeyCtrlS && ctrlxpressed {
				termbox.Sync()
			}
			if ev.Key == termbox.KeyCtrlQ && ctrlxpressed {
				break loop
			}
			if ev.Key == termbox.KeyCtrlC && ctrlxpressed {
				chmap := []termbox.InputMode{
					termbox.InputEsc | termbox.InputMouse,
					termbox.InputAlt | termbox.InputMouse,
					termbox.InputEsc,
					termbox.InputAlt,
				}
				inputmode++
				if inputmode >= len(chmap) {
					inputmode = 0
				}
				termbox.SetInputMode(chmap[inputmode])
			}
			if ev.Key == termbox.KeyCtrlX {
				ctrlxpressed = true
			} else {
				ctrlxpressed = false
			}

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			// draw_keyboard()
			// dispatch_press(&ev)
			// pretty_print_press(&ev)
			// termbox.Flush()
		// case termbox.EventResize:
		// 	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		// 	draw_keyboard()
		// 	pretty_print_resize(&ev)
		// 	termbox.Flush()
		// case termbox.EventMouse:
		// 	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		// 	draw_keyboard()
		// 	pretty_print_mouse(&ev)
		// 	termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}

}

func printTermbox(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func newMessage(x int, y int, format string) *Message{
	m := new(Message)
	m.x = x
	m.y = y
	m.format = format
	return m
}

