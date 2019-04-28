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

File `locale_en-US.json`:

```json
{
    "language": "en-US",
    "messages": [
        {
            "id": "home",
            "translate": "Home"
        }
    ]
}
```

File `locale_zh-CN.json`:

```json
{
    "language": "zh-CN",
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
i18n.SetFileMessage("conf/locales/locale_zh-CN.json")
il.Tr("home")
```

Code above will produce correspondingly:

- Chinese `zh-CN`：`首页`
