/*
   Copyright 2021 Edson Michaque

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package snip

type SnipList map[string]Snip

type SnipOption func(*Snip)

type Snip struct {
	Title       string            `json:"title" yaml:"title"`
	Description string            `json:"description" yaml:"description"`
	Script      string            `json:"script" yaml:"script"`
	Labels      map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

func New(options ...SnipOption) *Snip {
	newSnip := &Snip{}

	for _, option := range options {
		option(newSnip)
	}

	return newSnip
}

func NewList() map[string]*Snip {
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

func WithLabels(labels map[string]string) SnipOption {
	return func(s *snip) {
		s.Labels = labels
	}
}
