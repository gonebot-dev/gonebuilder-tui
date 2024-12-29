package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

func getLatestCommit() (commit CommitInfo, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/commits/main", os.Getenv("GONEREPO")), nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to get latest commit: %s", resp.Status)
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(&commit); err != nil {
		return
	}
	return
}

func getTree(commit CommitInfo) (files []FileInfo, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/git/trees/%s?recursive=true", os.Getenv("GONEREPO"), commit.SHA), nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to get tree: %s", resp.Status)
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
			req, Err := http.NewRequest("GET", fmt.Sprintf(
				"https://github.com/%s/raw/refs/heads/main/%s",
				os.Getenv("GONEREPO"),
				file.Path,
			), nil)
			if Err != nil {
				return
			}
			resp, Err = http.DefaultClient.Do(req)
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
			req, Err := http.NewRequest("GET", fmt.Sprintf(
				"https://github.com/%s/raw/refs/heads/main/%s",
				os.Getenv("GONEREPO"),
				file.Path,
			), nil)
			if Err != nil {
				return
			}
			resp, Err = http.DefaultClient.Do(req)
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
}
