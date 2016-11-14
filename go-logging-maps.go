package loggingmaps

import "bufio"

func SetParser(parser Parser) error {
	return nil
}

func SetOutput(writer *bufio.Writer) error {
	return nil
}

func LogString(message string, flags ...string) error {
	return nil
}

func Log(data map[string]interface{}) error {
	return nil
}

type Parser interface {
	Parse(data []byte) (map[string]interface{}, error)
	Unparse(data map[string]interface{}) ([]byte, error)
}

type JSONParser struct {
	Pretify   bool
	Identchar string
}

func (p JSONParser) Parse(data []byte) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	return result, nil
}
func (p JSONParser) Unparse(data map[string]interface{}) ([]byte, error) {
	result := []byte{}

	return result, nil
}
