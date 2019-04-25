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
	"github.com/wooos/i18n/utils/jsonutils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	_locale Locale

	Matcher = language.NewMatcher([]language.Tag{
		language.Chinese,
		language.English,
	})
)

type IL struct {
	dirPath     string
	languageTag language.Tag
}

type Locale struct {
	Language string          `json:"language"`
	Messages []LocaleMessage `json:"messages"`
}

type LocaleMessage struct {
	Key     string `json:"id"`
	Message string `json:"translate"`
}

func New(dirPath string) *IL {
	return &IL{dirPath: dirPath}
}

func _tag(lang string) language.Tag {
	t, _ := language.MatchStrings(Matcher, lang)
	return t
}

func (il *IL) setLanguageTag(tag language.Tag) error {
	il.languageTag = tag
	filename := il.dirPath + "/locale_" + il.languageTag.Parent().String() + ".json"
	ju := jsonutils.New()
	if err := ju.Load(filename, &_locale); err != nil {
		return err
	}
	return nil
}

func (il *IL) load(tag language.Tag) error {
	if err := il.setLanguageTag(tag); err != nil {
		return err
	}

	for _, v := range _locale.Messages {
		message.SetString(tag, v.Key, v.Message)
	}

	return nil
}

func (il *IL) Translate(lang string, key string) string {
	tag := _tag(lang)
	if err := il.load(tag); err != nil {
		return ""
	}

	p := message.NewPrinter(tag)
	return p.Sprintf(key)
}
