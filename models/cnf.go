package models

type Command struct {
	Name   string `json:"name"`
	Option string `json:"file"`
}

type Var struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Project struct {
	Name     string    `json:"name"`
	Arch     string    `json:"arch"`
	OS       string    `json:"os"`
	Vars     []Var     `json:"dev-vars"`
	Commands []Command `json:"commands"`
}

type ErrorTMP struct {
	Message string `json:"message"`
}
