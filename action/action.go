package action

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/PaluMacil/cc/config"
)

const (
	// Error Types
	ErrNoCommand      = "no command given"
	ErrInvalidCommand = "invalid command"
	ErrTooFewArgs     = "too few arguments given"
	// Command Types
	CommandAdd    = "add"
	CommandRemove = "remove"
	CommandSelect = "select"
	CommandList   = "list"
	CommandStatus = "status"
)

// Action holds command and argument information and
// provides and Execute method that takes an io.Writer
type Action struct {
	cmd  string
	arg  string
	arg2 string
}

// Execute performs an Action and, if there is output,
// write it to an io.Writer.
func (a Action) Execute(out io.Writer) error {
	log.Println("Executing action for", a.cmd)
	switch strings.Trim(a.cmd, " \n\r\t") {
	case CommandAdd:
		return add(a.arg, a.arg2)
	case CommandRemove:
		return remove(a.arg)
	case CommandSelect:
		return selectDefault(a.arg)
	case CommandList:
		list, err := list()
		if err != nil {
			return err
		}
		_, err = io.WriteString(out, list)
		return err
	case CommandStatus:
		status, err := status()
		if err != nil {
			return err
		}
		_, err = io.WriteString(out, status)
		return err
	default:
		return fmt.Errorf(ErrInvalidCommand)
	}
}

// New constructs an Action from the parameters
func New(cmd, arg, arg2 string) (Action, error) {
	if cmd == "" {
		return Action{}, fmt.Errorf(ErrNoCommand)
	}
	return Action{cmd, arg, arg2}, nil
}

// FromOSArgs constructs an Action from commandline arguments
func FromOSArgs(osArgs []string) (Action, error) {
	var action Action
	if len(osArgs) < 2 {
		return action, fmt.Errorf(ErrNoCommand)
	}
	action.cmd = osArgs[1]
	if len(osArgs) > 2 {
		action.arg = osArgs[2]
	}
	if len(osArgs) > 3 {
		action.arg2 = strings.Join(osArgs[3:], " ")
	}
	return action, nil
}

func add(name, path string) error {
	configPath := config.Path()
	conf, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("loading config during add action: %s", err)
	}

	conf.Compilers = append(conf.Compilers, config.Compiler{name, path})

	if err = conf.Save(configPath); err != nil {
		return fmt.Errorf("saving config during add action: %s", err)
	}

	return nil
}

func remove(name string) error {
	configPath := config.Path()
	conf, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("loading config during remove action: %s", err)
	}

	var newList []config.Compiler
	for _, compiler := range conf.Compilers {
		if compiler.Name != name {
			newList = append(newList, compiler)
		}
	}
	conf.Compilers = newList
	if conf.Default == name {
		conf.Default = ""
	}

	if err = conf.Save(configPath); err != nil {
		return fmt.Errorf("saving config during remove action: %s", err)
	}

	return nil
}

func selectDefault(name string) error {
	configPath := config.Path()
	conf, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("loading config during select action: %s", err)
	}

	conf.Default = name

	if err := conf.Save(configPath); err != nil {
		return fmt.Errorf("saving config during select action: %s", err)
	}

	return nil
}

func list() (string, error) {
	configPath := config.Path()
	conf, err := config.Load(configPath)
	if err != nil {
		return "", fmt.Errorf("loading config during list action: %s", err)
	}
	var b strings.Builder
	for _, c := range conf.Compilers {
		b.WriteString(c.Name)
		b.WriteString("\t")
		b.WriteString(c.Path)
		b.WriteString("\n")
	}

	return b.String(), nil
}

func status() (string, error) {
	configPath := config.Path()
	conf, err := config.Load(configPath)
	if err != nil {
		return "", fmt.Errorf("loading config during status action: %s", err)
	}

	return conf.Default, nil
}
