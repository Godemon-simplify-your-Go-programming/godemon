package models

type Project struct {
	Name     string    `json:"name"`
	Arch     string    `json:"arch"`
	OS       string    `json:"os"`
	Vars     []Var     `json:"dev-vars"`
	Commands []Command `json:"commands"`
	Files    []File    `json:"files"`
}
