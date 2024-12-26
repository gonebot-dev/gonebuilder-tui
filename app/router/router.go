// Register all scenes here.
package router

import tea "github.com/charmbracelet/bubbletea"

type Scene interface {
	tea.Model
	Name() string
	GetEmits() map[string]string
}

var scenes = make(map[string]tea.Model)

func Init() tea.Cmd {
	cmds := make([]tea.Cmd, len(scenes))
	for _, scene := range scenes {
		cmds = append(cmds, scene.Init())
	}
	return tea.Batch(cmds...)
}

func RegisterScene(scene Scene) {
	scenes[scene.Name()] = scene
}

func GetScene(name string) (Scene, tea.Cmd) {
	return scenes[name].(Scene), scenes[name].Init()
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
