package definition

type Job struct {
	Name         string            `yaml:"name"`
	Commands     []string          `yaml:"commands"`
	Option       map[string]string `yaml:"option,omitempty"`
	Dependencies []string          `yaml:"dependencies"`
}

type Dag struct {
	Name   string            `yaml:"name"`
	Option map[string]string `yaml:"option"`
	Jobs   []Job             `yaml:"jobs,omitempty"`
}
