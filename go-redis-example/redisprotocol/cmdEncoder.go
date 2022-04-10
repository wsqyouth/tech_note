package redisprotocol

import (
	"fmt"
	"strings"
)

type ICmdEncoder interface {
	Encode(cmd ICmd) string
}
type CmdEncoder struct {
}

func NewCmdEncoder() ICmdEncoder {
	return ICmdEncoder(new(CmdEncoder))
}

func (cmdEncoder *CmdEncoder) Encode(cmd ICmd) string {
	cmdLen := len(cmd.Args()) + 1

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s%d%s", PrifixArray, cmdLen, Splitter))
	builder.WriteString(fmt.Sprintf("%s%d%s", PrifixBatch, len(cmd.Name()), Splitter))
	builder.WriteString(fmt.Sprintf("%s%s", cmd.Name(), Splitter))

	for _, v := range cmd.Args() {
		builder.WriteString(fmt.Sprintf("%s%d%s", PrifixBatch, len(v.(string)), Splitter))
		builder.WriteString(fmt.Sprintf("%s%s", v.(string), Splitter))
	}
	return builder.String()
}
