package communitymodules

// This structure contains only the fields currently in use.
type Modules []Module

type Module struct {
	Name     string    `json:"name,omitempty"`
	Versions []Version `json:"versions,omitempty"`
}

type Version struct {
	Version     string `json:"version,omitempty"`
	ManagerPath string `json:"managerPath,omitempty"`
	Repository  string `json:"repository,omitempty"`
}
