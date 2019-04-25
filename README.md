i18n
====

Package i18n is for app Internationalization and Localization.

## Introduction

This package provides multiple-language options to improve user experience. 

You can use following command to install this module:

    go get -u github.com/wooos/i18n

## Usage

First of all, you have to import this package:

```go
import "github.com/wooos/i18n"
```

The format of locale files is json format.


## Minimal example

Here are two simplest locale file examples:

File `locale_en.json`:

```json
{
    "language": "en",
    "messages": [
        {
            "id": "home",
            "translate": "Home"
        }
    ]
}
```

File `locale_zh.json`:

```json
{
    "language": "zh",
    "messages": [
        {
            "id": "home",
            "translate": "首页"
        }
    ]
}
```

### Do Translation

Directly use package function to translate:

```go
il := i18n.New("conf/locales")
il.Tr("zh-CN", "home")
```

Code above will produce correspondingly:

- English `en-US`：`Home`
- Chinese `zh-CN`：`首页`
