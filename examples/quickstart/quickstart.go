package main

import (
	"fmt"
	"os"

	log "github.com/Jimskapt/go-logging-maps"
)

func main() {

	log.SetParser(log.JSONParser{Pretify: true, IdentChar: "\t"})

	// comment the following line to append logs in existing file :
	os.Create("./log.json")
	file, _ := os.OpenFile("./log.json", os.O_APPEND|os.O_RDWR, 0600)
	defer file.Close()
	log.SetOutput(file)

	fmt.Println(log.LogString("First start as quickstart.", "INIT", "START", "quickstart"))
	/*
		Returns nil and writes in ./log.json :
		{
			"flags": [
				"INIT",
				"START",
				"quickstart"
			],
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
			"file_name": "unknown.json",
			"flags": [
				"404",
				"INIT",
				"unknown.json"
			],
			"message": "The file {{.FileName}} was not found."
		}
	*/

}
