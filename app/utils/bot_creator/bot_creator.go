package botcreator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gonebot-dev/gonebuilder-tui/app/utils/api"
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

func formatName(name string) string {
	builder := strings.Builder{}
	for idx, char := range name {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (idx > 0 && (char >= '0' && char <= '9')) {
			builder.WriteRune(char)
		} else {
			builder.WriteRune('_')
		}
	}
	return builder.String()
}

func writeGoFile(path, name, version, desc string, adapters *[]api.AdapterInfo, plugins *[]api.PluginInfo) (err error) {
	adaptersImports := strings.Builder{}
	pluginsImports := strings.Builder{}
	adaptersLoads := strings.Builder{}
	pluginsLoads := strings.Builder{}
	filePath := filepath.Join(path, name+".go")
	fileStr := strings.Builder{}

	for _, adapter := range *adapters {
		adaptersImports.WriteString(fmt.Sprintf("\t%s \"%s\"\n", adapter.Name, adapter.Package))
		adaptersLoads.WriteString(fmt.Sprintf("\tgonebot.LoadAdapter(&%s)\n", adapter.Adapter))
	}
	for _, plugin := range *plugins {
		pluginsImports.WriteString(fmt.Sprintf("\t%s \"%s\"\n", plugin.Name, plugin.Package))
		pluginsLoads.WriteString(fmt.Sprintf("\tgonebot.LoadPlugin(&%s)\n", plugin.Plugin))
	}

	fileStr.WriteString("package main\n\n")
	fileStr.WriteString("import (\n")
	fileStr.WriteString("\t\"fmt\"\n")
	fileStr.WriteString("\t\"github.com/gonebot-dev/gonebot\"\n")
	fileStr.WriteString(pluginsImports.String())
	fileStr.WriteString(adaptersImports.String())
	fileStr.WriteString(")\n\n")
	fileStr.WriteString("func main() {\n")
	fileStr.WriteString(fmt.Sprintf("\tBotVersion := \"%s\"\n", version))
	fileStr.WriteString(fmt.Sprintf("\tBotDescription := \"%s\"\n", desc))
	fileStr.WriteString("\tfmt.Println(\"\\nBotVersion: \", BotVersion)\n")
	fileStr.WriteString("\tfmt.Println(\"BotDescription: \", BotDescription, \"\\n\")\n")
	fileStr.WriteString(pluginsLoads.String())
	fileStr.WriteString(adaptersLoads.String())
	fileStr.WriteString("\tgonebot.Run()\n")
	fileStr.WriteString("}\n")

	err = os.WriteFile(filePath, []byte(fileStr.String()), 0644)
	return
}

func CreateBot(folderPath, botName, botVersion, botDesc string, adapters *[]api.AdapterInfo, plugins *[]api.PluginInfo) (err error) {
	mutex.Lock()
	progress = 0
	createText = make([]logString, 0)
	createText = append(createText, logString{level: "INF", text: "Creating/Replacing folder..."})
	mutex.Unlock()
	botFolder := strings.ReplaceAll(filepath.Join(folderPath, botName), "\\", "/")
	if _, err = os.Stat(botFolder); err == nil || !os.IsNotExist(err) {
		addMsg(INF, "Removing existing folder...")
		if err = os.RemoveAll(botFolder); err != nil {
			addMsg(ERR, fmt.Sprintf("Error removing folder: %s", err.Error()))
			return
		}
		addMsg(INF, "Removed successfully!")
	}
	addMsg(INF, "Creating bot folder...")
	if err = os.Mkdir(botFolder, os.ModePerm); err != nil {
		addMsg(ERR, fmt.Sprintf("Error creating folder: %s", err.Error()))
		return
	}
	addMsg(SUC, "Created folder successfully!")
	pwd, _ := os.Getwd()
	os.Chdir(botFolder)
	moduleName := formatName(botName)
	addMsg(INF, fmt.Sprintf("Creating go module %s...", moduleName))
	if err = exec.Command("go", "mod", "init", moduleName).Run(); err != nil {
		addMsg(ERR, fmt.Sprintf("Error creating go module: exited with %s", err.Error()))
		return
	}
	addMsg(SUC, fmt.Sprintf("Created go module %s successfully!", moduleName))
	addMsg(INF, fmt.Sprintf("Writing %s...", moduleName))
	if err = writeGoFile(botFolder, moduleName, botVersion, botDesc, adapters, plugins); err != nil {
		addMsg(ERR, fmt.Sprintf("Error writing go file: exited with %s", err.Error()))
		return
	}
	addMsg(SUC, fmt.Sprintf("Wrote %s.go successfully!", moduleName))
	addMsg(INF, "Formatting go file...")
	if err = exec.Command("go", "fmt").Run(); err != nil {
		addMsg(ERR, fmt.Sprintf("Error formatting go file: exited with %s", err.Error()))
		return
	}
	addMsg(SUC, "Formatted go file successfully!")
	addMsg(INF, "Running go mod tidy...")
	if err = exec.Command("go", "mod", "tidy").Run(); err != nil {
		addMsg(ERR, fmt.Sprintf("Error running go mod tidy: exited with %s", err.Error()))
		return
	}
	addMsg(SUC, "Ran go mod tidy successfully!")
	addMsg(SUC, "Bot created successfully!")
	os.Chdir(pwd)
	return
}
