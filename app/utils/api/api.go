package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

var Plugins = make([]PluginInfo, 0)
var Adapters = make([]AdapterInfo, 0)

// Which commit is currently being used
var CurrentCommit = CommitInfo{SHA: ""}

func SyncRepo() (err error) {
	err = nil
	Plugins = make([]PluginInfo, 0)
	Adapters = make([]AdapterInfo, 0)

	latestCommit, err := getLatestCommit()
	if err != nil {
		return
	}
	if latestCommit.SHA == CurrentCommit.SHA {
		return
	}
	CurrentCommit = latestCommit
	files, err := getTree(latestCommit)
	if err != nil {
		return
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Path, ".json") {
			continue
		}
		if strings.HasPrefix(file.Path, "packages/adapters") {
			var adapter AdapterInfo
			var resp *http.Response
			resp, err = http.Get(fmt.Sprintf(
				"https://github.com/gonebot-dev/gonerepo/raw/refs/heads/main/%s",
				file.Path,
			))
			if err != nil {
				return
			}
			defer resp.Body.Close()
			if err = json.NewDecoder(resp.Body).Decode(&adapter); err != nil {
				return
			}
			pieces := strings.Split(strings.TrimSuffix(file.Path, ".json"), "/")
			adapter.Name = pieces[len(pieces)-1]
			Adapters = append(Adapters, adapter)
		} else if strings.HasPrefix(file.Path, "packages/plugins") {
			var plugin PluginInfo
			var resp *http.Response
			resp, err = http.Get(fmt.Sprintf(
				"https://github.com/gonebot-dev/gonerepo/raw/refs/heads/main/%s",
				file.Path,
			))
			if err != nil {
				return
			}
			defer resp.Body.Close()
			if err = json.NewDecoder(resp.Body).Decode(&plugin); err != nil {
				return
			}
			pieces := strings.Split(strings.TrimSuffix(file.Path, ".json"), "/")
			plugin.Name = pieces[len(pieces)-1]
			Plugins = append(Plugins, plugin)
		}
	}
	return
}
