package main

import (
	"bufio"
	"os"

	log "github.com/Jimskapt/go-logging-maps"
)

func main() {

	log.SetParser(log.JSONParser{Pretify: true, Identchar: "\t"})

	file, _ := os.Create("./log.json")
	logfile := bufio.NewWriter(file)
	defer logfile.Flush()
	log.SetOutput(logfile)

	log.LogString("First start as quickstart.", "INIT", "START", "quickstart")

	name := "unknown.json"
	data := map[string]interface{}{}
	data["Message"] = "The file {{.FileName}} was not found."
	data["FileName"] = name
	data["Flags"] = []string{"404", "INIT", name}

	log.Log(data)

}
