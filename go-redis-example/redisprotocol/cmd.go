package redisprotocol

type CmdName string

const (
	Cmd_Get CmdName = "get"
	Cmd_Set CmdName = "set"

	// redis protocol data type prifix
	PrifixString string = "+"
	PrifixError         = "-"
	PrifixNumber        = ":"
	PrifixBatch         = "$"
	PrifixArray         = "*"

	// spplitter
	Splitter string = "\r\n"
)

type ICmd interface {
	Name() CmdName
	Args() []interface{}
}
type Cmd struct {
	args []interface{}
	err  error
}

func NewCmd(opt CmdName, args []interface{}) *Cmd {
	cmd := new(Cmd)
	cmd.args = append(cmd.args, opt)
	cmd.args = append(cmd.args, args...)
	return cmd
}

func (cmd *Cmd) Name() CmdName {
	if len(cmd.args) == 0 {
		return ""
	}

	c, _ := cmd.args[0].(CmdName)
	return c
}

func (cmd *Cmd) Args() []interface{} {
	if len(cmd.args) < 2 {
		return nil
	}
	return cmd.args[1:]
}
