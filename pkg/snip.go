package snip

type Snip struct {
	Title       string   `json:"title" yaml:"title"`
	Description string   `json:"description" yaml:"description"`
	Script      string   `json:"script" yaml:"script"`
	Labels      []string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
