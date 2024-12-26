package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

func getLatestCommit() (commit CommitInfo, err error) {
	resp, err := http.Get("https://api.github.com/repos/gonebot-dev/gonerepo/commits/main")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Failed to get latest commit: %s", resp.Status)
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(&commit); err != nil {
		return
	}
	return
}

func getTree(commit CommitInfo) (files []FileInfo, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/gonebot-dev/gonerepo/git/trees/%s?recursive=true", commit.SHA))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Failed to get tree: %s", resp.Status)
		return
	}
	var treeResp TreeResponse
	if err = json.NewDecoder(resp.Body).Decode(&treeResp); err != nil {
		return
	}
	files = treeResp.Tree
	return
}

var Plugins = make([]list.Item, 0)
var Adapters = make([]list.Item, 0)

// Which commit is currently being used
var CurrentCommit = CommitInfo{SHA: ""}
var Finished = true
var Err error

func SyncRepo() {
	Finished = false
	Err = nil
	Plugins = make([]list.Item, 0)
	Adapters = make([]list.Item, 0)

	latestCommit, Err := getLatestCommit()
	if Err != nil {
		Finished = true
		return
	}
	if latestCommit.SHA == CurrentCommit.SHA {
		Finished = true
		return
	}
	CurrentCommit = latestCommit
	files, Err := getTree(latestCommit)
	if Err != nil {
		Finished = true
		return
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Path, ".json") {
			continue
		}
		if strings.HasPrefix(file.Path, "packages/adapters") {
			var adapter AdapterInfo
			var resp *http.Response
			resp, Err = http.Get(fmt.Sprintf(
				"https://github.com/gonebot-dev/gonerepo/raw/refs/heads/main/%s",
				file.Path,
			))
			if Err != nil {
				Finished = true
				return
			}
			defer resp.Body.Close()
			if Err = json.NewDecoder(resp.Body).Decode(&adapter); Err != nil {
				Finished = true
				return
			}
			pieces := strings.Split(strings.TrimSuffix(file.Path, ".json"), "/")
			adapter.Name = pieces[len(pieces)-1]
			Adapters = append(Adapters, adapter)
		} else if strings.HasPrefix(file.Path, "packages/plugins") {
			var plugin PluginInfo
			var resp *http.Response
			resp, Err = http.Get(fmt.Sprintf(
				"https://github.com/gonebot-dev/gonerepo/raw/refs/heads/main/%s",
				file.Path,
			))
			if Err != nil {
				Finished = true
				return
			}
			defer resp.Body.Close()
			if Err = json.NewDecoder(resp.Body).Decode(&plugin); Err != nil {
				Finished = true
				return
			}
			pieces := strings.Split(strings.TrimSuffix(file.Path, ".json"), "/")
			plugin.Name = pieces[len(pieces)-1]
			Plugins = append(Plugins, plugin)
		}
	}
	Finished = true
	return
}
