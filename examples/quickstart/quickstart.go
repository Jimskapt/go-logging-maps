package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	log "github.com/Jimskapt/go-logging-maps"
)

func main() {

	log.SetParser(log.JSONParser{Pretify: true, IdentChar: "\t"})

	log.Autofields = map[string](func() string){
		"file": func() string {
			_, file, _, _ := runtime.Caller(4)
			return file
		},
		"line": func() string {
			_, _, line, _ := runtime.Caller(4)
			return strconv.Itoa(line)
		},
		"date": func() string {
			return time.Now().UTC().Format(time.RFC3339)
		},
	}

	// comment the following line to append logs in the file :
	os.Create("./log.json")

	fmt.Println(log.SetOutput("./log.json"))

	fmt.Println("//////////////////////////////")

	fmt.Println(log.LogString("First start as quickstart.", "INIT", "START", "quickstart"))
	/*
		Returns <nil> and writes in ./log.json :
		[
			{
				"date": "2016-11-17T20:26:42Z",
				"file": "[your GOPATH]/src/github.com/Jimskapt/go-logging-maps/examples/quickstart/quickstart.go",
				"flags": [
					"INIT",
					"START",
					"quickstart"
				],
				"line": "38",
				"message": "First start as quickstart."
			}
		]
	*/

	fmt.Println("//////////////////////////////")

	name := "unknown.json"
	data := map[string]interface{}{
		"message":   "The file {{.FileName}} was not found.",
		"file_name": name,
		"flags":     []string{"404", "INIT", name},
	}

	fmt.Println(log.Log(data))

	/*
		Returns <nil> and writes in ./log.json :
			,
			{
				"date": "2016-11-17T20:26:42Z",
				"file": "[your GOPATH]/src/github.com/Jimskapt/go-logging-maps/examples/quickstart/quickstart.go",
				"file_name": "unknown.json",
				"flags": [
					"404",
					"INIT",
					"unknown.json"
				],
				"line": "65",
				"message": "The file {{.FileName}} was not found."
			}
		]
		(this data overwrites the "\n]", which is already exists in log.json - from previous example -)
	*/
}
