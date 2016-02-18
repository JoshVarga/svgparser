package utils

import (
	"strconv"
	"strings"
)

var allCommands commands

type commands struct {
	all        []string
	parameters map[string]int
	start      string
	end        string
}

func getCommands() commands {
	var parameters = map[string]int{
		"m": 2, "z": 0, "l": 2, "h": 1, "v": 1,
		"c": 6, "s": 4, "q": 4, "t": 2, "a": 7,
	}
	var all []string
	for k := range parameters {
		all = append(all, k)
	}
	return commands{all, parameters, "m", "z"}
}

func (c *commands) isCommand(token string) bool {
	for _, command := range c.all {
		if strings.ToLower(token) == command {
			return true
		}
	}
	return false
}

type PathParserError struct {
	msg string
}

func (err PathParserError) Error() string {
	return err.msg
}

type Token struct {
	value    string
	operator bool
}

type Command struct {
	Symbol string
	Params []float64
}

func (c *Command) IsAbsolute() bool {
	return c.Symbol == strings.ToUpper(c.Symbol)
}

func (c *Command) Compare(o *Command) bool {
	if c.Symbol != o.Symbol {
		return false
	}
	for i, param := range c.Params {
		if param != o.Params[i] {
			return false
		}
	}
	return true
}

type Subpath struct {
	Commands []*Command
}

func (s *Subpath) Compare(o *Subpath) bool {
	if len(s.Commands) != len(o.Commands) {
		return false
	}
	for i, command := range s.Commands {
		if !command.Compare(o.Commands[i]) {
			return false
		}
	}
	return true
}

type Path struct {
	Subpaths []*Subpath
}

func (p *Path) Compare(o *Path) bool {
	if len(p.Subpaths) != len(o.Subpaths) {
		return false
	}
	for i, subpath := range p.Subpaths {
		if !subpath.Compare(o.Subpaths[i]) {
			return false
		}
	}
	return true
}

func reverse(ops []float64) []float64 {
	for i, j := 0, len(ops)-1; i < j; i, j = i+1, j-1 {
		ops[i], ops[j] = ops[j], ops[i]
	}
	return ops
}

func addOperand(tokens []Token, operand string) ([]Token, string) {
	if operand != "" {
		tokens = append(tokens, Token{operand, false})
		operand = ""
	}
	return tokens, operand
}

// STEP 1
func Tokenize(raw string) []Token {
	var (
		tokens  []Token
		operand string
	)
	for _, r := range raw {
		char := string(r)
		switch {
		case allCommands.isCommand(char):
			tokens, operand = addOperand(tokens, operand)
			tokens = append(tokens, Token{char, true})
		case char == ".":
			if operand == "" {
				operand = "0"
			}
			if strings.Contains(operand, char) {
				tokens = append(tokens, Token{operand, false})
				operand = "0"
			}
			fallthrough
		case char >= "0" && char <= "9":
			operand += char
		case char == "-":
			tokens, operand = addOperand(tokens, operand)
			operand = char
		default:
			tokens, operand = addOperand(tokens, operand)
		}
	}
	tokens, operand = addOperand(tokens, operand)
	return tokens
}

// STEP 2
func TokensToCommands(tokens []Token) ([]*Command, error) {
	var (
		commands []*Command
		operands []float64
	)
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if token.operator {
			nParam := allCommands.parameters[strings.ToLower(token.value)]
			nOperand := len(operands)
			if nParam == 0 && nOperand == 0 {
				command := &Command{Symbol: token.value}
				commands = append([]*Command{command}, commands...)
			} else if nParam != 0 && nOperand%nParam == 0 {
				for i := 0; i < nOperand/nParam; i++ {
					command := &Command{token.value, reverse(operands[:nParam])}
					commands = append([]*Command{command}, commands...)
					operands = operands[nParam:]
				}
			} else {
				err := PathParserError{"Incorrect number of parameters for " + token.value}
				return nil, err
			}
		} else {
			if number, err := strconv.ParseFloat(token.value, 64); err != nil {
				return nil, err
			} else {
				operands = append(operands, number)
			}
		}
	}
	return commands, nil
}

// STEP 3
func CreateSubpaths(commands []*Command) *Path {
	path := &Path{}
	var subpath []*Command
	for i, command := range commands {
		switch strings.ToLower(command.Symbol) {
		case allCommands.start:
			if len(subpath) > 0 {
				path.Subpaths = append(path.Subpaths, &Subpath{subpath})
			}
			subpath = []*Command{command}
		case allCommands.end:
			subpathWithEnd := append(subpath, command)
			path.Subpaths = append(path.Subpaths, &Subpath{subpathWithEnd})
		default:
			subpath = append(subpath, command)
			if len(commands) == i+1 {
				path.Subpaths = append(path.Subpaths, &Subpath{subpath})
			}
		}
	}
	return path
}

func PathParser(raw string) (*Path, error) {
	allCommands = getCommands()
	commands, err := TokensToCommands(Tokenize(raw))
	if err != nil {
		return nil, err
	}
	return CreateSubpaths(commands), nil
}
