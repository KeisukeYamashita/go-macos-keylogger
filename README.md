# Go macOS Keylogger

> Go written keylogger utils for macOS

_NOTE: This is in development phase, don't use in production_

## Limitations

⚠ Currently, this package is not thread safe. If you are planning in multiple thread, please wait for development.

## Usage

Using this package is super easy and Go native.

```golang
package main

import (
	"fmt"
	"log"

	"github.com/KeisukeYamashita/go-macos-keylogger/pkg/keyboard"
	"github.com/KeisukeYamashita/go-macos-keylogger/pkg/keylogger"
)

func main() {
	kl, err := keylogger.New()
	if err != nil {
		log.Fatal(err)
	}

	f := func(key keyboard.Key, state keyboard.State) {
		fmt.Println(key)
	}

	kl.Listen(f)
}
```

will output

```console
$ go run main.go
```


## Maintainers

* [@KeisukeYamashita](https://github.com/KeisukeYamashita)

## LICENS

Copyright 2021 Go macOS KeyLogger contribuders. This package is released under the Apache License 2.0.
