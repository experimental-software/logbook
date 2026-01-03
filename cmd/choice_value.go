// Package cmd see https://github.com/spf13/pflag/issues/236#issuecomment-3098646174
package cmd

import "fmt"

type ChoiceValue struct {
	value    string
	validate func(string) error
}

func (f *ChoiceValue) Set(s string) error {
	err := f.validate(s)
	if err != nil {
		return err
	}

	f.value = s
	return nil
}
func (f *ChoiceValue) Type() string   { return "string" }
func (f *ChoiceValue) String() string { return f.value }

func StringChoice(choices []string) *ChoiceValue {
	return &ChoiceValue{
		validate: func(s string) error {
			for _, choice := range choices {
				if s == choice {
					return nil
				}
			}
			return fmt.Errorf("must be one of %v", choices)
		},
	}
}
