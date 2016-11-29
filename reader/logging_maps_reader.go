package main

import "github.com/gizak/termui"

var pointer = -1
var flagSelector = false
var flags []string
var flagsButtons []*termui.Par
var logList *termui.List

func main() {

	err := termui.Init()

	if err != nil {
		panic(err)
	}

	defer termui.Close()

	flags = []string{
		"INIT",
		"ERROR",
		"404",
		"WARNING",
		"127.0.0.1",
	}
	flagsButtons = make([]*termui.Par, 5)

	//x := 0
	for i, label := range flags {
		flagsButtons[i] = termui.NewPar(label)
		flagsButtons[i].TextFgColor = termui.ColorWhite
		flagsButtons[i].Height = 3
	}

	logs := []string{
		" Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		" In aliquet eu metus et bibendum.",
		" Aliquam viverra convallis libero, ut accumsan leo faucibus id.",
		" Etiam sed est ante.",
		" Duis vel enim nisi.",
		" Pellentesque sed euismod est, nec luctus ipsum.",
		" Praesent commodo odio ac nibh semper vehicula.",
		" Quisque nec congue urna, id congue justo.",
		" Nullam sagittis aliquam mi sed fringilla.",
		" Curabitur ligula sem, lacinia auctor massa tempus, vulputate feugiat metus.",
		" Sed et lobortis eros.",
	}

	logList = termui.NewList()
	logList.Items = logs
	logList.ItemFgColor = termui.ColorWhite
	logList.BorderLabel = "Logs"
	logList.Height = len(logs)

	//termui.Render(l)

	setBody()

	termui.Render(termui.Body)

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		if pointer == -1 {
			flagSelector = !flagSelector

			setBody()
			termui.Render(termui.Body)
		}
	})

	termui.Handle("/sys/kbd/d", func(termui.Event) {
		pointer++

		if pointer > 4 {
			pointer = 4
		}

		setBody()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		pointer--

		if pointer < -1 {
			pointer = -1
		}

		setBody()
		termui.Render(termui.Body)
	})

	termui.Loop()
}

func setBody() {
	termui.Body.Rows = []*termui.Row{}
	if flagSelector {

		for i := range flags {
			if i == pointer {
				flagsButtons[i].BorderFg = termui.ColorYellow
				flagsButtons[i].TextFgColor = termui.ColorYellow
			} else {
				flagsButtons[i].BorderFg = termui.ColorWhite
				flagsButtons[i].TextFgColor = termui.ColorWhite
			}
		}

		back := termui.NewPar("<")
		if pointer == -1 {
			back.BorderFg = termui.ColorYellow
			back.TextFgColor = termui.ColorYellow
		} else {
			back.BorderFg = termui.ColorWhite
			back.TextFgColor = termui.ColorWhite
		}
		back.Height = 3

		termui.Body.AddRows(
			termui.NewRow(
				termui.NewCol(1, 0,
					back,
				),
				termui.NewCol(2, 0,
					flagsButtons[0],
				),
				termui.NewCol(2, 0,
					flagsButtons[1],
				),
				termui.NewCol(2, 0,
					flagsButtons[2],
				),
				termui.NewCol(2, 0,
					flagsButtons[3],
				),
				termui.NewCol(2, 0,
					flagsButtons[4],
				),
			),
			termui.NewRow(
				termui.NewCol(12, 0, logList),
			),
		)
	} else {
		fb := termui.NewPar("Flag Selector")
		fb.Height = 3
		if pointer == -1 {
			fb.BorderFg = termui.ColorYellow
			fb.TextFgColor = termui.ColorYellow
		} else {
			fb.BorderFg = termui.ColorWhite
			fb.TextFgColor = termui.ColorWhite
		}

		termui.Body.AddRows(
			termui.NewRow(
				termui.NewCol(12, 0, fb),
			),
			termui.NewRow(
				termui.NewCol(12, 0, logList),
			),
		)
	}

	termui.Body.Align()
}
