package keylogger

/*
#cgo LDFLAGS: -framework Foundation -framework Foundation -framework Carbon
#include "keylogger.h"

void hello() {
	printf("hello");
}
*/
import "C"

import (
	"errors"
	"os/user"
)

type KeyLogger struct {
}

func New() (*KeyLogger, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	if u.Uid != "0" {
		return nil, errors.New("capturing key logs requires to run the script with root user")
	}

	return &KeyLogger{}, nil
}

func (k *KeyLogger) Listen() {
	C.hello()
}
