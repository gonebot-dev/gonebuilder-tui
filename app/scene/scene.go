package scene

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Scene interface {
	// Update the scene, return the next scene and any bubbletea command.
	//
	// Returns the original scene if the scene should not change.
	Update(msg tea.Msg) (Scene, string, tea.Cmd)
	// View of the scene.
	View() string
}

var Scenes = make(map[string]Scene)

func RegisterScene(name string, scene Scene) {
	Scenes[name] = scene
}

func GetScene(name string) Scene {
	return Scenes[name]
}

func Update(name string, msg tea.Msg) (next string, cmd tea.Cmd) {
	Scenes[name], next, cmd = Scenes[name].Update(msg)
	return
}

func View(name string) string {
	return Scenes[name].View()
}
