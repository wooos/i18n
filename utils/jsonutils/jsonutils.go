package jsonutils

import (
	"encoding/json"
	"io/ioutil"
)

type Json struct{}

// New return Json type
func New() *Json {
	return &Json{}
}

// Load
func (j *Json) Load(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}