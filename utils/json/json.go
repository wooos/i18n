package json

import (
	"encoding/json"
	"io/ioutil"
)

type Json struct{}

// New return Json type
func New() *Json {
	return &Json{}
}

// LoadFile
func (j *Json) LoadFile(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}

// LoadString
func (j *Json) LoadString(data string, v interface{}) error {
	if err := json.Unmarshal([]byte(data), v); err != nil {
		return err
	}

	return nil
}

// LoadBytes
func (j *Json) LoadBytes(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}