package models

type Command struct {
	Name   string `json:"name"`
	Option string `json:"file"`
	Path   string `json:"path"` // optional, for single file support
}
