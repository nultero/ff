package cmd

import (
	"ff/cmd/index"

	"github.com/nultero/tics"
)

var flavorText = tics.Make(" -> f's in the chat").Red().String()

var confFile = "confFile"
var dataDir = "dataDir"

var confMap = map[string]string{
	confFile: "$USER/.config/ff.yaml",
	dataDir:  "$USER/.ff",
}

var defaultSettings = []string{
	"editor: nvim",
	"notes path: $USER/.ff",
	"max columns: 4",
	"column padding: 2",
	"notes padding: 3",
}

var FileFlag string
var CategoryFlag bool
var NoteFlag bool

var dex = index.Log{Changed: false}
