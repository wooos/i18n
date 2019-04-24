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
	"github.com/wooos/i18n/utils/jsonutils"
)

var (
	Chinese = "zh"
	English = "en"

	_locale Locale
)

type IL struct {
	defaultLocale string
	dirPath       string
	locale        string
}

type Locale struct {
	Language string          `json:"language"`
	Messages []LocaleMessage `json:"messages"`
}

type LocaleMessage struct {
	Key       string `json:"ID"`
	Translate string `json:"translate"`
}

func NewLocale(dirPath, lang string) *IL {
	defaultLocale := dirPath + "/locale." + lang + ".json"
	j := jsonutils.New()
	if err := j.Load(defaultLocale, &_locale); err != nil {
		fmt.Printf("%v\n", err)
	}
	il := &IL{defaultLocale: defaultLocale, dirPath: dirPath, locale: lang}
	return il
}

func (il *IL) setLang(lang string) {
	il.locale = lang
	defaultLocale := il.dirPath + "/locale." + il.locale + ".json"
	j := jsonutils.New()
	j.Load(defaultLocale, &_locale)
	fmt.Printf("%v\n", defaultLocale)
	il.defaultLocale = defaultLocale
}

func (il *IL) Translate(key string, lang string) string {
	il.setLang(lang)
	for _, v := range _locale.Messages {
		if v.Key == key {
			return v.Translate
		}
	}
	return ""
}
