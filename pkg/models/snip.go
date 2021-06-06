package models

type SnipList map[string]Snip

type SnipOption func(*Snip)

type Snip struct {
	Title       string   `json:"title" yaml:"title"`
	Description string   `json:"description" yaml:"description"`
	Script      string   `json:"script" yaml:"script"`
	Labels      []string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

func NewSnip(options ...SnipOption) *Snip {
	newSnip := &Snip{}

	for _, option := range options {
		option(newSnip)
	}

	return newSnip
}

func NewSnipList() map[string]*Snip {
	return map[string]*Snip{}
}

func WithTitle(title string) SnipOption {
	return func(s *Snip) {
		s.Title = title
	}
}

func WithDescription(description string) SnipOption {
	return func(s *Snip) {
		s.Description = description
	}
}

func WithScript(script string) SnipOption {
	return func(s *Snip) {
		s.Script = script
	}
}
