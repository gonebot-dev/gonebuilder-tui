package botcreator

import (
	"fmt"
	"os"
	"sync"

	"github.com/gonebot-dev/gonebuilder-tui/app/base"
)

type logString struct {
	level string
	text  string
}

var mutex sync.RWMutex
var progress = 0
var createText []logString

const (
	INF = "INF"
	SUC = "SUC"
	WRN = "WRN"
	ERR = "ERR"
)

func addMsg(level, text string) {
	mutex.Lock()
	createText = append(createText, logString{level: level, text: text})
	if level == SUC {
		progress += 1
	}
	mutex.Unlock()
}

func GetInfo() (prog int, text []logString) {
	mutex.RLock()
	prog = progress
	copy(text, createText)
	mutex.RUnlock()
	return
}

func CreateBot(folderPath string) {
	var err error
	mutex.Lock()
	progress = 0
	createText = make([]logString, 0)
	createText = append(createText, logString{level: "INF", text: "Creating/Replacing folder..."})
	mutex.Unlock()
	if _, err = os.Stat(base.BotFolder); err == nil || !os.IsNotExist(err) {
		addMsg(INF, "Removing existing folder...")
		os.RemoveAll(base.BotFolder)
		addMsg(INF, "Removed successfully!")
	}
	addMsg(INF, "Creating bot folder...")
	if err = os.Mkdir(base.BotFolder, os.ModePerm); err != nil {
		addMsg(ERR, fmt.Sprintf("Error creating folder: %s", err.Error()))
		return
	}
	addMsg(SUC, "Created successfully!")
}
