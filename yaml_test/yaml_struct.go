package yaml_test

//MetadataS is struct
type MetadataS struct {
	name   string `yaml:"name"`
	labels string `yaml:"labels"`
}

//SelectorS is struct
type SelectorS struct {
	matchLabels string `yaml:"matchLabels"`
}

//TempSpec is struct
type TempSpec struct {
	containers   string
	nodeSelector string
	tolerations  string
}

//TemplateS is struct
type TemplateS struct {
	metadata MetadataS `yaml:"metadata"`
	spec     TempSpec  `yaml:"spec"`
}

//SpecS is struct
type SpecS struct {
	replicas int       `yaml:"replicas"`
	selector SelectorS `yaml:"selector"`
	template TemplateS `yaml:"template"`
}

//Deployment is struct
type Deployment struct {
	apiVersion string    `yaml:"apiVersion"`
	kind       string    `yaml:"kind"`
	metadata   MetadataS `yaml:"metadata"`
	spec       SpecS     `yaml:"spec"`
}
