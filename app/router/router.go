// Register all scenes here.
package router

import tea "github.com/charmbracelet/bubbletea"

type Scene struct {
	tea.Model
	Name string
}

var Scenes = make(map[string]tea.Model)

func Init() tea.Cmd {
	cmds := make([]tea.Cmd, len(Scenes))
	for _, scene := range Scenes {
		cmds = append(cmds, scene.Init())
	}
	return tea.Batch(cmds...)
}

func RegisterScene(name string, scene tea.Model) {
	sceneInstance := scene.(Scene)
	sceneInstance.Name = name
	Scenes[name] = sceneInstance
}

func GetScene(name string) Scene {
	return Scenes[name].(Scene)
}

func Update(name string, msg tea.Msg) (next string, cmd tea.Cmd) {
	nextModel, cmd := Scenes[name].Update(msg)
	nextScene := nextModel.(Scene)
	next = nextScene.Name
	Scenes[nextScene.Name] = nextScene
	return
}

func View(name string) string {
	return Scenes[name].View()
}
