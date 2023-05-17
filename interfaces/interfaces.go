package interfaces

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log() // => makes a shareholder for logging messages from a different variables but has the same function name
}

type loggableString string
type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func (text loggableString) log() {
	fmt.Println(text)
}

func createLog(data logger) {
	data.log()
}

// expecting any value using the concept of empty interfaces
func outputValue(value interface{}) {
	val, ok := value.(string) // => type assertion <checking that the value of provided the parameter is a string or other any checking type

	fmt.Println(val, ok)
	fmt.Println(value)
}

// no specific return type
func double(value interface{}) interface{} {
	switch val := value.(type) {
	case int:
		return val * 2
	case float64:
		return val * 2
	default:
		return value
	}
}
