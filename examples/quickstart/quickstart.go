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

	autofields := map[string](func() string){
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
	log.SetAutoFields(autofields)

	// comment the following line to append logs in existing file :
	os.Create("./log.json")
	file, _ := os.OpenFile("./log.json", os.O_APPEND|os.O_RDWR, 0600)
	defer file.Close()
	log.SetOutput(file)

	fmt.Println(log.LogString("First start as quickstart.", "INIT", "START", "quickstart"))
	/*
		Returns nil and writes in ./log.json :
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
	*/

	name := "unknown.json"
	data := map[string]interface{}{
		"message":   "The file {{.FileName}} was not found.",
		"file_name": name,
		"flags":     []string{"404", "INIT", name},
	}

	fmt.Println(log.Log(data))
	/*
		Returns nil and writes in ./log.json :
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
			"line": "61",
			"message": "The file {{.FileName}} was not found."
		}
	*/
}
