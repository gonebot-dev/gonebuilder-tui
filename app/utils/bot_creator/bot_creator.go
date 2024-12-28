package botcreator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/gonebot-dev/gonebuilder-tui/app/utils/api"
	"github.com/rs/zerolog"
)

var mutex sync.RWMutex
var progress = 0

func FormatName(name string) string {
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

func writeGoFile(path, name, version, desc string, adapters *[]list.Item, plugins *[]list.Item) (err error) {
	adaptersImports := strings.Builder{}
	pluginsImports := strings.Builder{}
	adaptersLoads := strings.Builder{}
	pluginsLoads := strings.Builder{}
	filePath := filepath.Join(path, name+".go")
	fileStr := strings.Builder{}

	for _, adap := range *adapters {
		adapter := adap.(api.AdapterInfo)
		adaptersImports.WriteString(fmt.Sprintf("\t%s \"%s\"\n", adapter.Name, adapter.Package))
		adaptersLoads.WriteString(fmt.Sprintf("\tgonebot.LoadAdapter(&%s)\n", adapter.Adapter))
	}
	for _, plug := range *plugins {
		plugin := plug.(api.PluginInfo)
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

func CreateBot(folderPath, botName, botVersion, botDesc string, adapters *[]list.Item, plugins *[]list.Item) (err error) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	logger := zerolog.New(output).With().Timestamp().Logger()
	mutex.Lock()
	progress = 0
	mutex.Unlock()
	botFolder := strings.ReplaceAll(filepath.Join(folderPath, botName), "\\", "/")
	if _, err = os.Stat(botFolder); err == nil || !os.IsNotExist(err) {
		logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Removing existing folder...", progress)
		if err = os.RemoveAll(botFolder); err != nil {
			logger.WithLevel(zerolog.ErrorLevel).Msgf("[%d/6] Error removing folder: %s", progress, err.Error())
			return
		}
		logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Removed successfully!", progress)
	}
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Creating bot folder...", progress)
	if err = os.Mkdir(botFolder, os.ModePerm); err != nil {
		logger.WithLevel(zerolog.ErrorLevel).Msgf("[%d/6] Error creating folder: %s", progress, err.Error())
		return
	}
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Created folder successfully!", progress)
	mutex.Lock()
	progress += 1
	mutex.Unlock()
	pwd, _ := os.Getwd()
	os.Chdir(botFolder)
	moduleName := FormatName(botName)
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Creating go module %s...", progress, moduleName)
	var cmdout []byte
	if cmdout, err = exec.Command("go", "mod", "init", moduleName).CombinedOutput(); err != nil {
		fmt.Println(string(cmdout))
		logger.WithLevel(zerolog.ErrorLevel).Msgf("[%d/6] Error creating go module: exited with %s", progress, err.Error())
		return
	}
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Created go module %s successfully!", progress, moduleName)
	mutex.Lock()
	progress += 1
	mutex.Unlock()
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Writing %s...", progress, moduleName)
	if err = writeGoFile(botFolder, moduleName, botVersion, botDesc, adapters, plugins); err != nil {
		logger.WithLevel(zerolog.ErrorLevel).Msgf("[%d/6] Error writing go file: exited with %s", progress, err.Error())
		return
	}
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Wrote %s.go successfully!", progress, moduleName)
	mutex.Lock()
	progress += 1
	mutex.Unlock()
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Formatting go file...", progress)
	if cmdout, err = exec.Command("go", "fmt").CombinedOutput(); err != nil {
		fmt.Println(string(cmdout))
		logger.WithLevel(zerolog.ErrorLevel).Msgf("[%d/6] Error formatting go file: exited with %s", progress, err.Error())
		return
	}
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Formatted go file successfully!", progress)
	mutex.Lock()
	progress += 1
	mutex.Unlock()
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Running go mod tidy(this should take some time)...", progress)
	if cmdout, err = exec.Command("go", "mod", "tidy").CombinedOutput(); err != nil {
		fmt.Println(string(cmdout))
		logger.WithLevel(zerolog.ErrorLevel).Msgf("[%d/6] Error running go mod tidy: exited with %s", progress, err.Error())
		return
	}
	fmt.Println(string(cmdout))
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Ran go mod tidy successfully!", progress)
	mutex.Lock()
	progress += 1
	mutex.Unlock()
	logger.WithLevel(zerolog.InfoLevel).Msgf("[%d/6] Bot created successfully!", progress)
	mutex.Lock()
	progress += 1
	mutex.Unlock()
	os.Chdir(pwd)
	return
}
