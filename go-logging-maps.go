package loggingmaps

import (
	"bufio"
	"os"
)

// See ./parsers.go for Parser & it instances (JSONParser)
var declaredParser Parser
var output *bufio.Writer
var isEmptyFile bool

// SetParser sets how data will be write (JSON ? XML ? TOML ? ...) on the output, thank to a Parser.
func SetParser(parser Parser) {
	declaredParser = parser
}

// SetOutput sets on which file we write logs.
// It need read rights on this file, because we need to check if it is already empty or not (in order to add a separator).
func SetOutput(file *os.File) {

	// check if file is empty or not.
	// if there is just one \n or \r in the file, it is also judged as empty.
	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanRunes)
	notEmpty := scanner.Scan()
	if notEmpty {
		if (scanner.Text() == "\n" || scanner.Text() == "\r") && scanner.Scan() == false {
			isEmptyFile = true
		} else {
			isEmptyFile = false
		}
	} else {
		isEmptyFile = true
	}

	// only save a writer to file
	output = bufio.NewWriter(file)
}

// LogString log the message and flags.
// It is a simplified form of Log() function.
func LogString(message string, flags ...string) error {

	data := map[string]interface{}{}
	data["message"] = message
	data["flags"] = flags

	return Log(data)
}

// Log is parsing data with the Parser and write this inside the output.
// Need at least to use SetParser() and SetOutput() before calling this function.
func Log(data map[string]interface{}) error {

	bytes, err := declaredParser.Unparse(data)
	if err != nil {
		return err
	}

	if !isEmptyFile {
		_, err = output.Write([]byte(declaredParser.EntrySeparator()))
		if err != nil {
			return err
		}
	}

	_, err = output.Write(bytes)
	if err != nil {
		return err
	}

	err = output.Flush()
	if err != nil {
		return err
	}

	if isEmptyFile {
		isEmptyFile = false
	}

	return err
}
