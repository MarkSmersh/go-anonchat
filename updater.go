package main

import (
	"github.com/MarkSmersh/go-anonchat/types/general"
)

type Updater struct {
	messages Caller[general.Message]
	commands Caller[general.Message]
}
