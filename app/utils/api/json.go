package api

import (
	"github.com/gonebot-dev/gonebuilder-tui/app/base"
)

type CommitInfo struct {
	SHA string `json:"sha"`
}

type FileInfo struct {
	Path string `json:"path"`
}

type TreeResponse struct {
	Tree []FileInfo `json:"tree"`
}

type PluginInfo struct {
	Name    string `json:"name"`
	Plugin  string `json:"plugin"`
	Desc    string `json:"description"`
	DescZH  string `json:"description_zh"`
	Package string `json:"package"`
}

func (pi PluginInfo) Title() string { return pi.Name }
func (pi PluginInfo) Description() string {
	return base.IfElse(base.Lang == "en", pi.Desc, pi.DescZH)
}
func (pi PluginInfo) FilterValue() string {
	return pi.Name + base.IfElse(base.Lang == "en", pi.Desc, pi.DescZH)
}

type AdapterInfo struct {
	Name    string `json:"name"`
	Adapter string `json:"adapter"`
	Desc    string `json:"description"`
	DescZH  string `json:"description_zh"`
	Package string `json:"Package"`
}

func (ai AdapterInfo) Title() string { return ai.Name }
func (ai AdapterInfo) Description() string {
	return base.IfElse(base.Lang == "en", ai.Desc, ai.DescZH)
}
func (ai AdapterInfo) FilterValue() string {
	return ai.Name + base.IfElse(base.Lang == "en", ai.Desc, ai.DescZH)
}
