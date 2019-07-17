package definition

type Job struct {
	Name    string            `yaml:"name"`
	Command string            `yaml:"command"`
	Option  map[string]string `yaml:"option,omitempty"`
}

type Dag struct {
	Name   string            `yaml:"name"`
	Option map[string]string `yaml:"option"`
	Jobs   []Job             `yaml:"jobs,omitempty"`
}
