package models

type Command struct {
	Name   string   `json:"name"`
	Option string   `json:"option"`
	Flags  []string `json:"flags"`
	Path   string   `json:"path"` // optional, for single file support
}
