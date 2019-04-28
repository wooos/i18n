// Copyright 2019 wooos
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package i18n is for app Internationalization and Localization.

package i18n

import (
	"fmt"
	"github.com/wooos/i18n/utils/json"
)

type Locale struct {
	Lang		string
	Message		[]LocaleJsonMessage
}

type LocaleJson struct {
	Language	string		`json:"language"`
	Messages	[]LocaleJsonMessage	`json:"messages"`
}

type LocaleJsonMessage struct {
	ID			string		`json:"id"`
	Translate	string		`json:"translate"`
}

var (
	localeJson = &LocaleJson{}
)

// SetFileMessage
func SetFileMessage(localeFile string) error {
	_json := json.New()
	if err := _json.LoadFile(localeFile, localeJson); err != nil {
		return err
	}

	return nil
}

// SetStringMessage
func SetStringMessage(localeString string) error {
	_json := json.New()
	if err := _json.LoadString(localeString, localeJson); err != nil {
		return err
	}

	return nil
}

// SetBytesMessage
func SetBytesMessage(localeByte []byte) error {
	_json := json.New()
	if err := _json.LoadBytes(localeByte, localeJson); err != nil {
		return err
	}

	return nil
}

// Tr return string
func Tr(format string, args ...interface{}) string {
	for _, message := range localeJson.Messages {
		if message.ID == format {
			format = message.Translate
		}
	}

	if len(args) > 0 {
		return fmt.Sprintf(format, args...)
	}

	return format
}