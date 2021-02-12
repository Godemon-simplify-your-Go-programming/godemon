package models

type Command struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Option string `json:"file"`
}

type Project struct {
	Name string `json:"name"`
	OS   string `json:"platformOS"`
	Arch string `json:"platformArch"`
	Path string `json:"path"`
}

type Commands struct {
	Commands []Command `json:"commands"`
}

type ProjectInfo struct {
	Project  Project `json:"project"`
	Commands Commands
}
