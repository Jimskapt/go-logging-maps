package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gizak/termui"
)

var pointer = -1
var flagSelector = false
var flags []Flag
var logList *termui.List

type Flag struct {
	Name      string
	Activated bool
	Button    *termui.Par
}

type Log struct {
	Flags []string `json:"flags"`
}

func main() {

	err := termui.Init()

	if err != nil {
		panic(err)
	}

	defer termui.Close()

	fileData, err := ioutil.ReadFile("./log.json")
	if err != nil {
		log.Fatal(err)
	}

	var parsedLogs []Log
	err = json.Unmarshal(fileData, &parsedLogs)
	if err != nil {
		panic(err)
	}

	logs := []string{}
	scanner := bufio.NewScanner(bytes.NewReader(fileData))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	flagsLabels := []string{
		"INIT",
		"ERROR",
		"404",
		"WARNING",
		"127.0.0.1",
	}
	flags = make([]Flag, len(flagsLabels))

	for i, flag := range flagsLabels {
		flags[i] = Flag{
			Name:      flag,
			Activated: true,
			Button:    termui.NewPar(flag),
		}
		flags[i].Button.TextFgColor = termui.ColorWhite
		flags[i].Button.Height = 3
		flags[i].Button.Width = len(flag)
	}

	/*logs := []string{
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
	}*/

	logList = termui.NewList()
	logList.Items = logs
	logList.ItemFgColor = termui.ColorWhite
	logList.BorderLabel = "Logs"
	logList.Height = len(logs)

	setBody()

	termui.Render(termui.Body)

	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		if pointer == -1 {
			flagSelector = !flagSelector

			setBody()
			termui.Render(termui.Body)
		} else {
			flags[pointer].Activated = !flags[pointer].Activated
			setBody()
			termui.Render(termui.Body)
		}
	})

	termui.Handle("/sys/kbd/C-c", escape)
	termui.Handle("/sys/kbd/<escape>", escape)

	termui.Handle("/sys/kbd/<right>", rightButton)
	termui.Handle("/sys/kbd/d", rightButton)

	termui.Handle("/sys/kbd/<left>", leftButton)
	termui.Handle("/sys/kbd/q", leftButton)

	termui.Loop()
}

func rightButton(e termui.Event) {
	if flagSelector {
		pointer++

		if pointer > len(flags)-1 {
			pointer = len(flags) - 1
		}

		setBody()
		termui.Render(termui.Body)
	}
}

func leftButton(e termui.Event) {
	if flagSelector {
		pointer--

		if pointer < -1 {
			pointer = -1
		}

		setBody()
		termui.Render(termui.Body)
	}
}

func escape(e termui.Event) {
	if e.Path == "/sys/kbd/<escape>" {
		if flagSelector {
			flagSelector = false
			pointer = -1

			setBody()
			termui.Render(termui.Body)
		} else {
			termui.StopLoop()
		}
	} else {
		termui.StopLoop()
	}
}

func setBody() {
	termui.Body.Rows = []*termui.Row{}
	if flagSelector {

		for i := range flags {
			if i == pointer {
				flags[i].Button.BorderFg = termui.ColorYellow

				if flags[i].Activated {
					flags[i].Button.TextBgColor = termui.ColorYellow
					flags[i].Button.TextFgColor = termui.ColorBlack
				} else {
					flags[i].Button.TextBgColor = termui.ColorBlack
					flags[i].Button.TextFgColor = termui.ColorYellow
				}
			} else {
				flags[i].Button.BorderFg = termui.ColorWhite

				if flags[i].Activated {
					flags[i].Button.TextBgColor = termui.ColorWhite
					flags[i].Button.TextFgColor = termui.ColorBlack
				} else {
					flags[i].Button.TextBgColor = termui.ColorBlack
					flags[i].Button.TextFgColor = termui.ColorWhite
				}
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

		max := len(flags)
		if len(flags) > 11 {
			max = 11
		}
		cols := make([]*termui.Row, max+1)
		cols[0] = termui.NewCol(1, 0, back)
		for i, flag := range flags {
			cols[i+1] = termui.NewCol(1, 0, flag.Button)
			if i+1 >= 11 {
				break
			}
		}

		termui.Body.AddRows(
			termui.NewRow(cols...),
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
