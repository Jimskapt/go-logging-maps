package loggingmaps

import "encoding/json"

// Parser sets how data will be saved (JSON ? XML ? TOML ? ...) on the output.
type Parser interface {
	Parse(data []byte) (map[string]interface{}, error)
	Unparse(data map[string]interface{}) ([]byte, error)
	EntrySeparator() string
	RootOpenElement() string
	RootCloseElement() string
}

// JSONParser is the JSON Parser
type JSONParser struct {
	Pretify   bool
	IdentChar string
}

func (p JSONParser) Parse(data []byte) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	err := json.Unmarshal(data, result)

	return result, err
}

func (p JSONParser) Unparse(data map[string]interface{}) ([]byte, error) {
	if p.Pretify {
		return json.MarshalIndent(data, "\t", p.IdentChar)
	}

	return json.Marshal(data)
}

func (p JSONParser) EntrySeparator() string {
	if p.Pretify {
		return ",\n\t"
	}

	return ","
}

func (p JSONParser) RootOpenElement() string {
	if p.Pretify {
		return "[\n\t"
	}

	return "["
}

func (p JSONParser) RootCloseElement() string {
	if p.Pretify {
		return "\n]"
	}

	return "]"
}
