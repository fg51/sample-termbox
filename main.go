package main


import (
	"github.com/nsf/termbox-go"
)


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

	termbox.SetCell(1, 2, rune('a'), termbox.ColorDefault, termbox.ColorDefault)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
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

