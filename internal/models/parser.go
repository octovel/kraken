package models

import "fmt"

type Parser struct {
	Options ParserOptions
}

func (p *Parser) Parse(args []string) (*ParserResult, error) {
	if p.Options.MinimumInputs != -1 && len(args) < p.Options.MinimumInputs {
		return nil, fmt.Errorf("too few arguments (expected %d, got %d)", len(args), p.Options.MinimumInputs)
	}
	if p.Options.MaximumInputs != -1 && len(args) > p.Options.MaximumInputs {
		return nil, fmt.Errorf("too many arguments (expected %d, got %d)", len(args), p.Options.MinimumInputs)
	}

	return nil, nil
}

type ParserResult struct{}

type ParserOptions struct {
	FlagPrefixes            []string
	ValueDelimiters         []string
	AllowStandaloneValues   bool
	EnableGroupedShortFlags bool
	EnableDoubleDashStop    bool

	MaximumTokenLength int

	MaximumFlags     int
	MinimumFlags     int
	MaximumInputs    int
	MinimumInputs    int
	MaximumArguments int
	MinimumArguments int

	ErrorOnDuplicate bool
	StrictOrdering   bool
}

func DefaultSettings() *ParserOptions {
	return &ParserOptions{
		FlagPrefixes:          []string{"--"},
		ValueDelimiters:       []string{"="},
		AllowStandaloneValues: true,
		MaximumTokenLength:    -1,
		MaximumFlags:          -1,
		MinimumFlags:          -1,
		MaximumInputs:         -1,
		MinimumInputs:         -1,
		MaximumArguments:      -1,
		MinimumArguments:      -1,
	}
}
