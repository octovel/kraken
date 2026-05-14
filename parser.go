package parser

import (
	"fmt"
)

type Parser struct {
	Options Options
}

type Result struct{}

func (p *Parser) Parse(args []string) (*Result, error) {
	if p.Options.MinimumInputs != -1 && len(args) < p.Options.MinimumInputs {
		return nil, fmt.Errorf("too few arguments (expected %d, got %d)", p.Options.MinimumInputs, len(args))
	}
	if p.Options.MaximumInputs != -1 && len(args) > p.Options.MaximumInputs {
		return nil, fmt.Errorf("too many arguments (expected %d, got %d)", p.Options.MaximumInputs, len(args))
	}

	if p.Options.ErrorOnNoCommand && len(args) == 0 {
		return nil, fmt.Errorf("no command provided")
	}

	return nil, nil
}

type Options struct {
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
	ErrorOnNoCommand bool
	StrictOrdering   bool
}

func New(options *Options) *Parser {
	if options == nil {
		return &Parser{
			Options: *DefaultSettings(),
		}
	}

	_options := *options

	if len(_options.FlagPrefixes) == 0 {
		_options.FlagPrefixes = []string{"--"}
	}
	if len(_options.ValueDelimiters) == 0 {
		_options.ValueDelimiters = []string{"="}
	}

	fixOption := func(value *int) {
		if *value == 0 {
			*value = -1
		}
	}

	fixOption(&_options.MaximumTokenLength)
	fixOption(&_options.MaximumFlags)
	fixOption(&_options.MinimumFlags)
	fixOption(&_options.MaximumInputs)
	fixOption(&_options.MinimumInputs)
	fixOption(&_options.MaximumArguments)
	fixOption(&_options.MinimumArguments)

	return &Parser{
		Options: _options,
	}
}

func DefaultSettings() *Options {
	return &Options{
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
