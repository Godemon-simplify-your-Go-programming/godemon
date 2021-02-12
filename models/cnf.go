package models

type Command struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Option string `json:"file"`
}

type Commands struct {
	Commands []Command `json:"commands"`
}
