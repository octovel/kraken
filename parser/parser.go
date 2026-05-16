package parser

import (
	"fmt"
	"os"
)

// Creates and returns a new parser instance with [Options].
func New(options *Options) *Parser {
	if options == nil {
		return &Parser{
			Options: *DefaultOptions(),
		}
	}

	_options := *options

	if _options.FlagPrefix == "" {
		_options.FlagPrefix = "--"
	}
	if _options.ValueDelimiter == ' ' {
		_options.ValueDelimiter = '='
	}

	fixOption := func(value *int) {
		if *value == 0 {
			*value = -1
		}
	}

	fixOption(&_options.MaximumTokenLength)
	fixOption(&_options.MaximumFlags)
	fixOption(&_options.MinimumFlags)
	fixOption(&_options.MaximumTokens)
	fixOption(&_options.MinimumTokens)
	fixOption(&_options.MaximumArguments)
	fixOption(&_options.MinimumArguments)

	return &Parser{
		Options: _options,
	}
}

// Returns the a pointer to a struct representing the
// default options for the parser.
func DefaultOptions() *Options {
	return &Options{
		FlagPrefix:               "--",
		ValueDelimiter:           '=',
		AllowStandaloneValues:    true,
		HandleExecutableArgument: true,
		MaximumTokenLength:       -1,
		MaximumFlags:             -1,
		MinimumFlags:             -1,
		MaximumTokens:            -1,
		MinimumTokens:            -1,
		MaximumArguments:         -1,
		MinimumArguments:         -1,
	}
}

// Parses the given arguments vector into a structured data struct
// you can interact with. It is used by other components of this library.
func (parser *Parser) Parse(arguments []string) (*Result, error) {
	_arguments := arguments
	if parser.Options.HandleExecutableArgument {
		if len(_arguments) != 0 && _arguments[0] == os.Args[0] {
			_arguments = arguments[1:]
		}
	}

	// make a call to the internal parser
	result, err := parse(_arguments, parser)
	if err != nil {
		return nil, fmt.Errorf("an error occured during the parsing process: %s", err.Error())
	}

	if parser.Options.MinimumTokens != -1 && len(_arguments) < parser.Options.MinimumTokens {
		return nil, fmt.Errorf("too few arguments (expected %d, got %d)", parser.Options.MinimumTokens, len(_arguments))
	}
	if parser.Options.MaximumTokens != -1 && len(_arguments) > parser.Options.MaximumTokens {
		return nil, fmt.Errorf("too many arguments (expected %d, got %d)", parser.Options.MaximumTokens, len(_arguments))
	}

	if parser.Options.ErrorOnNoCommand && len(arguments) == 0 {
		return nil, fmt.Errorf("no command provided")
	}

	return result, err
}

func parse(arguments []string, parser *Parser) (*Result, error) {
	// temp
	return nil, nil
}

type Parser struct {
	Options Options
	Result  Result
}

// Represents the result of the parser.
type Result struct{}

// Options are used the change and customise the behaviour of the parser.
// For the default value of each option, see https://wiki.octovel.org/kraken
type Options struct {
	// Define a custom flag prefix. If the token the parser is iterating
	// over starts with this string, it is thus considered a flag by
	// the parser. The default value for this option is:
	//  "--"
	FlagPrefix string
	// Define a custom name and value delimiter for flags. If the flag the
	// parser is iterating over has this rune, it splits it from there into
	// two parts; the name and the value. The default value for this option is:
	//  "="
	ValueDelimiter rune
	// Similar to [Options.FlagPrefix], define a custom short flags prefix.
	// If the token the parser is iterating over starts with this string, it
	// is thus considered short flags and split into all its runes. The default
	// value for this option is:
	//  "" // disabled by default
	GroupShortFlagsPrefix rune
	// If enabled, it treats flags without a value as boolean a flag.
	// The default value for this option is:
	//  true
	AllowStandaloneValues bool
	// If enabled, it seeks for the executable file path in provided tokens and
	// omits it automatically if found. We recommend enabling this option.
	// The default value for this option is:
	//  true
	HandleExecutableArgument bool
	// If enabled, treats all tokens starting with one of the
	// [Options.GroupShortFlagsPrefix] as group short flags.
	// Note that this property overrides [Options.FlagPrefix].
	// The default value for this option is:
	//  false
	EnableGroupedShortFlags bool
	// If enabled, treats all tokens left as arguments,
	// even if they start with a flag prefix.
	EnableDoubleDashStop bool
	// Define the maximum length of an input. If exceeded, the parser
	// returns an error. Set it to -1 to remove the limit. The default
	// value for this option is:
	//  -1
	MaximumTokenLength int
	// Define a limit of flags the user is not allowed exceed.
	// If exceeded, the parser returns an error. Set it to -1
	// to remove the limit. The default value for this option is:
	//  -1
	MaximumFlags int
	// Define a limit of flags the user has to input.
	// If not met, the parser returns an error. Set it to -1
	// to remove the limit. The default value for this option is:
	//  -1
	MinimumFlags int
	// Define a limit of tokens the user is not allowed exceed.
	// If exceeded, the parser returns an error. Set it to -1
	// to remove the limit. The default value for this option is:
	//  -1
	MaximumTokens int
	// Define a limit of tokens the user has to input.
	// If not met, the parser returns an error. Set it to -1
	// to remove the limit. The default value for this option is:
	//  -1
	MinimumTokens int
	// Define a limit of arguments the user is not allowed exceed.
	// If exceeded, the parser returns an error. Set it to -1
	// to remove the limit. The default value for this option is:
	//  -1
	MaximumArguments int
	// Define a limit of arguments the user has to input.
	// If not met, the parser returns an error. Set it to -1
	// to remove the limit. The default value for this option is:
	//  -1
	MinimumArguments int
	// If enabled, the parser will return an error if it
	// encounters the same flag twice. Otherwise the duplicate
	// flag is ignored. The default value for this option is:
	//  false
	ErrorOnDuplicate bool
	// If enabled, the parser will return an error if
	// no command was found during parsing. Otherwise the
	// process continues. The default value for this option is:
	//  false
	ErrorOnNoCommand bool
	// If enabled, the parser will return an error if
	// if finds a flag after an argument. The default value
	// for this option is:
	//  false
	StrictOrdering bool
}
