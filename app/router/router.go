// Register all scenes here.
package router

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Scene interface {
	tea.Model
	Name() string
}

type SwitchSceneMsg struct {
	Next string
}

func NextScene(name string) tea.Cmd {
	return func() tea.Msg {
		return SwitchSceneMsg{
			Next: name,
		}
	}
}

type EchoMsg struct{}

var EchoTick = func() tea.Msg {
	return EchoMsg{}
}

var scenes = make(map[string]tea.Model)

func Init(name string) tea.Cmd {
	return scenes[name].Init()
}

func RegisterScene(scene Scene) {
	scenes[scene.Name()] = scene
}

func Update(name string, msg tea.Msg) (next string, cmd tea.Cmd) {
	nextModel, cmd := scenes[name].Update(msg)
	nextScene := nextModel.(Scene)
	next = nextScene.Name()
	if next != name {
		scenes[nextScene.Name()], cmd = nextScene.Update(msg)
	} else {
		scenes[nextScene.Name()] = nextScene
	}
	return
}

func View(name string) string {
	return scenes[name].View()
}
