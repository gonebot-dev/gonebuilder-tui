package api

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
	Name          string `json:"name"`
	Plugin        string `json:"plugin"`
	Description   string `json:"description"`
	DescriptionZH string `json:"description_zh"`
	Package       string `json:"package"`
}

type AdapterInfo struct {
	Name          string `json:"name"`
	Adapter       string `json:"adapter"`
	Description   string `json:"description"`
	DescriptionZH string `json:"description_zh"`
	Package       string `json:"Package"`
}
