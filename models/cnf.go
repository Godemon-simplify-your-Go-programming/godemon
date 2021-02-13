package models

type Command struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Option string `json:"file"`
}

type Commands struct {
	Commands []Command `json:"commands"`
}

type Var struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Project struct {
	Name string `json:"name"`
	Arch string `json:"arch"`
	OS   string `json:"os"`
	Path string `json:"path"`
	Vars []Var  `json:"dev-vars"`
}
