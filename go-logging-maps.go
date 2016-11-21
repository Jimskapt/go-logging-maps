package loggingmaps

import (
	"bufio"
	"os"
)

// See ./parsers.go for Parser & it instances (JSONParser)
var declaredParser Parser
var output *os.File
var isEmptyFile bool
var Autofields = map[string](func() string){}

// SetParser sets how data will be write (JSON ? XML ? TOML ? ...) on the output, thank to a Parser.
func SetParser(parser Parser) {
	declaredParser = parser
}

// SetOutput sets on which file we write logs.
// It need read rights on this file, because we need to check if it is already empty or not (in order to add a separator).
func SetOutput(filepath string) error {

	// check if file is empty or not.
	// if there is just one \n or \r in the file, it is also judged as empty.
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

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
	output, err = os.OpenFile(filepath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	return nil
}

// LogString log the message and flags.
// It is a simplified form of Log() function.
func LogString(message string, flags ...string) error {

	data := map[string]interface{}{}
	data["message"] = message
	data["flags"] = flags

	addAutoFields(data)

	return Log(data)
}

func generateAutoFields() map[string]string {
	result := map[string]string{}

	for key, function := range Autofields {
		result[key] = function()
	}

	return result
}

func addAutoFields(data map[string]interface{}) {
	for key, value := range generateAutoFields() {
		if data[key] == nil {
			data[key] = value
		}
	}
}

// Log is parsing data with the Parser and write this inside the output.
// Need at least to use SetParser() and SetOutput() before calling this function.
func Log(data map[string]interface{}) error {

	addAutoFields(data)

	bytes, err := declaredParser.Unparse(data)
	if err != nil {
		return err
	}

	closer := ""
	if !isEmptyFile {
		bytes = append([]byte(declaredParser.EntrySeparator()), bytes...)
		closer = declaredParser.RootCloseElement()
	} else {
		bytes = append([]byte(declaredParser.RootOpenElement()), bytes...)
	}
	bytes = append(bytes, []byte(declaredParser.RootCloseElement())...)

	fi, err := output.Stat()
	if err != nil {
		return err
	}

	_, err = output.WriteAt(bytes, fi.Size()-int64(len(([]byte)(closer))))
	if err != nil {
		return err
	}

	if isEmptyFile {
		isEmptyFile = false
	}

	return err
}
