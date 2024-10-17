package command

import (
	"fmt"
	"strings"
)

// Command 命令，仅作为描述
type Command struct {
	command []string
	desc    string
}

// NewCommand 创建一个命令
//
//	desc: 描述
//	cms: 命令
func NewCommand(desc string, cms ...string) Command {
	return Command{
		command: cms,
		desc:    desc,
	}
}

type Commands []Command

// NewCommands 创建一个命令列表
func NewCommands(cms ...Command) Commands {
	var c Commands
	c = append(c, cms...)
	return c
}

func (c Commands) String() string {
	var builder strings.Builder
	for _, command := range c {
		for _, oneCommand := range command.command {
			builder.WriteString(fmt.Sprintf("%s ", oneCommand))
		}
		builder.WriteString(fmt.Sprintf(": %s\n", command.desc))
	}
	return builder.String()
}
