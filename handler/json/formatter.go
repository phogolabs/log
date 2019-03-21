package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// PrintfColorFunc prints the value in color
type PrintfColorFunc = func(format string, a ...interface{}) string

// PrintColorFunc prints the text in color
type PrintColorFunc = func(a ...interface{}) string

// ColorFunc returns the color formatter
type ColorFunc func(k, v interface{}) ColorFormatter

// ColorFormatter formats a color
type ColorFormatter interface {
	SprintfFunc() PrintfColorFunc
	SprintFunc() PrintColorFunc
}

// Formatter is a struct to format JSON data. `color` is github.com/fatih/color: https://github.com/fatih/color
type Formatter struct {
	// JSON key color. Default is `color.New(color.FgBlue, color.Bold)`.
	KeyColor ColorFormatter

	// JSON string value color. Default is `color.New(color.FgGreen, color.Bold)`.
	StringColor ColorFormatter

	// JSON boolean value color. Default is `color.New(color.FgYellow, color.Bold)`.
	BoolColor ColorFormatter

	// JSON number value color. Default is `color.New(color.FgCyan, color.Bold)`.
	NumberColor ColorFormatter

	// JSON null value color. Default is `color.New(color.FgBlack, color.Bold)`.
	NullColor ColorFormatter

	// ColorFn returns the color for given key and value
	ColorFn ColorFunc

	// Max length of JSON string value. When the value is 1 and over, string is truncated to length of the value.
	// Default is 0 (not truncated).
	StringMaxLength int

	// DisablePretty to disable formatting
	DisablePretty bool

	// Boolean to disable color. Default is false.
	DisabledColor bool

	// Indent space number. Default is 2.
	Indent int

	// Newline string. To print without new lines set it to empty string. Default is \n.
	Newline string
}

// NewFormatter returns a new formatter with following default values.
func NewFormatter() *Formatter {
	return &Formatter{
		KeyColor:        color.New(color.FgBlue, color.Bold),
		StringColor:     color.New(color.FgGreen, color.Bold),
		BoolColor:       color.New(color.FgYellow, color.Bold),
		NumberColor:     color.New(color.FgCyan, color.Bold),
		NullColor:       color.New(color.FgBlack, color.Bold),
		StringMaxLength: 0,
		DisabledColor:   true,
		DisablePretty:   true,
		Indent:          2,
		Newline:         "\n",
	}
}

// Marshal marshals and formats JSON data.
func (f *Formatter) Marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	if !f.DisablePretty {
		data, err = f.format(data)
	}

	return data, err
}

// Format formats JSON string.
func (f *Formatter) format(data []byte) ([]byte, error) {
	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	t := reflect.TypeOf(v)

	return []byte(f.pretty(t.Name(), v, 1)), nil
}

func (f *Formatter) sprintfColor(c ColorFormatter, format string, args ...interface{}) string {
	if f.DisabledColor || c == nil {
		return fmt.Sprintf(format, args...)
	}
	return c.SprintfFunc()(format, args...)
}

func (f *Formatter) sprintColor(c ColorFormatter, s string) string {
	if f.DisabledColor || c == nil {
		return fmt.Sprint(s)
	}
	return c.SprintFunc()(s)
}

func (f *Formatter) color(k interface{}, v interface{}) ColorFormatter {
	if colorFn := f.ColorFn; colorFn != nil {
		if color := colorFn(k, v); color != nil {
			return color
		}
	}

	switch v.(type) {
	case string:
		return f.StringColor
	case float64:
		return f.NumberColor
	case bool:
		return f.BoolColor
	case nil:
		return f.NullColor
	default:
		return nil
	}
}

func (f *Formatter) pretty(k, v interface{}, depth int) string {
	color := f.color(k, v)

	switch value := v.(type) {
	case string:
		return f.processString(color, value)
	case float64:
		return f.sprintColor(color, strconv.FormatFloat(value, 'f', -1, 64))
	case bool:
		return f.sprintColor(color, strconv.FormatBool(value))
	case nil:
		return f.sprintColor(color, "null")
	case map[string]interface{}:
		return f.processMap(k, value, depth)
	case []interface{}:
		return f.processArray(k, value, depth)
	}

	return ""
}

func (f *Formatter) processString(color ColorFormatter, value string) string {
	runes := []rune(value)

	if f.StringMaxLength != 0 && len(runes) >= f.StringMaxLength {
		value = string(runes[0:f.StringMaxLength]) + "..."
	}

	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(value)

	value = string(buffer.Bytes())
	value = strings.TrimSuffix(value, "\n")

	return f.sprintColor(color, value)
}

func (f *Formatter) processMap(k interface{}, m map[string]interface{}, depth int) string {
	if len(m) == 0 {
		return "{}"
	}

	currentIndent := f.generateIndent(depth - 1)
	nextIndent := f.generateIndent(depth)
	rows := []string{}
	keys := []string{}

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := m[key]

		k := f.sprintfColor(f.KeyColor, `"%s"`, key)
		v := f.pretty(key, value, depth+1)

		valueIndent := " "

		if f.Newline == "" {
			valueIndent = ""
		}

		row := fmt.Sprintf("%s%s:%s%s", nextIndent, k, valueIndent, v)
		rows = append(rows, row)
	}

	return fmt.Sprintf("{%s%s%s%s}", f.Newline, strings.Join(rows, ","+f.Newline), f.Newline, currentIndent)
}

func (f *Formatter) processArray(key interface{}, arr []interface{}, depth int) string {
	if len(arr) == 0 {
		return "[]"
	}

	currentIndent := f.generateIndent(depth - 1)
	nextIndent := f.generateIndent(depth)
	rows := []string{}

	for _, value := range arr {
		item := f.pretty(key, value, depth+1)

		row := nextIndent + item
		rows = append(rows, row)
	}

	return fmt.Sprintf("[%s%s%s%s]", f.Newline, strings.Join(rows, ","+f.Newline), f.Newline, currentIndent)
}

func (f *Formatter) generateIndent(depth int) string {
	return strings.Repeat(" ", f.Indent*depth)
}
